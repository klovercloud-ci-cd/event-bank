package v1

import (
	"encoding/json"
	"github.com/klovercloud-ci/api/common"
	v1 "github.com/klovercloud-ci/core/v1"
	"github.com/klovercloud-ci/core/v1/api"
	"github.com/klovercloud-ci/core/v1/service"
	"github.com/klovercloud-ci/enums"
	"github.com/labstack/echo/v4"
	"io/ioutil"
	"log"
	"strconv"
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

func (p logEventApi) GetByProcessId(context echo.Context) error {
	processId:=context.Param("processId")
	option := getQueryOption(context)
	logs,total:=p.logEventService.GetByProcessId(processId,option)
	metadata := common.GetPaginationMetadata(option.Pagination.Page, option.Pagination.Limit, total, int64(len(logs)))
	if option.Pagination.Page > 0 {
		metadata.Links = append(metadata.Links, map[string]string{"prev": context.Path() + "?order=" + context.QueryParam("order") + "&page=" + strconv.FormatInt(option.Pagination.Page-1, 10) + "&limit=" + strconv.FormatInt(option.Pagination.Limit, 10)})
	}
	metadata.Links = append(metadata.Links, map[string]string{"self": context.Path() + "?order=" + context.QueryParam("order") + "&page=" + strconv.FormatInt(option.Pagination.Page, 10) + "&limit=" + strconv.FormatInt(option.Pagination.Limit, 10)})

	if (option.Pagination.Page+1)*option.Pagination.Limit < metadata.TotalCount {
		metadata.Links = append(metadata.Links, map[string]string{"next": context.Path() + "?order=" + context.QueryParam("order") + "&page=" + strconv.FormatInt(option.Pagination.Page+1, 10) + "&limit=" + strconv.FormatInt(option.Pagination.Limit, 10)})
	}
	return common.GenerateSuccessResponse(context,logs,&metadata,"")
}

func getQueryOption(context echo.Context) v1.LogEventQueryOption {
	option := v1.LogEventQueryOption{}
	page := context.QueryParam("page")
	limit := context.QueryParam("limit")
	if page == "" {
		option.Pagination.Page = enums.DEFAULT_PAGE
		option.Pagination.Limit = enums.DEFAULT_PAGE_LIMIT
	} else {
		option.Pagination.Page, _ = strconv.ParseInt(page ,10, 64)
		option.Pagination.Limit, _ = strconv.ParseInt(limit ,10, 64)
	}
	return option
}



func NewLogEventApi(logEventService service.LogEvent) api.LogEvent {
	return &logEventApi{
		logEventService: logEventService,
	}
}
