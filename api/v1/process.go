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

type processApi struct {
	processService service.Process
}

// Save ... Save process
// @Summary Save process
// @Description Stores process
// @Tags Process
// @Accept json
// @Produce json
// @Param data body v1.Process true "Process Data"
// @Success 200 {object} common.ResponseDTO
// @Failure 404 {object} common.ResponseDTO
// @Router /api/v1/processes [POST]
func (p processApi) Save(context echo.Context) error {
	var data v1.Process
	body, err := ioutil.ReadAll(context.Request().Body)
	if err != nil {
		log.Println("Input Error:", err.Error())
		return common.GenerateErrorResponse(context, nil, err.Error())
	}
	if err := json.Unmarshal(body, &data); err != nil {
		return common.GenerateErrorResponse(context, nil, err.Error())
	}
	p.processService.Store(data)
	return common.GenerateSuccessResponse(context, "", nil, "Operation Successful!")
}

// Get... Get Process List or count process
// @Summary Get Process List or count process
// @Description Get Process List or count process
// @Tags Process
// @Produce json
// @Param companyId query string true "Company Id"
// @Param repositoryId query string false "Repository Id"
// @Param appId query string true "App Id"
// @Param operation query string false "Operation[countTodaysProcessByCompanyId]"
// @Success 200 {object} common.ResponseDTO
// @Router /api/v1/processes [GET]
func (p processApi) Get(context echo.Context) error {
	companyId := context.QueryParam("companyId")
	repositoryId := context.QueryParam("repositoryId")
	appId := context.QueryParam("appId")
	operation := context.QueryParam("operation")
	if operation == "countTodaysProcessByCompanyId" {
		return common.GenerateSuccessResponse(context, p.processService.CountTodaysRanProcessByCompanyId(companyId), nil, "")
	}
	return common.GenerateSuccessResponse(context, p.processService.GetByCompanyIdAndRepositoryIdAndAppName(companyId, repositoryId, appId, v1.ProcessQueryOption{}), nil, "")
}

// NewProcessApi returns Process type api
func NewProcessApi(processService service.Process) api.Process {
	return &processApi{
		processService: processService,
	}
}
