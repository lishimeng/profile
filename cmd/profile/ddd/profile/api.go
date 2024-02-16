package profile

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/tool"
	"github.com/lishimeng/go-log"
)

type CreateProfileResp struct {
	app.Response
	Id string `json:"id,omitempty"`
}

func createProfile(ctx iris.Context) {

	log.Info("create user profile")
	var err error
	var resp CreateProfileResp
	p, err := serviceCreateProfile()
	if err != nil {
		log.Info("createProfile fail")
		log.Info(err)
		resp.Code = tool.RespCodeError
		tool.ResponseJSON(ctx, resp)
		return
	}

	resp.Code = tool.RespCodeSuccess
	resp.Id = p.UserCode
	tool.ResponseJSON(ctx, resp)
}
