package main

import (
	"github.com/klovercloud-ci/api"
	"github.com/klovercloud-ci/config"
)
// @title Klovercloud-ci-event-store API
// @description Klovercloud-ci-event-store API

func main(){
	e:=config.New()
	api.Routes(e)
	e.Logger.Fatal(e.Start(":"+ config.ServerPort))
}
//swag init --parseDependency --parseInternal