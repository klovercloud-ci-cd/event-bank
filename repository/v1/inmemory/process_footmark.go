package inmemory

import (
	v1 "github.com/klovercloud-ci-cd/event-bank/core/v1"
	"github.com/klovercloud-ci-cd/event-bank/core/v1/repository"
)

type processFootmarkRepository struct {
}

func (l processFootmarkRepository) GetFootmarkByProcessIdAndStepAndFootmark(processId, step, footmark string, claim int) *v1.ProcessFootmark {
	//TODO implement me
	panic("implement me")
}

func (l processFootmarkRepository) GetByProcessIdAndStepAndClaim(processId, step string, claim int) []v1.ProcessFootmark {
	return nil
}

func (l processFootmarkRepository) GetByProcessId(processId string) []v1.ProcessFootmark {
	return nil
}

func (l processFootmarkRepository) Store(event v1.ProcessFootmark) {
}

// NewProcessFootmarkRepository returns ProcessLifeCycleEventRepository type object
func NewProcessFootmarkRepository() repository.ProcessFootmarkRepository {
	return &processFootmarkRepository{}

}
