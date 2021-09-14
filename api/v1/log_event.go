package v1

import (
	"encoding/json"
	"github.com/klovercloud-ci/api/common"
	v1 "github.com/klovercloud-ci/core/v1"
	"github.com/klovercloud-ci/core/v1/api"
	"github.com/klovercloud-ci/core/v1/service"
	"github.com/labstack/echo/v4"
	"io/ioutil"
	"log"
)

type logEventApi struct {
	logEventService service.LogEvent
}

func (p logEventApi) Save(context echo.Context) error {
	var data v1.LogEvent
	body, err := ioutil.ReadAll(context.Request().Body)
	if  err != nil{
		log.Println("Input Error:", err.Error())
		return common.GenerateErrorResponse(context,nil,err.Error())
	}
	if err := json.Unmarshal(body, &data); err != nil {
		return common.GenerateErrorResponse(context,nil,err.Error())
	}
	p.logEventService.Store(data)
	return common.GenerateSuccessResponse(context,"",nil,"Operation Successful!")
}


func NewLogEventApi(logEventService service.LogEvent) api.LogEvent {
	return &logEventApi{
		logEventService: logEventService,
	}
}
