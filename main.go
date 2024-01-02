package main

import (
	"context"
	"fmt"
	"log"

	database "github.com/brianepv1/sv-microservice-login-golang/database"
	"github.com/brianepv1/sv-microservice-login-golang/internal/api"
	"github.com/brianepv1/sv-microservice-login-golang/internal/repository"
	"github.com/brianepv1/sv-microservice-login-golang/internal/service"
	settings "github.com/brianepv1/sv-microservice-login-golang/settings"
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.Provide(
			context.Background,
			settings.New,
			database.New,
			repository.New,
			service.New,
			api.New,
			echo.New,
		),
		fx.Invoke(
			setLifeCycle,

			func(s *settings.Settings) {
				log.Println(s)
			},
		/**
		func(db *sqlx.DB) {
			_, err := db.Query("SELECT * FROM users")
			if err != nil {
				panic(err)
			}
		},
		func(ctx context.Context, serv service.Service) {
			err := serv.RegisterUser(ctx, "my@email.com", "Brian", "validPassword")
			if err != nil {
				panic(err)
			}

			u, err := serv.LoginUser(ctx, "my@email.com", "validPassword")
			if err != nil {
				panic(err)
			}

			if u.FirstName != "Brian" {
				panic("wrong name")
			}
		},
		*/
		),
	)

	app.Run()
}

func setLifeCycle(lc fx.Lifecycle, a *api.API, s *settings.Settings, e *echo.Echo) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			address := fmt.Sprintf(":%s", s.Port)
			go a.Start(e, address)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return nil
		},
	})
}
