package v1

import (
	"github.com/klovercloud-ci/dependency"
	"github.com/labstack/echo/v4"
)

func Router(g *echo.Group) {
	LogEventRouter(g.Group("/logs"))
}

func LogEventRouter(g *echo.Group) {
	logEventRouter := NewLogEventApi(dependency.GetPipelineService())
	g.POST("", logEventRouter.Save, AuthenticationAndAuthorizationHandler)
	g.GET("/:processId", logEventRouter.GetByProcessId,AuthenticationAndAuthorizationHandler)
}
