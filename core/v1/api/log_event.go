package api

import (
	"github.com/labstack/echo/v4"
)

// LogEvent Log Event api operations
type LogEvent interface {
	Save(context echo.Context) error
}
