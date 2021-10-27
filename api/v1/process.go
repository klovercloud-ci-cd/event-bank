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

type processApi struct {
	processService service.Process
}



func (p processApi) Save(context echo.Context) error {
	var data v1.Process
	body, err := ioutil.ReadAll(context.Request().Body)
	if  err != nil{
		log.Println("Input Error:", err.Error())
		return common.GenerateErrorResponse(context,nil,err.Error())
	}
	if err := json.Unmarshal(body, &data); err != nil {
		return common.GenerateErrorResponse(context,nil,err.Error())
	}
	p.processService.Store(data)
	return common.GenerateSuccessResponse(context,"",nil,"Operation Successful!")
}

func (p processApi) Get(context echo.Context) error {
	companyId := context.QueryParam("companyId")
	repositoryId := context.QueryParam("repositoryId")
	appId:=context.QueryParam("appId")
	operation:=context.QueryParam("operation")
	if operation=="countTodaysProcessByCompanyId"{
		return common.GenerateSuccessResponse(context,p.processService.CountTodaysRanProcessByCompanyId(companyId),nil,"")
	}
	return common.GenerateSuccessResponse(context,p.processService.GetByCompanyIdAndRepositoryIdAndAppName(companyId,repositoryId,appId,v1.ProcessQueryOption{}),nil,"")
}

func NewProcessApi(processService service.Process) api.Process {
	return &processApi{
	processService: processService,
	}
}
