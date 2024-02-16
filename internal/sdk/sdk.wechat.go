package sdk

import (
	"github.com/lishimeng/app-starter/factory"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/go-sdk/wechat"
)

func CreateWechatClient(config WechatConfig) {
	log.Info("init wechat client")
	var client *wechat.Client
	client = wechat.New(config.AppID, config.AppSecret)
	factory.Add(client)
}
