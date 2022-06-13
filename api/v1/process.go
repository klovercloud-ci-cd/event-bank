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
	"log"
	"strconv"
	"strings"
	"time"
)

type processApi struct {
	processService  service.Process
	footmarkService service.ProcessFootmark
	logEventService service.LogEvent
}

// Get... Get Process by Id
// @Summary Get Process by Id
// @Description Get Process by Id
// @Tags Process
// @Produce json
// @Param processId path string true "ProcessId"
// @Param companyId query string true "Company Id"
// @Success 200 {object} common.ResponseDTO{v1.Process}
// @Router /api/v1/processes/{processId} [GET]
func (p processApi) GetById(context echo.Context) error {
	processId := context.Param("processId")
	companyId := context.QueryParam("companyId")
	data := p.processService.GetById(companyId, processId)
	return common.GenerateSuccessResponse(context, data, nil, "Operation Successful!")
}

// Get... Get logs
// @Summary Get Logs
// @Description Gets logs by processId
// @Tags Process
// @Produce json
// @Param processId path string true "Pipeline ProcessId"
// @Param page query int64 false "Page number"
// @Param limit query int64 false "Record count"
// @Success 200 {object} common.ResponseDTO{data=[]string}
// @Router /api/v1/processes/{processId}/logs [GET]
func (p processApi) GetLogsById(context echo.Context) error {
	processId := context.Param("processId")
	option := getQueryOption(context)
	logs, total := p.logEventService.GetByProcessId(processId, option)
	metadata := common.GetPaginationMetadata(option.Pagination.Page, option.Pagination.Limit, total, int64(len(logs)))
	uri := strings.Split(context.Request().RequestURI, "?")[0]
	if option.Pagination.Page > 0 {
		metadata.Links = append(metadata.Links, map[string]string{"prev": uri + "?order=" + context.QueryParam("order") + "&page=" + strconv.FormatInt(option.Pagination.Page-1, 10) + "&limit=" + strconv.FormatInt(option.Pagination.Limit, 10)})
	}
	metadata.Links = append(metadata.Links, map[string]string{"self": uri + "?order=" + context.QueryParam("order") + "&page=" + strconv.FormatInt(option.Pagination.Page, 10) + "&limit=" + strconv.FormatInt(option.Pagination.Limit, 10)})

	if (option.Pagination.Page+1)*option.Pagination.Limit < metadata.TotalCount {
		metadata.Links = append(metadata.Links, map[string]string{"next": uri + "?order=" + context.QueryParam("order") + "&page=" + strconv.FormatInt(option.Pagination.Page+1, 10) + "&limit=" + strconv.FormatInt(option.Pagination.Limit, 10)})
	}

	return common.GenerateSuccessResponse(context, logs, &metadata, "")
}

// Get... Get logs
// @Summary Get Logs
// @Description Gets logs by processId, step, and footmark
// @Tags Process
// @Produce json
// @Param processId path string true "Pipeline ProcessId"
// @Param step path string true "Pipeline step"
// @Param footmark path string true "footmarks"
// @Param claims query string true "claims"
// @Param page query int64 false "Page number"
// @Param limit query int64 false "Record count"
// @Success 200 {object} common.ResponseDTO{data=[]string}
// @Router /api/v1/processes/{processId}/steps/{step}/footmarks/{footmark}/logs [GET]
func (p processApi) GetLogsByProcessIdAndStepAndFootmark(context echo.Context) error {
	processId := context.Param("processId")
	step := context.Param("step")
	footmark := context.Param("footmark")
	claims := context.QueryParam("claims")
	claim, _ := strconv.Atoi(claims)
	option := getQueryOption(context)
	logs, total := p.logEventService.GetByProcessIdAndStepAndFootmark(processId, step, footmark, claim, option)
	metadata := common.GetPaginationMetadata(option.Pagination.Page, option.Pagination.Limit, total, int64(len(logs)))
	uri := strings.Split(context.Request().RequestURI, "?")[0]
	if option.Pagination.Page > 0 {
		metadata.Links = append(metadata.Links, map[string]string{"prev": uri + "?order=" + context.QueryParam("order") + "&page=" + strconv.FormatInt(option.Pagination.Page-1, 10) + "&limit=" + strconv.FormatInt(option.Pagination.Limit, 10)})
	}
	metadata.Links = append(metadata.Links, map[string]string{"self": uri + "?order=" + context.QueryParam("order") + "&page=" + strconv.FormatInt(option.Pagination.Page, 10) + "&limit=" + strconv.FormatInt(option.Pagination.Limit, 10)})

	if (option.Pagination.Page+1)*option.Pagination.Limit < metadata.TotalCount {
		metadata.Links = append(metadata.Links, map[string]string{"next": uri + "?order=" + context.QueryParam("order") + "&page=" + strconv.FormatInt(option.Pagination.Page+1, 10) + "&limit=" + strconv.FormatInt(option.Pagination.Limit, 10)})
	}

	return common.GenerateSuccessResponse(context, logs, &metadata, "")
}

