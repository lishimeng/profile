package profile

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/tool"
	"github.com/lishimeng/go-log"
)

type CreateProfileResp struct {
	app.Response
	Id string `json:"id,omitempty"`
}

type CreateProfileReq struct {
	Code string `json:"code,omitempty"`
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

	p, err := serviceCreateProfile(req.Code)
	if err != nil {
		log.Info("createProfile fail")
		log.Info(err)
		resp.Code = tool.RespCodeError
		tool.ResponseJSON(ctx, resp)
		return
	}

	resp.Code = tool.RespCodeSuccess
	resp.Id = p.UserCode
	tool.ResponseJSON(ctx, resp)
}
