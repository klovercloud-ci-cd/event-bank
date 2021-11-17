package dependency

import (
	"github.com/klovercloud-ci-cd/event-bank/config"
	"github.com/klovercloud-ci-cd/event-bank/core/v1/logic"
	"github.com/klovercloud-ci-cd/event-bank/core/v1/service"
	"github.com/klovercloud-ci-cd/event-bank/enums"
	in_memory "github.com/klovercloud-ci-cd/event-bank/repository/v1/inmemory"
	"github.com/klovercloud-ci-cd/event-bank/repository/v1/mongo"
)

// GetV1LogEventService returns LogEvent service
func GetV1LogEventService() service.LogEvent {

	var logEventService service.LogEvent
	if config.Database == enums.MONGO {
		logEventService = logic.NewLogEventService(mongo.NewLogEventRepository(3000))

	}
	if config.Database == enums.INMEMORY {
		logEventService = logic.NewLogEventService(in_memory.NewLogEventRepository())
	}
	return logic.NewLogEventService(logEventService)
}

// GetV1ProcessEventService returns ProcessEvent service
func GetV1ProcessEventService() service.ProcessEvent {
	var processEventService service.ProcessEvent
	if config.Database == enums.MONGO {
		processEventService = logic.NewProcessEventService(in_memory.NewProcessEventRepository())
	}
	if config.Database == enums.INMEMORY {
		processEventService = logic.NewProcessEventService(in_memory.NewProcessEventRepository())
	}
	return logic.NewProcessEventService(processEventService)
}

// GetV1ProcessService returns Process service
func GetV1ProcessService() service.Process {
	var processService service.Process
	if config.Database == enums.MONGO {
		processService = logic.NewProcessService(mongo.NewProcessRepository(3000))
	}
	if config.Database == enums.INMEMORY {
		processService = logic.NewProcessService(in_memory.NewProcessRepository())
	}
	return logic.NewProcessService(processService)
}

// GetV1ProcessLifeCycleEventService returns ProcessLifeCycleEvent service
func GetV1ProcessLifeCycleEventService() service.ProcessLifeCycleEvent {
	var processLifeCycleEventService service.ProcessLifeCycleEvent
	if config.Database == enums.MONGO {
		processLifeCycleEventService = logic.NewProcessLifeCycleEventService(mongo.NewProcessLifeCycleRepository(3000))
	}
	if config.Database == enums.INMEMORY {
		processLifeCycleEventService = logic.NewProcessLifeCycleEventService(mongo.NewProcessLifeCycleRepository(3000))
	}
	return processLifeCycleEventService
}

// GetV1JwtService returns Jwt service
func GetV1JwtService() service.Jwt {
	return logic.NewJwtService()
}
