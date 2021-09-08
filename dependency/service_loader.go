package dependency

import (
	"github.com/klovercloud-ci/config"
	"github.com/klovercloud-ci/core/v1/logic"
	"github.com/klovercloud-ci/core/v1/service"
	"github.com/klovercloud-ci/enums"
	in_memory "github.com/klovercloud-ci/repository/v1/in-memory"
	"github.com/klovercloud-ci/repository/v1/mongo"
)

func GetPipelineService() service.LogEvent{

	var logEventService service.LogEvent
	if config.Database==enums.Mongo{
		logEventService=logic.NewLogEventService(mongo.NewLogEventRepository(3000))

	}
	if config.Database == enums.Inmemory{
		logEventService=logic.NewLogEventService(in_memory.NewLogEventRepository())
	}
	return logic.NewLogEventService(logEventService)
}