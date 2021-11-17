package inmemory

import (
	v1 "github.com/klovercloud-ci-cd/event-bank/core/v1"
	"github.com/klovercloud-ci-cd/event-bank/core/v1/repository"
)

type processRepository struct {
}

func (p processRepository) CountTodaysRanProcessByCompanyId(companyId string) int64 {
	panic("implement me")
}

func (p processRepository) Store(process v1.Process) {
	panic("implement me")
}

func (p processRepository) GetByCompanyIdAndRepositoryIdAndAppName(companyId, repositoryId, appId string, option v1.ProcessQueryOption) []v1.Process {
	panic("implement me")
}

// NewProcessRepository returns ProcessRepository type object
func NewProcessRepository() repository.ProcessRepository {
	return &processRepository{}

}
