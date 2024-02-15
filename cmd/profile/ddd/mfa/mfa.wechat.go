package mfa

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/factory"
	"github.com/lishimeng/app-starter/tool"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/go-sdk/wechat"
)

type WechatBindReq struct {
	UnionId string
	Code    string
}

// 绑定微信union id
func bindWechat(ctx iris.Context) {

	log.Info("bind wx phone_number")

	var err error
	var req WechatBindReq
	var resp app.Response

	err = ctx.ReadJSON(&req)
	if err != nil {
		log.Info(err)
		resp.Code = tool.RespCodeError
		tool.ResponseJSON(ctx, resp)
		return
	}
	// TODO unionId+code 校验

	// code-->手机号(wechat)
	var client *wechat.Client
	err = factory.Get(client)
	if err != nil {
		log.Info("no wechat client")
		log.Info(err)
		resp.Code = tool.RespCodeNotFound
		tool.ResponseJSON(ctx, resp)
		return
	}
	var mfaCode = ctx.Params().Get("code")

	wxResp, err := client.GetPhoneNumber(req.Code)
	if err != nil {
		log.Info("get wx phone_number error")
		log.Info(err)
		resp.Code = tool.RespCodeNotFound
		tool.ResponseJSON(ctx, resp)
		return
	}
	phoneNumber := wxResp.PhoneInfo.PurePhoneNumber

	// 绑定微信手机号
	err = serviceBindWechat(mfaCode, req.UnionId, phoneNumber)
	if err != nil {
		log.Info("bind wx phone_number error: %s[%s]=%s", mfaCode, req.UnionId, phoneNumber)
		log.Info(err)
		resp.Code = tool.RespCodeNotFound
		tool.ResponseJSON(ctx, resp)
		return
	}

	resp.Code = tool.RespCodeSuccess
	tool.ResponseJSON(ctx, resp)
}
