package ddd

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/profile/cmd/profile/ddd/mfa"
	"github.com/lishimeng/profile/cmd/profile/ddd/profile"
	"github.com/lishimeng/profile/cmd/profile/ddd/user"
)

func Route(app *iris.Application) {

	root := app.Party("/api")
	mfa.Route(root.Party("/mfa"))
	profile.Route(root.Party("/profile"))
	user.Route(root.Party("/user"))
}
