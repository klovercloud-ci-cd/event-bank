package logic

import (
	v1 "github.com/klovercloud-ci-cd/event-bank/core/v1"
	"github.com/klovercloud-ci-cd/event-bank/core/v1/repository"
	"github.com/klovercloud-ci-cd/event-bank/core/v1/service"
	"time"
)

type processService struct {
	repo repository.ProcessRepository
}

func (p processService) GetById(companyId, processId string) v1.Process {
	return p.repo.GetById(companyId, processId)
}

func (p processService) CountProcessByCompanyIdAndDate(companyId string, from, to time.Time) int64 {
	return p.repo.CountProcessByCompanyIdAndDate(companyId, from, to)
}

func (p processService) CountTodaysRanProcessByCompanyId(companyId string) int64 {
	return p.repo.CountTodaysRanProcessByCompanyId(companyId)
}

func (p processService) Store(process v1.Process) {
	process.CreatedAt = time.Now().UTC()
	p.repo.Store(process)
}

func (p processService) GetByCompanyIdAndRepositoryIdAndAppName(companyId, repositoryId, appId string, option v1.ProcessQueryOption) ([]v1.Process, int64) {
	return p.repo.GetByCompanyIdAndRepositoryIdAndAppName(companyId, repositoryId, appId, option)
}

func (p processService) GetByCompanyIdAndCommitId(companyId, commitId string, option v1.ProcessQueryOption) ([]v1.Process, int64) {
	return p.repo.GetByCompanyIdAndCommitId(companyId, commitId, option)
}

// NewProcessService returns Process type service
func NewProcessService(repo repository.ProcessRepository) service.Process {
	return &processService{
		repo: repo,
	}
}
