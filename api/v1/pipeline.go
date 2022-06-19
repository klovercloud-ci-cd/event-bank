package v1

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/websocket"
	"github.com/klovercloud-ci-cd/event-bank/api/common"
	v1 "github.com/klovercloud-ci-cd/event-bank/core/v1"
	"github.com/klovercloud-ci-cd/event-bank/core/v1/api"
	"github.com/klovercloud-ci-cd/event-bank/core/v1/service"
	"github.com/klovercloud-ci-cd/event-bank/enums"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type pipelineApi struct {
	pipelineService     service.Pipeline
	logEventService     service.LogEvent
	processEventService service.ProcessEvent
}

var (
	upgrader = websocket.Upgrader{}
)

// Get... Get Pipeline
// @Summary Get Pipeline
// @Description Gets status count of the pipeline by company id and a given date
//@Tags Pipeline
// @Produce json
// @Param action query string true "action [dashboard_data]"
// @Param companyId query string true "Company Id"
// @Param from query string true "From Data"
// @Param to query string true "To Data"
// @Success 200 {object} common.ResponseDTO{data=v1.PipelineStatusCount}
// @Router /api/v1/pipelines [GET]
func (p pipelineApi) Get(context echo.Context) error {
	action := context.QueryParam("action")
	if action == "dashboard_data" {
		companyId := context.QueryParam("companyId")
		if companyId == "" {
			return common.GenerateErrorResponse(context, "[ERROR]: Company Id is not provided.", "Operation Failed.")
		}
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
		data := p.pipelineService.GetStatusCount(companyId, fromDate, toDate)
		return common.GenerateSuccessResponse(context, data, nil, "Operation Successful.")
	}
	return common.GenerateErrorResponse(context, "[ERROR]: Invalid action type is given.", "Operation Failed")
}

// Get... Get Pipeline
// @Summary Get Pipeline
// @Description Gets pipeline or logs by pipeline processId If action is "get_pipeline", then pipeline will be returned or logs will be returned.
// @Tags Pipeline
// @Produce json
// @Param processId path string true "Pipeline ProcessId"
// @Param action query string false "action"
// @Param page query int64 false "Page number"
// @Param limit query int64 false "Record count"
// @Success 200 {object} common.ResponseDTO{data=[]string}
// @Router /api/v1/pipelines/{processId} [GET]
func (p pipelineApi) GetById(context echo.Context) error {
	action := context.QueryParam("action")
	if action == "get_pipeline" {
		return p.GetByProcessId(context)
	}
	return p.GetLogs(context)
}

func convertDatetoDateTime(date string) (time.Time, error) {
	layout := "2006-1-2"
	d, err := time.Parse(layout, date)
	if err != nil {
		return time.Time{}, errors.New("invalid date format")
	}
	dateTime, err := time.Parse(time.RFC3339, d.Format(time.RFC3339))
	if err != nil {
		return time.Time{}, errors.New("invalid date format")
	}
	return dateTime, nil
}

// Get... Get by process
// @Summary Get by process id
// @Description Gets pipeline by process id
// @Tags Pipeline
// @Produce json
// @Param commitId path string true "processId"
// @Success 200 {object} common.ResponseDTO{data=[]string}
// @Router /api/v1/pipelines/{commitId} [GET]
func (p pipelineApi) GetByProcessId(context echo.Context) error {
	processId := context.Param("processId")
	companyId := context.QueryParam("companyId")
	data := p.pipelineService.GetByProcessId(processId)
	if data.MetaData.CompanyId != companyId {
		return common.GenerateSuccessResponse(context, v1.Pipeline{}, nil, "")
	}
	return common.GenerateSuccessResponse(context, data, nil, "")
}

// Get... Get logs
// @Summary Get Logs
// @Description Gets logs by pipeline processId
// @Tags Pipeline
// @Produce json
// @Param processId path string true "Pipeline ProcessId"
// @Param page query int64 false "Page number"
// @Param limit query int64 false "Record count"
// @Success 200 {object} common.ResponseDTO{data=[]string}
// @Router /api/v1/pipelines/{processId} [GET]
func (p pipelineApi) GetLogs(context echo.Context) error {
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

func getQueryOption(context echo.Context) v1.LogEventQueryOption {
	option := v1.LogEventQueryOption{}
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

func (p pipelineApi) GetEvents(context echo.Context) error {
	companyId := context.QueryParam("companyId")
	userId := context.QueryParam("userId")
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	ws, err := upgrader.Upgrade(context.Response(), context.Request(), nil)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	defer ws.Close()

	status := make(chan map[string]interface{})
	for {
		go p.processEventService.ReadEventByCompanyIdAndUserId(status, companyId, userId)
		jsonStr, err := json.Marshal(<-status)
		if err != nil {
			log.Println("[ERROR]: Failed to marshal", err.Error())
		}
		err = ws.WriteMessage(websocket.TextMessage, []byte(jsonStr))
		if err != nil {
			log.Println("[ERROR]: Failed to write", err.Error())
			ws.Close()
		}
		_, _, err = ws.ReadMessage()
		if err != nil {
			log.Println("[ERROR]: Failed to read", err.Error())
			ws.Close()
		}

	}
}

// NewPipelineApi returns Pipeline type api
func NewPipelineApi(pipelineService service.Pipeline, logEventService service.LogEvent, processEventService service.ProcessEvent) api.Pipeline {
	return &pipelineApi{
		pipelineService:     pipelineService,
		logEventService:     logEventService,
		processEventService: processEventService,
	}
}
