package api

import "github.com/labstack/echo/v4"

type ProcessEvent interface {
	Save(context echo.Context) error
}
