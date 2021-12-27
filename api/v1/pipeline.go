package v1

import (
	"encoding/json"
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
)

type pipelineApi struct {
	logEventService     service.LogEvent
	processEventService service.ProcessEvent
}

var (
	upgrader = websocket.Upgrader{}
)

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
	processId := context.QueryParam("processId")
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	ws, err := upgrader.Upgrade(context.Response(), context.Request(), nil)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	defer ws.Close()

	status := make(chan map[string]interface{})
	for {
		go p.processEventService.ReadEventByProcessId(status, processId)
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
func NewPipelineApi(logEventService service.LogEvent, processEventService service.ProcessEvent) api.Pipeline {
	return &pipelineApi{
		logEventService:     logEventService,
		processEventService: processEventService,
	}
}
