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

type processEventApi struct {
	processEventService service.ProcessEvent
}

func (p processEventApi) Save(context echo.Context) error {
	var data v1.PipelineProcessEvent
	body, err := ioutil.ReadAll(context.Request().Body)
	if  err != nil{
		log.Println("Input Error:", err.Error())
		return common.GenerateErrorResponse(context,nil,err.Error())
	}
	if err := json.Unmarshal(body, &data); err != nil {
		return common.GenerateErrorResponse(context,nil,err.Error())
	}
	p.processEventService.Store(data)
	return common.GenerateSuccessResponse(context,"",nil,"Operation Successful!")
}


func NewProcessEventApi(processEventService service.ProcessEvent) api.ProcessEvent {
	return &processEventApi{
		processEventService: processEventService,
	}
}
