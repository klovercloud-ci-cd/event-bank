package api

import "github.com/labstack/echo/v4"

// ProcessEvent Process Event api operations
type ProcessEvent interface {
	Get(context echo.Context) error
	Save(context echo.Context) error
	DequeueByCompanyId(context echo.Context, companyId string) error
	GetByCompanyIdAndProcessId(context echo.Context, companyId string) error
}
