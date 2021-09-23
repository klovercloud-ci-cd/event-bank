package repository
import v1 "github.com/klovercloud-ci/core/v1"

type ProcessRepository interface {
	Store( process v1.Process)
	GetByCompanyIdAndRepositoryIdAndAppName(companyId,repositoryId,appId string,option v1.ProcessQueryOption)[]v1.Process
}
