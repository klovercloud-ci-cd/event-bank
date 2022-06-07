package v1

// PipelineProcessEvent Pipeline ProcessEvent struct
type PipelineProcessEvent struct {
	Id        string                 `bson:"id" json:"id"`
	ProcessId string                 `bson:"process_id" json:"process_id"`
	CompanyId string                 `bson:"company_id" json:"company_id"`
	Read      bool                   `bson:"read" json:"read"`
	Data      map[string]interface{} `bson:"data" json:"data"`
}
