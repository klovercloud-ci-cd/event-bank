package api

import "github.com/labstack/echo/v4"

// ProcessEvent Process Event api operations
type ProcessEvent interface {
	Save(context echo.Context) error
}
