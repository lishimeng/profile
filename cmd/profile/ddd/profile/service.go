package profile

import (
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/persistence"
	"github.com/lishimeng/profile/internal/db/model"
	"github.com/lishimeng/x/util"
	"strings"
)

func serviceCreateProfile() (p model.UserProfile, err error) {

	err = app.GetOrm().Transaction(func(ctx persistence.TxContext) (e error) {

		// 创建用户
		p.Status = app.Enable
		p.UserCode = createRandCode()
		_, e = ctx.Context.Insert(&p)
		if e != nil {
			return
		}

		// 创建mfa
		var mfa model.Mfa
		mfa.MfaCode = createRandCode()
		mfa.UserCode = p.UserCode
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

func createRandCode() string {
	code := util.UUIDString()
	code = strings.ToLower(strings.ReplaceAll(code, "-", ""))
	return code
}