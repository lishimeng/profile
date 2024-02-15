package main

import (
	"context"
	"fmt"
	"github.com/lishimeng/app-starter"
	etc2 "github.com/lishimeng/app-starter/etc"
	"github.com/lishimeng/app-starter/factory"
	"github.com/lishimeng/app-starter/persistence"
	"github.com/lishimeng/app-starter/token"
	"github.com/lishimeng/hufu/cmd/oauth/ddd"
	"github.com/lishimeng/hufu/cmd/oauth/static"
	"github.com/lishimeng/hufu/internal/db/model"
	"github.com/lishimeng/hufu/internal/etc"
	"net/http"
	"time"
)
import _ "github.com/lib/pq"

func main() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	err := _main()
	if err != nil {
		fmt.Println(err)
	}
	time.Sleep(time.Millisecond * 200)
}

func _main() (err error) {
	configName := "config"

	application := app.New()

	err = application.Start(func(ctx context.Context, builder *app.ApplicationBuilder) error {

		var err error
		err = builder.LoadConfig(&etc.Config, func(loader etc2.Loader) {
			loader.SetFileSearcher(configName, ".").SetEnvPrefix("").SetEnvSearcher()
		})
		if err != nil {
			return err
		}
		dbConfig := persistence.PostgresConfig{
			UserName:  etc.Config.Db.User,
			Password:  etc.Config.Db.Password,
			Host:      etc.Config.Db.Host,
			Port:      etc.Config.Db.Port,
			DbName:    etc.Config.Db.Database,
			InitDb:    true,
			AliasName: "default",
			SSL:       etc.Config.Db.Ssl,
		}

		builder.EnableDatabase(dbConfig.Build(),
			model.Tables()...).
			SetWebLogLevel("debug").
			EnableOrmLog().
			EnableStaticWeb(func() http.FileSystem {
				return http.FS(static.Static)
			}).
			EnableWeb(etc.Config.Web.Listen, ddd.Route).
			EnableTokenValidator(func(injectFunc app.TokenValidatorInjectFunc) {
				key := []byte(etc.Config.Token.Key)
				provider := token.NewJwtProvider(
					token.WithIssuer(etc.Config.Token.Issuer),
					token.WithAlg(etc.Config.Token.Alg),
					token.WithKey(key, key))
				session := token.NewLocalStorage(provider)
				injectFunc(session)
				factory.Add(provider)
			}).
			PrintVersion()
		return err
	})
	return
}
