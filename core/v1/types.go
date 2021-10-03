package v1

import (
	"github.com/klovercloud-ci/enums"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

type LogEventQueryOption struct {
	Pagination struct {
		Page  int64
		Limit int64
	}
	Step string
}
type ProcessQueryOption struct {

}
type PipelineApplyOption struct {
	Purging enums.PIPELINE_PURGING
}
type Subject struct {
	Step,Log string
	EventData map[string]interface{}
	ProcessLabel map[string]string
	ProcessId string
}

type Resource struct {
	Step string `json:"step"`
	ProcessId string `json:"process_id"`
	Descriptors *[]unstructured.Unstructured  `json:"descriptors" yaml:"descriptors"`
	Type     enums.PIPELINE_RESOURCE_TYPE `json:"type"`
	Name string                  `json:"name"`
	Namespace string             `json:"namespace"`
	Images [] string`json:"images"`
}