package inmemory

import (
	v1 "github.com/klovercloud-ci-cd/event-bank/core/v1"
	"github.com/klovercloud-ci-cd/event-bank/core/v1/repository"
)

type processFootmarkRepository struct {
}

func (l processFootmarkRepository) GetFootmarkByProcessIdAndStepAndFootmark(processId, step, footmark string) *v1.ProcessFootmark {
	panic("implement me")
}

func (l processFootmarkRepository) GetByProcessIdAndStep(processId, step string) []v1.ProcessFootmark {
	return nil
}

func (l processFootmarkRepository) GetByProcessId(processId string) []v1.ProcessFootmark {
	return nil
}

func (l processFootmarkRepository) Store(event v1.ProcessFootmark) {
}
// NewProcessFootmarkRepository returns ProcessLifeCycleEventRepository type object
func NewProcessFootmarkRepository() repository.ProcessFootmarkRepository {
	return &processFootmarkRepository{
	}

}
