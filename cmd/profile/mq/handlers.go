package mq

import (
	"github.com/lishimeng/app-starter/amqp"
	"github.com/lishimeng/profile/cmd/profile/mq/profile"
)

func Handlers() (handlers []amqp.Handler) {
	handlers = append(handlers, &profile.Handler{})
	return
}
