package main

import (
	"github.com/klovercloud-ci-cd/event-bank/api"
	"github.com/klovercloud-ci-cd/event-bank/config"
	_ "github.com/klovercloud-ci-cd/event-bank/docs"
)

// @title Klovercloud-ci-event-bank API
// @description Klovercloud-ci-event-bank API

func main() {
	e := config.New()
	api.Routes(e)
	e.Logger.Fatal(e.Start(":" + config.ServerPort))
}

//swag init --parseDependency --parseInternal
