package v1

// ProcessFootmark ProcessFootmark struct
type ProcessFootmark struct{
	ProcessId string    `bson:"process_id"`
	Step      string    `bson:"step"`
	Footmark      string    `bson:"footmark"`
}