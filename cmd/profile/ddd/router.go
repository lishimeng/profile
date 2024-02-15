package ddd

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/profile/cmd/profile/ddd/mfa"
	"github.com/lishimeng/profile/cmd/profile/ddd/profile"
	"github.com/lishimeng/profile/cmd/profile/ddd/user"
)

func Route(app *iris.Application) {

	mfa.Route(app.Party("/mfa"))
	profile.Route(app.Party("/profile"))
	user.Route(app.Party("/user"))
}
