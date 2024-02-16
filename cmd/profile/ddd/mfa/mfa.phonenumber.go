package mfa

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/tool"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/profile/internal/store"
	"github.com/lishimeng/x/util"
)

type PhoneNumberApplyReq struct {
	PhoneNumber string `json:"phone_number"`
}
type PhoneNumberApplyResp struct {
	app.Response
	PhoneNumber string `json:"phone_number"`
}

const mfaPhoneTpl = `mfa_phone_num_%s_%s`

func sendPhoneNumberCode(ctx iris.Context) {
	var err error
	var req PhoneNumberApplyReq
	var resp PhoneNumberApplyResp
	err = ctx.ReadJSON(&req)
	if err != nil {
		log.Info(err)
		resp.Code = tool.RespCodeError
		tool.ResponseJSON(ctx, resp)
		return
	}
	var mfaCode = ctx.Params().Get("code")
	if req.PhoneNumber == "" {
		// TODO 检查手机号格式
		log.Info("unknown phone_number %s", req.PhoneNumber)
		resp.Code = tool.RespCodeNotFound
		tool.ResponseJSON(ctx, resp)
		return
	}
	// 生成验证码
	randCode := util.RandStr(4)
	store.GetManager().
		GetDefaultStore().
		Save(fmt.Sprintf(mfaPhoneTpl, mfaCode, randCode), req.PhoneNumber) // 保存
	// TODO 发送验证码到用户手机

	resp.Code = tool.RespCodeSuccess
	resp.PhoneNumber = req.PhoneNumber
	tool.ResponseJSON(ctx, resp)
	return
}

type PhoneNumberReq struct {
	PhoneNumber string `json:"phone_number"`
	Code        string `json:"code"`
}

func bindPhoneNumber(ctx iris.Context) {
	var err error
	var req PhoneNumberReq
	var resp app.Response
	err = ctx.ReadJSON(&req)
	if err != nil {
		log.Info(err)
		resp.Code = tool.RespCodeError
		tool.ResponseJSON(ctx, resp)
		return
	}
	var mfaCode = ctx.Params().Get("code")
	if req.PhoneNumber == "" {
		// TODO 检查手机号格式
		log.Info("unknown phone_number %s", req.PhoneNumber)
		resp.Code = tool.RespCodeNotFound
		tool.ResponseJSON(ctx, resp)
		return
	}
	if req.Code == "" {
		// TODO 检查手机号格式
		log.Info("unknown code %s", req.Code)
		resp.Code = tool.RespCodeNotFound
		tool.ResponseJSON(ctx, resp)
		return
	}

	phoneNumber, ok := store.GetManager().
		GetDefaultStore().
		Load(fmt.Sprintf(mfaPhoneTpl, mfaCode, req.Code))
	if !ok {
		log.Info("no cache of code %s", req.Code)
		resp.Code = tool.RespCodeNotFound
		tool.ResponseJSON(ctx, resp)
		return
	}
	if phoneNumber != req.PhoneNumber {
		log.Info("not match %s[%s] != %s", req.Code, req.PhoneNumber, phoneNumber)
		resp.Code = tool.RespCodeNotFound
		tool.ResponseJSON(ctx, resp)
		return
	}
	// 绑定手机号
	err = serviceBindPhoneNumber(mfaCode, phoneNumber)
	if err != nil {
		log.Info(err)
		resp.Code = tool.RespCodeError
		tool.ResponseJSON(ctx, resp)
		return
	}

	resp.Code = tool.RespCodeSuccess
	tool.ResponseJSON(ctx, resp)
}
