package main

import (
	"github.com/klovercloud-ci-cd/event-bank/api"
	"github.com/klovercloud-ci-cd/event-bank/config"
	"github.com/klovercloud-ci-cd/event-bank/dependency"
	_ "github.com/klovercloud-ci-cd/event-bank/docs"
	"github.com/labstack/echo-contrib/jaegertracing"
	"io"
	"log"
	"time"
)

// @title Klovercloud-ci-event-bank API
// @description Klovercloud-ci-event-bank API
func main() {
	e := config.New()
	c := jaegertracing.New(e, nil)
	defer func(c io.Closer) {
		err := c.Close()
		if err != nil {
			panic(err)
		}
	}(c)
	go UpdatePipelineStepStatus()
	api.Routes(e)
	e.Logger.Fatal(e.Start(":" + config.ServerPort))
}

//
func UpdatePipelineStepStatus() {
	log.Println("Updating pipeline step status, time:", time.Now().UTC())
	p := dependency.GetV1ProcessLifeCycleEventService()
	p.UpdateStatusesByTime(time.Now().UTC().Add(time.Minute * -20))
	time.Sleep(time.Minute * 20)
	UpdatePipelineStepStatus()
}

//swag init --parseDependency --parseInternal
