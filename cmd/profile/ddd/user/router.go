package user

import "github.com/kataras/iris/v12"

func Route(root iris.Party) {
	root.Post("/", addUser)
	root.Post("/wechat", addUser)
}
