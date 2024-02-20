package login

import (
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/persistence"
	"github.com/lishimeng/profile/internal/data"
	"github.com/lishimeng/profile/internal/db/model"
)

func svsGetUnionId(unionId string) (item model.MfaItem, err error) {

	err = app.GetOrm().Context.
		QueryTable(new(model.MfaItem)).
		Filter("Status", app.Enable).
		Filter("Category", model.MfaWechat).
		Filter("Sn", unionId).
		One(&item)

	return
}

func svsBindNewWxUser(unionId string) (p model.UserProfile, item model.MfaItem, err error) {
	err = app.GetOrm().Transaction(func(ctx persistence.TxContext) (e error) {
		p.UserCode = data.CreateRandCode()
		p.Status = app.Enable
		_, e = ctx.Context.Insert(&p)
		if e != nil {
			return
		}
		item.MfaCode = p.UserCode
		item.Sn = unionId
		item.Category = model.MfaWechat
		item.Status = app.Enable
		_, e = ctx.Context.Insert(&item)
		return
	})
	return
}
