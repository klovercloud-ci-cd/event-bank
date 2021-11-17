package service

import (
	v1 "github.com/klovercloud-ci-cd/event-bank/core/v1"
)

// Process Process operations.
type Process interface {
	Store(process v1.Process)
	GetByCompanyIdAndRepositoryIdAndAppName(companyId, repositoryId, appId string, option v1.ProcessQueryOption) []v1.Process
	CountTodaysRanProcessByCompanyId(companyId string) int64
}
