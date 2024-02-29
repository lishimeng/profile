package profile

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/amqp/rabbit"
	"github.com/lishimeng/app-starter/tool"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/profile/cmd/profile/models"
	"github.com/lishimeng/profile/internal/db/model"
)

type CreateProfileResp struct {
	app.Response
	Id string `json:"id,omitempty"`
}

type CreateProfileReq struct {
	Code string `json:"code,omitempty"` // user code
}

func createProfile(ctx iris.Context) {

	log.Info("create user profile")
	var err error
	var req CreateProfileReq
	var resp CreateProfileResp

	err = ctx.ReadJSON(&req)
	if err != nil {
		log.Info("read json fail")
		log.Info(err)
		resp.Code = tool.RespCodeError
		tool.ResponseJSON(ctx, resp)
		return
	}
	log.Info("req.code:%s", req.Code)
	if req.Code == "" {
		log.Info("code is empty")
		resp.Code = tool.RespCodeError
		tool.ResponseJSON(ctx, resp)
		return
	}

	list, err := serviceGetProfile(req.Code)
	if err != nil {
		log.Info(err)
		resp.Code = tool.RespCodeError
		tool.ResponseJSON(ctx, resp)
		return
	}
	if len(list) > 0 {
		log.Info("duplicate: %s", req.Code)
		resp.Code = tool.RespCodeError
		resp.Message = "duplicate user_code"
		tool.ResponseJSON(ctx, resp)
		return
	}

	var mqReq models.BasicProfileCreate
	mqReq.Uid = req.Code
	var tx = rabbit.Message{
		Payload: mqReq,
		Router:  rabbit.Route{Exchange: rabbit.DefaultExchange, Key: "mq_profile_create"},
	}
	tx.SetOption(rabbit.UUIDMsgIdOption, rabbit.JsonEncodeOption)
	err = app.GetAmqp().Publish(tx)
	if err != nil {
		log.Info(err)
		resp.Code = tool.RespCodeError
		resp.Message = "duplicate user_code"
		tool.ResponseJSON(ctx, resp)
		return
	}

	resp.Code = tool.RespCodeSuccess
	resp.Id = req.Code
	tool.ResponseJSON(ctx, resp)
}

type GetProfileResp struct {
	app.Response
	UserCode              string `json:"userCode,omitempty"`
	RealName              string `json:"realName,omitempty"`
	IdCard                string `json:"idCard,omitempty"`
	IdCardVerified        bool   `json:"idCardVerified,omitempty"`
	PhoneNumber           string `json:"PhoneNumber,omitempty"`
	PhoneNumberVerified   bool   `json:"phoneNumberVerified,omitempty"`
	WechatUnionId         string `json:"wechatUnionId,omitempty"`
	WechatUnionIdVerified bool   `json:"wechatUnionIdVerified,omitempty"`
	CreateTime            int64  `json:"createTime,omitempty"`
}

func getProfileSpec(ctx iris.Context) {
	var err error
	var resp GetProfileResp

	var code = ctx.Params().Get("code")
	log.Info("get profile, code: %s", code)

	list, err := serviceGetProfile(code)
	if err != nil {
		log.Info(err)
		resp.Code = tool.RespCodeNotFound
		tool.ResponseJSON(ctx, resp)
		return
	}

	if len(list) == 0 {
		log.Info("no profile of: %s", code)
		resp.Code = tool.RespCodeNotFound
		tool.ResponseJSON(ctx, resp)
		return
	}

	var p = list[0]

	resp.Code = tool.RespCodeSuccess
	resp.UserCode = p.UserCode
	resp.RealName = p.RealName
	resp.IdCard = p.IdCard
	resp.IdCardVerified = p.IdCardVerified == model.Verified
	resp.CreateTime = p.CreateTime.Unix()
	tool.ResponseJSON(ctx, resp)
}
