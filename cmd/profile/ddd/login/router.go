package login

import "github.com/kataras/iris/v12"

func Route(root iris.Party) {
	root.Post("/wx/mini_code", wechatLogin)
	root.Post("/wx/unionid", unionIdLogin)
}
