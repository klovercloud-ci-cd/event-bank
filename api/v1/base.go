package v1

import (
	"github.com/klovercloud-ci/dependency"
	"github.com/labstack/echo/v4"
)

func Router(g *echo.Group) {
	LogEventRouter(g.Group("/logs"))
	PipelineRouter(g.Group("/pipelines"))
	ProcessEventRouter(g.Group("/processes_events"))
	ProcessRouter(g.Group("/processes"))

}

func LogEventRouter(g *echo.Group) {
	logEventRouter := NewLogEventApi(dependency.GetLogEventService())
	g.POST("", logEventRouter.Save, AuthenticationAndAuthorizationHandler)
}

func ProcessEventRouter(g *echo.Group) {
	processEventRouter := NewProcessEventApi(dependency.GetProcessEventService())
	g.POST("", processEventRouter.Save, AuthenticationAndAuthorizationHandler)
}
func ProcessRouter(g *echo.Group) {
	processRouter := NewProcessApi(dependency.GetProcessService())
	g.POST("", processRouter.Save, AuthenticationAndAuthorizationHandler)
	g.GET("",processRouter.GetByCompanyIdAndRepositoryIdAndAppName,AuthenticationAndAuthorizationHandler)
}

func PipelineRouter(g *echo.Group) {
	pipelineRouter := NewPipelineApi(dependency.GetLogEventService(),dependency.GetProcessEventService())
	g.GET("/:processId",pipelineRouter.GetLogs,AuthenticationAndAuthorizationHandler)
	g.GET("/ws",pipelineRouter.GetEvents,AuthenticationAndAuthorizationHandler)
}
