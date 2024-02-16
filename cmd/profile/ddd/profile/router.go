package profile

import "github.com/kataras/iris/v12"

func Route(root iris.Party) {
	root.Post("/", createProfile)
	root.Get("/{code}", getProfileSpec)
}
