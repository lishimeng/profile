package profile

import (
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/persistence"
	"github.com/lishimeng/profile/internal/db/model"
)

func serviceGetProfile(userCode string) (p []model.UserProfile, err error) {

	_, err = app.GetOrm().Context.
		QueryTable(new(model.UserProfile)).
		Filter("UserCode", userCode).
		Limit(10).All(&p)
	if err != nil {
		return
	}
	return
}

func serviceCreateProfile(userCode string) (p model.UserProfile, err error) {

	err = app.GetOrm().Transaction(func(ctx persistence.TxContext) (e error) {

		// 创建用户
		p.Status = app.Enable
		p.UserCode = userCode
		_, e = ctx.Context.Insert(&p)
		if e != nil {
			return
		}

		// 创建mfa
		var mfa model.Mfa
		mfa.MfaCode = p.UserCode
		mfa.MfaType = model.MfaUnknown
		mfa.Status = app.Enable
		_, e = ctx.Context.Insert(&mfa)
		if e != nil {
			return
		}

		return
	})

	return
}
