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
	"strconv"
)

type processLifeCycleEventApi struct {
	processLifeCycleEventService service.ProcessLifeCycleEvent
}

func (p processLifeCycleEventApi) Pull(context echo.Context) error {
	agentName:=context.QueryParam("agent")
	count,_:=strconv.ParseInt(context.QueryParam("count"), 10, 64)
	steptype:=context.QueryParam("step_type")
	if steptype!=""{
		return common.GenerateSuccessResponse(context,p.processLifeCycleEventService.PullNonInitializedAndAutoTriggerEnabledEventsByStepType(count,steptype),nil,"")
	}
	return common.GenerateSuccessResponse(context,p.processLifeCycleEventService.PullPausedAndAutoTriggerEnabledResourcesByAgentName(count,agentName),nil,"")
}

func (p processLifeCycleEventApi) Save(context echo.Context) error {
	type ProcessLifeCycleEventList struct {
		Events [] v1.ProcessLifeCycleEvent `bson:"events" json :"events"`
	}
	var data ProcessLifeCycleEventList
	body, err := ioutil.ReadAll(context.Request().Body)
	if  err != nil{
		log.Println("Input Error:", err.Error())
		return common.GenerateErrorResponse(context,nil,err.Error())
	}
	if err := json.Unmarshal(body, &data); err != nil {
		return common.GenerateErrorResponse(context,nil,err.Error())
	}
	p.processLifeCycleEventService.Store(data.Events)
	return common.GenerateSuccessResponse(context,"",nil,"Operation Successful!")
}


func NewProcessLifeCycleEventApi(processLifeCycleEventService service.ProcessLifeCycleEvent) api.ProcessLifeCycleEvent {
	return &processLifeCycleEventApi{
		processLifeCycleEventService: processLifeCycleEventService,
	}
}
