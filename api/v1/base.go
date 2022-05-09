package v1

import (
	"github.com/klovercloud-ci-cd/event-bank/dependency"
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
func FootmarkRouter(g *echo.Group) {

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
	processRouter := NewProcessApi(dependency.GetV1ProcessService(), dependency.GetV1ProcessFootmarkService(), dependency.GetV1LogEventService())
	g.POST("", processRouter.Save, AuthenticationAndAuthorizationHandler)
	g.GET("", processRouter.Get, AuthenticationAndAuthorizationHandler)
	g.GET("/:processId/steps/:step/footmarks", processRouter.GetFootmarksByProcessIdAndStep, AuthenticationAndAuthorizationHandler)
	g.GET("/:processId/logs", processRouter.GetLogsById, AuthenticationAndAuthorizationHandler)
	g.GET("/:processId/steps/:step/footmarks/:footmark/logs", processRouter.GetLogsByProcessIdAndStepAndFootmark, AuthenticationAndAuthorizationHandler)
}

// PipelineRouter api/v1/pipelines router/*
func PipelineRouter(g *echo.Group) {
	pipelineRouter := NewPipelineApi(dependency.GetV1PipelineService(), dependency.GetV1LogEventService(), dependency.GetV1ProcessEventService())
	g.GET("/:processId", pipelineRouter.Get, AuthenticationAndAuthorizationHandler)
	g.GET("/ws", pipelineRouter.GetEvents, AuthenticationAndAuthorizationHandlerForWebSocket)
}

// ProcessLifeCycleRouter api/v1/process_life_cycle_events/* router
func ProcessLifeCycleRouter(g *echo.Group) {
	processLifeCycleEventRouter := NewProcessLifeCycleEventApi(dependency.GetV1ProcessLifeCycleEventService())
	g.POST("", processLifeCycleEventRouter.Save, AuthenticationAndAuthorizationHandler)
	g.GET("", processLifeCycleEventRouter.Pull, AuthenticationAndAuthorizationHandler)
	g.PUT("", processLifeCycleEventRouter.Update, AuthenticationAndAuthorizationHandler)
}