// GetFootmarksByProcessIdAndStep... GetFootmarksByProcessIdAndStep Footmark List
// @Summary Get Footmark List
// @Description Get Footmark List
// @Tags Process
// @Produce json
// @Param processId path string true "Process Id"
// @Param step path string true "step name"
// @Success 200 {object} common.ResponseDTO{data=[]string}
// @Router /api/v1/processes/{processId}/steps/{step} [GET]
func (p processApi) GetFootmarksByProcessIdAndStep(context echo.Context) error {
	process := context.Param("processId")
	step := context.Param("step")
	footmarks := p.footmarkService.GetByProcessIdAndStep(process, step)
	return common.GenerateSuccessResponse(context, v1.ProcessFootmark{}.GetFootMarks(footmarks), nil, "")
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
// @Param appId query string false "App Id"
// @Param appId query string false "Commit Id"
// @Param from query string false "From Date"
// @Param to query string false "To Date"
// @Param operation query string false "Operation[countTodaysProcessByCompanyId/countProcessByCompanyIdAndDate]"
// @Success 200 {object} common.ResponseDTO{data=[]v1.Process}
// @Router /api/v1/processes [GET]
func (p processApi) Get(context echo.Context) error {
	companyId := context.QueryParam("companyId")
	repositoryId := context.QueryParam("repositoryId")
	appId := context.QueryParam("appId")
	operation := context.QueryParam("operation")
	option := getProcessQueryOption(context)
	commitId := context.QueryParam("commitId")
	if operation == "countTodaysProcessByCompanyId" {
		return common.GenerateSuccessResponse(context, p.processService.CountTodaysRanProcessByCompanyId(companyId), nil, "")
	} else if operation == "countProcessByCompanyIdAndDate" {
		from := context.QueryParam("from")
		to := context.QueryParam("to")
		var fromDate time.Time
		var toDate time.Time
		if from != "" {
			date, err := convertDatetoDateTime(from)
			if err != nil {
				return common.GenerateErrorResponse(context, "[ERROR]: Invalid Date Format", err.Error())
			}
			fromDate = date
			if to != "" {
				date, err = convertDatetoDateTime(to)
				if err != nil {
					return common.GenerateErrorResponse(context, "[ERROR]: Invalid Date Format", err.Error())
				}
				toDate = date
			} else {
				toDate = fromDate.AddDate(0, 0, 9).Add(time.Hour * 23).Add(time.Minute * 59).Add(time.Second * 59)
			}
		} else {
			toDate = time.Now().UTC()
			fromDate = toDate.AddDate(0, 0, -9).Add(-time.Hour * 23).Add(-time.Minute * 59).Add(-time.Second * 59)
		}
		return common.GenerateSuccessResponse(context, p.processService.CountProcessByCompanyIdAndDate(companyId, fromDate, toDate), nil, "Operation Successful.")
	} else if commitId == "" {
		return p.GetByCompanyIdAndRepositoryIdAndAppName(context, companyId, repositoryId, appId, option)
	} else {
		return p.GetByCompanyIdAndCommitId(context, companyId, commitId, option)
	}
}

func (p processApi) GetByCompanyIdAndRepositoryIdAndAppName(context echo.Context, companyId, repositoryId, appId string, option v1.ProcessQueryOption) error {
	data, total := p.processService.GetByCompanyIdAndRepositoryIdAndAppName(companyId, repositoryId, appId, option)
	metadata := common.GetPaginationMetadata(option.Pagination.Page, option.Pagination.Limit, total, int64(len(data)))
	uri := strings.Split(context.Request().RequestURI, "?")[0]
	if option.Pagination.Page > 0 {
		metadata.Links = append(metadata.Links, map[string]string{"prev": uri + "?order=" + context.QueryParam("order") + "&page=" + strconv.FormatInt(option.Pagination.Page-1, 10) + "&limit=" + strconv.FormatInt(option.Pagination.Limit, 10)})
	}
	metadata.Links = append(metadata.Links, map[string]string{"self": uri + "?order=" + context.QueryParam("order") + "&page=" + strconv.FormatInt(option.Pagination.Page, 10) + "&limit=" + strconv.FormatInt(option.Pagination.Limit, 10)})

	if (option.Pagination.Page+1)*option.Pagination.Limit < metadata.TotalCount {
		metadata.Links = append(metadata.Links, map[string]string{"next": uri + "?order=" + context.QueryParam("order") + "&page=" + strconv.FormatInt(option.Pagination.Page+1, 10) + "&limit=" + strconv.FormatInt(option.Pagination.Limit, 10)})
	}
	return common.GenerateSuccessResponse(context, data, &metadata, "")
}

func (p processApi) GetByCompanyIdAndCommitId(context echo.Context, companyId, commitId string, option v1.ProcessQueryOption) error {
	data, total := p.processService.GetByCompanyIdAndCommitId(companyId, commitId, option)
	metadata := common.GetPaginationMetadata(option.Pagination.Page, option.Pagination.Limit, total, int64(len(data)))
	uri := strings.Split(context.Request().RequestURI, "?")[0]
	if option.Pagination.Page > 0 {
		metadata.Links = append(metadata.Links, map[string]string{"prev": uri + "?order=" + context.QueryParam("order") + "&page=" + strconv.FormatInt(option.Pagination.Page-1, 10) + "&limit=" + strconv.FormatInt(option.Pagination.Limit, 10)})
	}
	metadata.Links = append(metadata.Links, map[string]string{"self": uri + "?order=" + context.QueryParam("order") + "&page=" + strconv.FormatInt(option.Pagination.Page, 10) + "&limit=" + strconv.FormatInt(option.Pagination.Limit, 10)})

	if (option.Pagination.Page+1)*option.Pagination.Limit < metadata.TotalCount {
		metadata.Links = append(metadata.Links, map[string]string{"next": uri + "?order=" + context.QueryParam("order") + "&page=" + strconv.FormatInt(option.Pagination.Page+1, 10) + "&limit=" + strconv.FormatInt(option.Pagination.Limit, 10)})
	}
	return common.GenerateSuccessResponse(context, data, &metadata, "")
}

func getProcessQueryOption(context echo.Context) v1.ProcessQueryOption {
	option := v1.ProcessQueryOption{}
	page := context.QueryParam("page")
	limit := context.QueryParam("limit")
	option.Step = context.QueryParam("step")
	if page == "" {
		option.Pagination.Page = enums.DEFAULT_PAGE
		option.Pagination.Limit = enums.DEFAULT_PAGE_LIMIT
	} else {
		option.Pagination.Page, _ = strconv.ParseInt(page, 10, 64)
		option.Pagination.Limit, _ = strconv.ParseInt(limit, 10, 64)
	}
	return option
}

// NewProcessApi returns Process type api
func NewProcessApi(processService service.Process, footmarkService service.ProcessFootmark, logEventService service.LogEvent) api.Process {
	return &processApi{
		processService:  processService,
		footmarkService: footmarkService,
		logEventService: logEventService,
	}
}
