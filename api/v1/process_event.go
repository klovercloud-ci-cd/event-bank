package v1

import (
	"encoding/json"
	"github.com/klovercloud-ci-cd/event-bank/api/common"
	v1 "github.com/klovercloud-ci-cd/event-bank/core/v1"
	"github.com/klovercloud-ci-cd/event-bank/core/v1/api"
	"github.com/klovercloud-ci-cd/event-bank/core/v1/service"
	"github.com/klovercloud-ci-cd/event-bank/enums"
	"github.com/labstack/echo/v4"
	"io/ioutil"
	"strconv"
)

type processEventApi struct {
	processEventService service.ProcessEvent
}

// Get ... Get Proccess Event By Company Id And Process Id
// @Summary Get Proccess Event By Company Id And Process Id
// @Description Get Proccess Event By Company Id And Process Id
// @Tags ProcessEvent
// @Accept json
// @Produce json
// @Param scope query string false "scope [notification]"
// @Param companyId query string true "Company Id"
// @Param processId query string false "Process Id when scope is notification [Optional]"
// @Param page query int64 false "Page number"
// @Param limit query int64 false "Record count"
// @Success 200 {object} common.ResponseDTO
// @Failure 404 {object} common.ResponseDTO
// @Router /api/v1/processes_events [GET]
func (p processEventApi) Get(context echo.Context) error {
	companyId := context.QueryParam("companyId")
	scope := context.QueryParam("scope")
	if scope == "notification" {
		return p.GetByCompanyIdAndProcessId(context, companyId)
	}
	return p.DequeueByCompanyId(context, companyId)
}

func (p processEventApi) GetByCompanyIdAndProcessId(context echo.Context, companyId string) error {
	processId := context.QueryParam("processId")
	option := getProcessQueryOption(context)
	data := p.processEventService.GetByCompanyIdAndProcessId(companyId, processId, option)
	return common.GenerateSuccessResponse(context, data, nil, "Operation Successful!")
}

func (p processEventApi) getProcessEventQueryOption(context echo.Context) v1.ProcessQueryOption {
	option := v1.ProcessQueryOption{}
	page := context.QueryParam("page")
	limit := context.QueryParam("limit")
	if page == "" {
		option.Pagination.Page = enums.DEFAULT_PAGE
		option.Pagination.Limit = enums.DEFAULT_PAGE_LIMIT
	} else {
		option.Pagination.Page, _ = strconv.ParseInt(page, 10, 64)
		option.Pagination.Limit, _ = strconv.ParseInt(limit, 10, 64)
	}
	return option
}

func (p processEventApi) DequeueByCompanyId(context echo.Context, companyId string) error {
	data := p.processEventService.DequeueByCompanyId(companyId)
	return common.GenerateSuccessResponse(context, data, nil, "Operation Successful!")
}

// Save ... Save Pipeline process event
// @Summary Save Pipeline process event
// @Description Stores Pipeline process event
// @Tags ProcessEvent
// @Accept json
// @Produce json
// @Param data body v1.PipelineProcessEvent true "PipelineProcessEvent Data"
// @Success 200 {object} common.ResponseDTO
// @Failure 404 {object} common.ResponseDTO
// @Router /api/v1/processes_events [POST]
func (p processEventApi) Save(context echo.Context) error {
	var data v1.PipelineProcessEvent
	body, err := ioutil.ReadAll(context.Request().Body)
	if err != nil {
		return common.GenerateErrorResponse(context, nil, err.Error())
	}
	if err := json.Unmarshal(body, &data); err != nil {
		return common.GenerateErrorResponse(context, nil, err.Error())
	}
	p.processEventService.Store(data)
	return common.GenerateSuccessResponse(context, "", nil, "Operation Successful!")
}

// NewProcessEventApi returns ProcessEvent type api
func NewProcessEventApi(processEventService service.ProcessEvent) api.ProcessEvent {
	return &processEventApi{
		processEventService: processEventService,
	}
}
