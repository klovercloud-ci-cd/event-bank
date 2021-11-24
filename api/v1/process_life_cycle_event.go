package v1

import (
	"encoding/json"
	"github.com/klovercloud-ci-cd/event-bank/api/common"
	v1 "github.com/klovercloud-ci-cd/event-bank/core/v1"
	"github.com/klovercloud-ci-cd/event-bank/core/v1/api"
	"github.com/klovercloud-ci-cd/event-bank/core/v1/service"
	"github.com/labstack/echo/v4"
	"io/ioutil"
	"log"
	"strconv"
)

type processLifeCycleEventApi struct {
	processLifeCycleEventService service.ProcessLifeCycleEvent
}

// Pull... Pull Steps
// @Summary Pull Steps
// @Description Pulls auto trigger enabled steps
// @Tags ProcessLifeCycle
// @Produce json
// @Param agent path string true "Agen name"
// @Param count path int64 true "Pull size"
// @Param step_type path string false "Step type [BUILD, DEPLOY]"
// @Success 200 {object} common.ResponseDTO{data=[]v1.ProcessLifeCycleEvent}
// @Router /api/v1/process_life_cycle_events [GET]
func (p processLifeCycleEventApi) Pull(context echo.Context) error {
	agentName := context.QueryParam("agent")
	count, _ := strconv.ParseInt(context.QueryParam("count"), 10, 64)
	steptype := context.QueryParam("step_type")
	if steptype != "" {
		return common.GenerateSuccessResponse(context, p.processLifeCycleEventService.PullNonInitializedAndAutoTriggerEnabledEventsByStepType(count, steptype), nil, "")
	}
	return common.GenerateSuccessResponse(context, p.processLifeCycleEventService.PullPausedAndAutoTriggerEnabledResourcesByAgentName(count, agentName), nil, "")
}

// Save ... Save process lifecycle event
// @Summary Save process lifecycle event
// @Description Stores process lifecycle event
// @Tags ProcessLifeCycle
// @Accept json
// @Produce json
// @Param data body v1.ProcessLifeCycleEventList true "ProcessLifeCycleEventList Data"
// @Success 200 {object} common.ResponseDTO
// @Failure 404 {object} common.ResponseDTO
// @Router /api/v1/process_life_cycle_events [POST]
func (p processLifeCycleEventApi) Save(context echo.Context) error {

	var data v1.ProcessLifeCycleEventList
	body, err := ioutil.ReadAll(context.Request().Body)
	if err != nil {
		log.Println("Input Error:", err.Error())
		return common.GenerateErrorResponse(context, nil, err.Error())
	}
	if err := json.Unmarshal(body, &data); err != nil {
		return common.GenerateErrorResponse(context, nil, err.Error())
	}
	p.processLifeCycleEventService.Store(data.Events)
	return common.GenerateSuccessResponse(context, "", nil, "Operation Successful!")
}

// NewProcessLifeCycleEventApi returns ProcessLifeCycleEvent type api
func NewProcessLifeCycleEventApi(processLifeCycleEventService service.ProcessLifeCycleEvent) api.ProcessLifeCycleEvent {
	return &processLifeCycleEventApi{
		processLifeCycleEventService: processLifeCycleEventService,
	}
}
