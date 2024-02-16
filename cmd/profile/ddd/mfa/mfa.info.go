package mfa

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/persistence"
	"github.com/lishimeng/app-starter/tool"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/profile/internal/db/model"
)

type Spec struct {
	app.Response
	MfaCode  string            `json:"mfaCode,omitempty"`
	UserCode string            `json:"userCode,omitempty"`
	MfaType  model.MfaCategory `json:"mfaType,omitempty"`
	Items    []SpecItem        `json:"items,omitempty"`
}

type SpecItem struct {
	Sn       string `json:"sn,omitempty"`
	Category string `json:"category,omitempty"`
}

func apiGetMfa(ctx iris.Context) {
	log.Info("get mfa spec")

	var err error
	var resp Spec
	var mfa model.Mfa
	var items []model.MfaItem

	mfaCode := ctx.Params().Get("code")
	err = app.GetOrm().Transaction(func(ctx persistence.TxContext) (e error) {
		e = ctx.Context.QueryTable(new(model.Mfa)).
			Filter("MfaCode", mfaCode).
			Filter("Status", app.Enable).
			One(&mfa)
		if e != nil {
			return
		}
		_, e = ctx.Context.QueryTable(new(model.MfaItem)).
			Filter("MfaCode", mfaCode).
			Filter("Status", app.Enable).
			All(&items)
		if e != nil {
			return
		}
		return
	})
	if err != nil {
		resp.Code = tool.RespCodeNotFound
		tool.ResponseJSON(ctx, resp)
		return
	}

	resp.Code = tool.RespCodeSuccess
	resp.MfaCode = mfa.MfaCode
	resp.UserCode = mfa.UserCode
	resp.MfaType = mfa.MfaType
	if len(items) > 0 {
		for _, item := range items {
			resp.Items = append(resp.Items, SpecItem{
				Sn: item.Sn, Category: string(item.Category),
			})
		}
	}
	tool.ResponseJSON(ctx, resp)
}
