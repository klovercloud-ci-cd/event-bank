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

// ReclaimAbleStatus reclaim process status
type ReclaimAbleStatus struct {
	ClaimAble bool `json:"claim_able" bson:"claim_able"`
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
	Step           string                       `json:"step" bson:"step" yaml:"step"`
	ProcessId      string                       `json:"process_id" bson:"process_id" yaml:"process_id"`
	Descriptors    *[]unstructured.Unstructured `json:"descriptors" yaml:"descriptors" bson:"descriptors"`
	Type           enums.PIPELINE_RESOURCE_TYPE `json:"type" bson:"type" yaml:"type"`
	Name           string                       `json:"name" bson:"name" yaml:"name"`
	Namespace      string                       `json:"namespace" bson:"namespace" yaml:"namespace"`
	Images         []string                     `json:"images" bson:"images" yaml:"images"`
	Pipeline       *Pipeline                    `bson:"pipeline" json:"pipeline" yaml:"pipeline"`
	Claim          int                          `bson:"claim" json:"claim" yaml:"claim"`
	RolloutRestart bool                         `bson:"rollout_restart" json:"rollout_restart" yaml:"rollout_restart"`
}

// CompanyMetadata company metadata
type CompanyMetadata struct {
	Labels                    map[string]string `bson:"labels" json:"labels" yaml:"labels"`
	NumberOfConcurrentProcess int64             `bson:"number_of_concurrent_process" json:"number_of_concurrent_process" yaml:"number_of_concurrent_process"`
	TotalProcessPerDay        int64             `bson:"total_process_per_day" json:"total_process_per_day" yaml:"total_process_per_day"`
}

// PipelineMetadata pipeline metadata
type PipelineMetadata struct {
	CompanyId       string          `json:"company_id" yaml:"company_id" bson:"company_id"`
	CompanyMetadata CompanyMetadata `json:"company_metadata" yaml:"company_metadata" bson:"company_metadata"`
	AllowedBranches string          `json:"allowed_branches" bson:"allowed_branches"`
}

// PipelineStatusCount pipeline status count info
type PipelineStatusCount struct {
	Pipeline struct {
		Completed      int64 `json:"completed"`
		Failed         int64 `json:"failed"`
		Running        int64 `json:"running"`
		Paused         int64 `json:"paused"`
		NonInitialized int64 `json:"nonInitialized"`
	} `json:"pipeline"`
}

// ProcessLifeCycleEventList process life cycle event list
type ProcessLifeCycleEventList struct {
	Events []ProcessLifeCycleEvent `bson:"events" json:"events" yaml:"events"`
}
