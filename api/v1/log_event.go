package v1

import (
	"encoding/json"
	"github.com/klovercloud-ci-cd/event-store/api/common"
	v1 "github.com/klovercloud-ci-cd/event-store/core/v1"
	"github.com/klovercloud-ci-cd/event-store/core/v1/api"
	"github.com/klovercloud-ci-cd/event-store/core/v1/service"
	"github.com/labstack/echo/v4"
	"io/ioutil"
	"log"
)

type logEventApi struct {
	logEventService service.LogEvent
}

// Save ... Save log
// @Summary Save log
// @Description Stores logs
// @Tags Log
// @Accept json
// @Produce json
// @Param data body v1.LogEvent true "LogEvent Data"
// @Success 200 {object} common.ResponseDTO
// @Failure 404 {object} common.ResponseDTO
// @Router /api/v1/logs [POST]
func (p logEventApi) Save(context echo.Context) error {
	var data v1.LogEvent
	body, err := ioutil.ReadAll(context.Request().Body)
	if err != nil {
		log.Println("Input Error:", err.Error())
		return common.GenerateErrorResponse(context, nil, err.Error())
	}
	if err := json.Unmarshal(body, &data); err != nil {
		return common.GenerateErrorResponse(context, nil, err.Error())
	}
	p.logEventService.Store(data)
	return common.GenerateSuccessResponse(context, "", nil, "Operation Successful!")
}

// NewLogEventApi returns LogEvent type api
func NewLogEventApi(logEventService service.LogEvent) api.LogEvent {
	return &logEventApi{
		logEventService: logEventService,
	}
}
