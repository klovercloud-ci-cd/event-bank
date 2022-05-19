package logic

import (
	v1 "github.com/klovercloud-ci-cd/event-bank/core/v1"
	"github.com/klovercloud-ci-cd/event-bank/core/v1/repository"
	"github.com/klovercloud-ci-cd/event-bank/core/v1/service"
)

type processFootmarkService struct {
	repo repository.ProcessFootmarkRepository
}

func (p processFootmarkService) GetFootmarkByProcessIdAndStepAndFootmark(processId, step, footmark string) *v1.ProcessFootmark {
	return p.repo.GetFootmarkByProcessIdAndStepAndFootmark(processId, step, footmark)
}

func (p processFootmarkService) GetByProcessIdAndStep(processId, step string) []v1.ProcessFootmark {
	return p.repo.GetByProcessIdAndStep(processId, step)
}

func (p processFootmarkService) Store(processFootmark v1.ProcessFootmark) {
	p.repo.Store(processFootmark)
}

func (p processFootmarkService) GetByProcessId(processId string) []v1.ProcessFootmark {
	return p.repo.GetByProcessId(processId)
}

// NewProcessFootmarkService returns Process footmark service
func NewProcessFootmarkService(repo repository.ProcessFootmarkRepository) service.ProcessFootmark {
	return &processFootmarkService{
		repo: repo,
	}
}
