package api

import "github.com/labstack/echo/v4"

type Process interface {
	Save(context echo.Context) error
	GetByCompanyIdAndRepositoryIdAndAppName(context echo.Context) error
}

