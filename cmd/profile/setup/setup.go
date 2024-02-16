package setup

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/factory"
	"github.com/lishimeng/profile/internal/db/model"
	"github.com/lishimeng/profile/internal/sdk"
	"github.com/lishimeng/profile/internal/store"
)

func Setup(_ context.Context) (err error) {
	initStoreManager()
	err = initSdkClient()
	if err != nil {
		return
	}
	err = initNotifySdk()
	if err != nil {
		return
	}
	return
}

func initStoreManager() {
	m := store.NewStoreManager()
	factory.Add(&m)
}

func initSdkClient() (err error) {

	var config model.SdkConfig
	err = app.GetOrm().Context.QueryTable(new(model.SdkConfig)).One(&config)
	if err != nil {
		return
	}

	// 初始化微信sdk客户端
	var wx = config.Wechat
	bs, err := base64.StdEncoding.DecodeString(wx)
	if err != nil {
		return
	}
	var wxConfig sdk.WechatConfig
	err = json.Unmarshal(bs, &wxConfig)
	if err != nil {
		return
	}
	sdk.CreateWechatClient(wxConfig)

	// 初始化天地图sdk客户端
	var t = config.Tianditu
	bs, err = base64.StdEncoding.DecodeString(t)
	if err != nil {
		return
	}
	var tdConfig sdk.TiandituConfig
	err = json.Unmarshal(bs, &tdConfig)
	if err != nil {
		return
	}
	sdk.CreateTianditu(tdConfig)
	return
}

func initNotifySdk() (err error) {
	// TODO 初始化notify sdk
	return
}
