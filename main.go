package main

import (
	"github.com/klovercloud-ci-cd/klovercloud-ci-event-store/api"
	"github.com/klovercloud-ci-cd/klovercloud-ci-event-store/config"
	_ "github.com/klovercloud-ci-cd/klovercloud-ci-event-store/docs"
)

// @title Klovercloud-ci-event-store API
// @description Klovercloud-ci-event-store API

func main() {
	e := config.New()
	api.Routes(e)
	e.Logger.Fatal(e.Start(":" + config.ServerPort))
}

//swag init --parseDependency --parseInternal
