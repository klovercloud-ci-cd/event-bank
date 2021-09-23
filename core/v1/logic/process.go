package logic

import (
	"github.com/klovercloud-ci/core/v1/repository"
	"github.com/klovercloud-ci/core/v1/service"
	v1 "github.com/klovercloud-ci/core/v1"
	"time"
)

type processService struct {
	repo repository.ProcessRepository
}

func (p processService) Store(process v1.Process) {
	process.CreatedAt=time.Now().UTC()
	p.repo.Store(process)
}

func (p processService) GetByCompanyIdAndRepositoryIdAndAppName(companyId, repositoryId, appId string, option v1.ProcessQueryOption) []v1.Process {
	return p.repo.GetByCompanyIdAndRepositoryIdAndAppName(companyId,repositoryId,appId,option)
}

func NewProcessService(repo repository.ProcessRepository) service.Process {
	return &processService{
		repo: repo,
	}
}