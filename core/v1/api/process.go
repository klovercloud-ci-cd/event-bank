package api

import "github.com/labstack/echo/v4"

// Process Process api operations
type Process interface {
	Save(context echo.Context) error
	Get(context echo.Context) error
	GetFootmarksByProcessIdAndStep(context echo.Context) error
}
