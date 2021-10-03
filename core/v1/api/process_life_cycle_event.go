package api
import "github.com/labstack/echo/v4"

type ProcessLifeCycleEvent interface {
	Save(context echo.Context) error
	Pull(context echo.Context) error
}

