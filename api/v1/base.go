package v1

import (
	"github.com/klovercloud-ci-cd/event-store/dependency"
	"github.com/labstack/echo/v4"
)

// Router api/v1 base router
func Router(g *echo.Group) {
	LogEventRouter(g.Group("/logs"))
	PipelineRouter(g.Group("/pipelines"))
	ProcessEventRouter(g.Group("/processes_events"))
	ProcessRouter(g.Group("/processes"))
	ProcessLifeCycleRouter(g.Group("/process_life_cycle_events"))

}

// LogEventRouter api/v1/logs/* router
func LogEventRouter(g *echo.Group) {
	logEventRouter := NewLogEventApi(dependency.GetV1LogEventService())
	g.POST("", logEventRouter.Save, AuthenticationAndAuthorizationHandler)
}

// ProcessEventRouter api/v1/processes_events router/*
func ProcessEventRouter(g *echo.Group) {
	processEventRouter := NewProcessEventApi(dependency.GetV1ProcessEventService())
	g.POST("", processEventRouter.Save, AuthenticationAndAuthorizationHandler)
}

// ProcessRouter api/v1/processes router/*
func ProcessRouter(g *echo.Group) {
	processRouter := NewProcessApi(dependency.GetV1ProcessService())
	g.POST("", processRouter.Save, AuthenticationAndAuthorizationHandler)
	g.GET("", processRouter.Get, AuthenticationAndAuthorizationHandler)
}

// PipelineRouter api/v1/pipelines router/*
func PipelineRouter(g *echo.Group) {
	pipelineRouter := NewPipelineApi(dependency.GetV1LogEventService(), dependency.GetV1ProcessEventService())
	g.GET("/:processId", pipelineRouter.GetLogs, AuthenticationAndAuthorizationHandler)
	g.GET("/ws", pipelineRouter.GetEvents, AuthenticationAndAuthorizationHandler)
}

// ProcessLifeCycleRouter api/v1/process_life_cycle_events/* router
func ProcessLifeCycleRouter(g *echo.Group) {
	processLifeCycleEventRouter := NewProcessLifeCycleEventApi(dependency.GetV1ProcessLifeCycleEventService())
	g.POST("", processLifeCycleEventRouter.Save, AuthenticationAndAuthorizationHandler)
	g.GET("", processLifeCycleEventRouter.Pull, AuthenticationAndAuthorizationHandler)
}
