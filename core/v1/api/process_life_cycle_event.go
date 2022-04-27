package api

import (
	"github.com/labstack/echo/v4"
)

// ProcessLifeCycleEvent Process Life Cycle Event api operations
type ProcessLifeCycleEvent interface {
	Save(context echo.Context) error
	Pull(context echo.Context) error
	Update(context echo.Context) error
}
