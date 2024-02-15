package mfa

import (
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/persistence"
	model "github.com/lishimeng/profile/internal/db"
)

func serviceBindPhoneNumber(mfa string, phoneNumber string) (err error) {

	err = app.GetOrm().Transaction(func(ctx persistence.TxContext) error {
		var list []model.MfaItem
		var item model.MfaItem
		var e error
		e = ctx.Context.QueryTable(new(model.MfaItem)).
			Filter("Category", model.MfaPhoneNumber).
			Filter("MfaCode", mfa).
			One(&list)
		if e != nil {
			return e
		}
		if len(list) > 0 {
			item = list[0]
		}
		item.Category = model.MfaPhoneNumber
		item.MfaCode = mfa
		item.Sn = phoneNumber
		item.Status = app.Enable

		_, e = ctx.Context.InsertOrUpdate(&item)
		return e
	})
	// 增加一条
	return
}
