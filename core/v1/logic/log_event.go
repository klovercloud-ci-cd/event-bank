package logic

import (
	v1 "github.com/klovercloud-ci-cd/event-bank/core/v1"
	"github.com/klovercloud-ci-cd/event-bank/core/v1/repository"
	"github.com/klovercloud-ci-cd/event-bank/core/v1/service"
)

type logEventService struct {
	repo                   repository.LogEventRepository
	processFootmarkService service.ProcessFootmark
}

func (l logEventService) GetByProcessIdAndStepAndFootmark(processId string, step string, footmark string, option v1.LogEventQueryOption) ([]string, int64) {
	return l.repo.GetByProcessIdAndStepAndFootmark(processId, step, footmark, option)
}

func (l logEventService) Store(log v1.LogEvent) {
	if log.Log != "" {
		l.repo.Store(log)
	}
	if log.Footmark != "" {
		if l.processFootmarkService.GetFootmarkByProcessIdAndStepAndFootmark(log.ProcessId, log.Step, log.Footmark) == nil {
			l.processFootmarkService.Store(v1.ProcessFootmark{
				ProcessId: log.ProcessId,
				Step:      log.Step,
				Footmark:  log.Footmark,
			})
		}
	}
}

func (l logEventService) GetByProcessId(processId string, option v1.LogEventQueryOption) ([]string, int64) {
	return l.repo.GetByProcessId(processId, option)
}

// NewLogEventService returns LogEvent type service
func NewLogEventService(repo repository.LogEventRepository, processFootmarkService service.ProcessFootmark) service.LogEvent {
	return &logEventService{
		repo:                   repo,
		processFootmarkService: processFootmarkService,
	}
}
