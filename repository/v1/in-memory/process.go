package in_memory
import (
	v1 "github.com/klovercloud-ci/core/v1"
	"github.com/klovercloud-ci/core/v1/repository"
)
type processRepository struct {

}

func (p processRepository) CountTodaysRanProcessByCompanyId(companyId string) int64 {
	panic("implement me")
}

func (p processRepository) Store(process v1.Process) {
	panic("implement me")
}

func (p processRepository) GetByCompanyIdAndRepositoryIdAndAppName(companyId, repositoryId, appId string,option v1.ProcessQueryOption) []v1.Process {
	panic("implement me")
}

func NewProcessRepository() repository.ProcessRepository {
	return &processRepository{
	}

}