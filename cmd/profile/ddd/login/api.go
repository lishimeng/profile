package login

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/factory"
	"github.com/lishimeng/app-starter/tool"
	"github.com/lishimeng/go-sdk/wechat"
	"github.com/lishimeng/profile/internal/data"
)

type WxJsCodeReq struct {
	Code string
}
type WxLoginResp struct {
	app.Response
	data.UserToken
}

func wechatLogin(ctx iris.Context) {
	// 微信js_code
	// 微信openid
	// 微信session_key

	var err error
	var req WxJsCodeReq
	var resp WxLoginResp
	err = ctx.ReadJSON(&req)
	if err != nil {
		resp.Code = tool.RespCodeError
		tool.ResponseJSON(ctx, resp)
		return
	}

	if len(req.Code) == 0 {
		resp.Code = tool.RespCodeNotFound
		tool.ResponseJSON(ctx, resp)
		return
	}

	// TODO cache重复请求(code), 至少2小时

	var wxClient *wechat.Client
	err = factory.Get(wxClient)
	if err != nil {
		resp.Code = tool.RespCodeError
		tool.ResponseJSON(ctx, resp)
		return
	}
	result, err := wxClient.JsCode2Session(req.Code)
	if len(req.Code) == 0 {
		resp.Code = tool.RespCodeNotFound
		tool.ResponseJSON(ctx, resp)
		return
	}
	if result.ErrCode != 0 {
		// TODO
		resp.Code = tool.RespCodeNotFound
		tool.ResponseJSON(ctx, resp)
		return
	}
	// 换取union_id
	var unionId = result.UnionId
	mfaItem, err := svsGetUnionId(unionId)
	if err != nil {
		// 创建
		// create profile
		// bind wechat(union_id)
		_, mfaItem, err = svsBindNewWxUser(unionId)
		if err != nil {
			resp.Code = tool.RespCodeError
			tool.ResponseJSON(ctx, resp)
			return
		}
	}

	resp.Code = tool.RespCodeSuccess
	resp.Uid = mfaItem.MfaCode
	tool.ResponseJSON(ctx, resp)
	// TODO token
}

func unionIdLogin(ctx iris.Context) {
	// union_id不存在, 返回404

	var err error
	var req WxJsCodeReq
	var resp WxLoginResp
	err = ctx.ReadJSON(&req)
	if err != nil {
		resp.Code = tool.RespCodeError
		tool.ResponseJSON(ctx, resp)
		return
	}

	if len(req.Code) == 0 {
		resp.Code = tool.RespCodeNotFound
		tool.ResponseJSON(ctx, resp)
		return
	}

	// TODO cache重复请求(code), 至少2小时

	var wxClient *wechat.Client
	err = factory.Get(wxClient)
	if err != nil {
		resp.Code = tool.RespCodeError
		tool.ResponseJSON(ctx, resp)
		return
	}
	result, err := wxClient.JsCode2Session(req.Code)
	if len(req.Code) == 0 {
		resp.Code = tool.RespCodeNotFound
		tool.ResponseJSON(ctx, resp)
		return
	}
	if result.ErrCode != 0 {
		// TODO
		resp.Code = tool.RespCodeNotFound
		tool.ResponseJSON(ctx, resp)
		return
	}
	// 换取union_id
	var unionId = result.UnionId
	mfaItem, err := svsGetUnionId(unionId)
	if err != nil {
		resp.Code = tool.RespCodeNotFound
		tool.ResponseJSON(ctx, resp)
		return
	}
	resp.Code = tool.RespCodeSuccess
	resp.Uid = mfaItem.MfaCode
	tool.ResponseJSON(ctx, resp)
	// TODO token
}
