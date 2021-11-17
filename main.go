package main

import (
	"github.com/klovercloud-ci-cd/event-bank/api"
	"github.com/klovercloud-ci-cd/event-bank/config"
	_ "github.com/klovercloud-ci-cd/event-bank/docs"
)

// @title Klovercloud-ci-event-store API
// @description Klovercloud-ci-event-store API

func main() {
	e := config.New()
	api.Routes(e)
	e.Logger.Fatal(e.Start(":" + config.ServerPort))
}

//swag init --parseDependency --parseInternal
