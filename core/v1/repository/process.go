package repository

import v1 "github.com/klovercloud-ci-cd/event-bank/core/v1"

// ProcessRepository Process Repository operations.
type ProcessRepository interface {
	Store(process v1.Process)
	GetById(companyId, processId string) v1.Process
	GetByCompanyIdAndRepositoryIdAndAppName(companyId, repositoryId, appId string, option v1.ProcessQueryOption) ([]v1.Process, int64)
	GetByCompanyIdAndCommitId(companyId, commitId string, option v1.ProcessQueryOption) ([]v1.Process, int64)
	CountTodaysRanProcessByCompanyId(companyId string) int64
}
