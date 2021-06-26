package main

import (
	"challenge/controllers"
	"challenge/pkg/config"
	"challenge/pkg/db"
	"challenge/repositories"
	"challenge/services"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

func main() {
	log.Info("Loading app config")
	appConfig, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	log.Info("Try to connect to DB")
	dbClient, err := db.Connect(appConfig.DB)
	if err != nil {
		log.Fatal(err)
	}

	log.Info("Init app backbone")
	factRepository := repositories.NewFactRepository(dbClient)
	factService := services.NewFactService(factRepository)
	factController := controllers.NewFactController(factService)

	e := echo.New()
	e.HideBanner = true

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	v1 := e.Group("/api/v1")

	v1.GET("/facts", factController.GetFacts)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", appConfig.AppPort)))
}
