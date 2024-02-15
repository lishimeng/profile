package ddd

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/profile/cmd/profile/ddd/mfa"
)

func Route(app *iris.Application) {

	mfa.Route(app.Party("/mfa"))
}
