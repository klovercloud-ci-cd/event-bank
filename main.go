package main

import (
	"github.com/klovercloud-ci-cd/event-bank/api"
	"github.com/klovercloud-ci-cd/event-bank/config"
	_ "github.com/klovercloud-ci-cd/event-bank/docs"
	"github.com/labstack/echo-contrib/jaegertracing"
)

// @title Klovercloud-ci-event-bank API
// @description Klovercloud-ci-event-bank API

func main() {
	e := config.New()
	c := jaegertracing.New(e, nil)
	defer c.Close()
	api.Routes(e)
	e.Logger.Fatal(e.Start(":" + config.ServerPort))
}

//swag init --parseDependency --parseInternal
