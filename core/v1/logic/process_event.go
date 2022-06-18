package logic

import (
	v1 "github.com/klovercloud-ci-cd/event-bank/core/v1"
	"github.com/klovercloud-ci-cd/event-bank/core/v1/repository"
	"github.com/klovercloud-ci-cd/event-bank/core/v1/service"
)

type processEventService struct {
	repo repository.ProcessEventRepository
}

func (p *processEventService) GetByCompanyIdAndProcessId(companyId, processId string, option v1.ProcessQueryOption) ([]v1.PipelineProcessEvent, int64) {
	return p.repo.GetByCompanyIdAndProcessId(companyId, processId, option)
}

func (p *processEventService) ReadEventByCompanyIdAndUserId(c chan map[string]interface{}, companyId, userId string) {
	c <- p.DequeueByCompanyIdAndUserId(companyId, userId)
}

func (p processEventService) Store(data v1.PipelineProcessEvent) {
	p.repo.Store(data)
}

func (p processEventService) GetByCompanyId(companyId string) map[string]interface{} {
	return p.repo.GetByCompanyId(companyId)
}

func (p processEventService) DequeueByCompanyIdAndUserId(companyId, userId string) map[string]interface{} {
	return p.repo.DequeueByCompanyIdAndUserId(companyId, userId)
}

// NewProcessEventService returns ProcessEvent type service
func NewProcessEventService(repo repository.ProcessEventRepository) service.ProcessEvent {
	return &processEventService{
		repo: repo,
	}
}
