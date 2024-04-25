package main

import (
	"github.com/labstack/echo/v4"
	"github.com/zer0day88/brick-test/infra/db/porstgres"
	"github.com/zer0day88/brick-test/internal/app/domain/entities"
	"github.com/zer0day88/brick-test/internal/app/route"
	"github.com/zer0day88/brick-test/pkg/config"
	"github.com/zer0day88/brick-test/pkg/environment"
	"github.com/zer0day88/brick-test/pkg/logger"
)

func main() {

	log := logger.New()

	config.Load(log)

	db, err := postgres.Init()

	if err != nil {
		log.Fatal().Err(err).Send()
	}

	if config.Key.Environment == environment.Development {
		errMigrate := db.AutoMigrate(&entities.Transfer{})
		if errMigrate != nil {
			log.Fatal().Err(errMigrate).Send()
		}
	}

	e := echo.New()
	//e.Use(middleware.Recover())

	route.InitRoute(e, db, log)

	err = e.Start(":" + config.Key.Port)
	if err != nil {
		log.Fatal().Msg("Failed to start server")
	}

}
