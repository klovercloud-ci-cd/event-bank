package v1

import (
	"github.com/klovercloud-ci-cd/event-bank/enums"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

// LogEventQueryOption log event query params
type LogEventQueryOption struct {
	Pagination struct {
		Page  int64
		Limit int64
	}
	Step string
}

// ProcessQueryOption process query params
type ProcessQueryOption struct {
	Pagination struct {
		Page  int64
		Limit int64
	}
	Step string
}

// PipelineApplyOption pipeline apply options
type PipelineApplyOption struct {
	Purging enums.PIPELINE_PURGING
}

// Subject subject that observers listen
type Subject struct {
	Step, Log    string
	EventData    map[string]interface{}
	ProcessLabel map[string]string
	ProcessId    string
}

// DeployableResource agent applicable workload info.
type DeployableResource struct {
	Step        string                       `json:"step"`
	ProcessId   string                       `json:"process_id"`
	Descriptors *[]unstructured.Unstructured `json:"descriptors" yaml:"descriptors"`
	Type        enums.PIPELINE_RESOURCE_TYPE `json:"type"`
	Name        string                       `json:"name"`
	Namespace   string                       `json:"namespace"`
	Images      []string                     `json:"images"`
	Pipeline    *Pipeline                    `bson:"pipeline" json:"pipeline"`
	Claim int  `bson:"claim" json:"claim"`
}

// CompanyMetadata company metadata
type CompanyMetadata struct {
	Labels                    map[string]string `bson:"labels" json:"labels" yaml:"labels"`
	NumberOfConcurrentProcess int64             `bson:"number_of_concurrent_process" json:"number_of_concurrent_process" yaml:"number_of_concurrent_process"`
	TotalProcessPerDay        int64             `bson:"total_process_per_day" json:"total_process_per_day" yaml:"total_process_per_day"`
}

// PipelineMetadata pipeline metadata
type PipelineMetadata struct {
	CompanyId       string          `json:"company_id" yaml:"company_id"`
	CompanyMetadata CompanyMetadata `json:"company_metadata" yaml:"company_metadata"`
}

// ProcessLifeCycleEventList process life cycle event list
type ProcessLifeCycleEventList struct {
	Events []ProcessLifeCycleEvent `bson:"events" json:"events" yaml:"events"`
}
