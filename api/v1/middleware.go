package v1

import (
	"github.com/klovercloud-ci-cd/event-bank/api/common"
	"github.com/klovercloud-ci-cd/event-bank/config"
	"github.com/klovercloud-ci-cd/event-bank/dependency"
	"github.com/labstack/echo/v4"
)

// AuthenticationAndAuthorizationHandler handle user authentication and authorization here.
func AuthenticationAndAuthorizationHandler(handler echo.HandlerFunc) echo.HandlerFunc {
	return func(context echo.Context) (err error) {
		if config.EnableAuthentication {
			token := context.Request().Header.Get("token")
			if token == "" {
				return common.GenerateErrorResponse(context, "[ERROR]: Invalid token!", "Failed to parse token!")
			}
			res, _ := dependency.GetV1JwtService().ValidateToken(token)
			if !res {
				return common.GenerateErrorResponse(context, "[ERROR]: Invalid token!", "Please provide a valid token!")
			}
		}
		return handler(context)
	}
}

// AuthenticationAndAuthorizationHandlerForWebSocket handle user authentication and authorization here.
func AuthenticationAndAuthorizationHandlerForWebSocket(handler echo.HandlerFunc) echo.HandlerFunc {
	return func(context echo.Context) (err error) {
		if config.EnableAuthentication {
			token := context.QueryParam("token")
			if token == "" {
				return common.GenerateErrorResponse(context, "[ERROR]: Invalid token!", "Failed to parse token!")
			}
			res, _ := dependency.GetV1JwtService().ValidateToken(token)
			if !res {
				return common.GenerateErrorResponse(context, "[ERROR]: Invalid token!", "Please provide a valid token!")
			}
		}
		return handler(context)
	}
}
