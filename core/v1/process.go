package v1

import "time"

// Process Process struct
type Process struct {
	ProcessId    string                 `bson:"process_id" json:"process_id"`
	CompanyId    string                 `bson:"company_id" json:"company_id"`
	AppId        string                 `bson:"app_id" json:"app_id"`
	RepositoryId string                 `bson:"repository_id" json:"repository_id"`
	Data         map[string]interface{} `bson:"data" json:"data"`
	CreatedAt    time.Time              `bson:"created_at" json:"created_at"`
}
