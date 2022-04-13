package v1

// PipelineProcessEvent Pipeline ProcessEvent struct
type PipelineProcessEvent struct {
	ProcessId string                 `bson:"process_id" json:"process_id"`
	Data      map[string]interface{} `bson:"data" json:"data"`
}
