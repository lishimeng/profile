package setup

import (
	"context"
	"github.com/lishimeng/app-starter/factory"
	"github.com/lishimeng/go-sdk/wechat"
)

func Setup(_ context.Context) (err error) {
	initWechatClient()
	return
}

func initWechatClient() {
	var client *wechat.Client
	// TODO init
	factory.Add(client)
}
