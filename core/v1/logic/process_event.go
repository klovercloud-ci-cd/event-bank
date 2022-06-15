package logic

import (
	v1 "github.com/klovercloud-ci-cd/event-bank/core/v1"
	"github.com/klovercloud-ci-cd/event-bank/core/v1/repository"
	"github.com/klovercloud-ci-cd/event-bank/core/v1/service"
)

type processEventService struct {
	repo repository.ProcessEventRepository
}

func (p *processEventService) GetByCompanyIdAndProcessId(companyId, processId string, option v1.ProcessQueryOption) []v1.PipelineProcessEvent {
	return p.repo.GetByCompanyIdAndProcessId(companyId, processId, option)
}

func (p *processEventService) ReadEventByCompanyId(c chan map[string]interface{}, processId string) {
	c <- p.DequeueByCompanyId(processId)
}

func (p processEventService) Store(data v1.PipelineProcessEvent) {
	p.repo.Store(data)
}

func (p processEventService) GetByCompanyId(companyId string) map[string]interface{} {
	return p.repo.GetByCompanyId(companyId)
}

func (p processEventService) DequeueByCompanyId(companyId string) map[string]interface{} {
	return p.repo.DequeueByCompanyId(companyId)
}

// NewProcessEventService returns ProcessEvent type service
func NewProcessEventService(repo repository.ProcessEventRepository) service.ProcessEvent {
	return &processEventService{
		repo: repo,
	}
}
