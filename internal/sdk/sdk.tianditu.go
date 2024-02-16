package sdk

import (
	"github.com/lishimeng/app-starter/factory"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/go-sdk/tianditu"
)

func CreateTianditu(config TiandituConfig) {
	log.Info("init tianditu sdk")
	tClient := tianditu.NewClient(tianditu.WithKey(config.Key))
	factory.Add(&tClient)
}
