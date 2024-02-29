package profile

import (
	"encoding/json"
	"github.com/lishimeng/app-starter/amqp/rabbit"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/profile/cmd/profile/models"
)

type Handler struct {
}

func (h *Handler) Subscribe(payload interface{}, _ rabbit.TxHandler, _ rabbit.ServerContext) {
	var err error
	s := payload.([]byte)
	var req models.BasicProfileCreate
	err = json.Unmarshal(s, &req)
	if err != nil {
		log.Info(err)
		return
	}
	p, err := serviceCreateProfile(req.Uid)
	if err != nil {
		log.Info("createProfile fail")
		log.Info(err)
		return
	}
	log.Info("success create profile: %s", p.UserCode)
}

func (h *Handler) Router() rabbit.Route {
	return rabbit.Route{
		Queue: "mq_profile_create",
	}
}
