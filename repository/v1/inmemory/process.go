package inmemory

import (
	v1 "github.com/klovercloud-ci-cd/event-bank/core/v1"
	"github.com/klovercloud-ci-cd/event-bank/core/v1/repository"
	"time"
)

type processRepository struct {
}

func (p processRepository) CountProcessByCompanyIdAndDate(companyId string, from, to time.Time) int64 {
	//TODO implement me
	panic("implement me")
}

func (p processRepository) GetById(companyId, processId string) v1.Process {
	//TODO implement me
	panic("implement me")
}

func (p processRepository) GetByCompanyIdAndCommitId(companyId, commitId string, option v1.ProcessQueryOption) ([]v1.Process, int64) {
	panic("implement me")
}

func (p processRepository) CountTodaysRanProcessByCompanyId(companyId string) int64 {
	panic("implement me")
}

func (p processRepository) Store(process v1.Process) {
	panic("implement me")
}

func (p processRepository) GetByCompanyIdAndRepositoryIdAndAppName(companyId, repositoryId, appId string, option v1.ProcessQueryOption) ([]v1.Process, int64) {
	panic("implement me")
}

// NewProcessRepository returns ProcessRepository type object
func NewProcessRepository() repository.ProcessRepository {
	return &processRepository{}

}
