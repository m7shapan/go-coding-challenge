package main

import (
	"challenge/pkg/config"
	"challenge/pkg/db"
	"challenge/server"

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

	server := server.Server{
		AppConfig: appConfig,
		DB:        dbClient,
	}

	server.Start()
}
