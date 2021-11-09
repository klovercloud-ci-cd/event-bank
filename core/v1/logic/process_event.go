package logic

import (
	v1 "github.com/klovercloud-ci/core/v1"
	"github.com/klovercloud-ci/core/v1/repository"
	"github.com/klovercloud-ci/core/v1/service"
)

type processEventService struct {
	repo repository.ProcessEventRepository
}

func (p *processEventService) ReadEventByProcessId(c chan map[string]interface{}, processId string) {
	c <- p.DequeueByProcessId(processId)
}

func (p processEventService) Store(data v1.PipelineProcessEvent) {
	p.repo.Store(data)
}

func (p processEventService) GetByProcessId(processId string) map[string]interface{} {
	return p.repo.GetByProcessId(processId)
}

func (p processEventService) DequeueByProcessId(processId string) map[string]interface{} {
	return p.repo.DequeueByProcessId(processId)
}

// NewProcessEventService returns ProcessEvent type service
func NewProcessEventService(repo repository.ProcessEventRepository) service.ProcessEvent {
	return &processEventService{
		repo: repo,
	}
}
