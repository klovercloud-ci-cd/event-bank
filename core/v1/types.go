package v1

type LogEventQueryOption struct {
	Pagination struct {
		Page  int64
		Limit int64
	}
	Step string
}

type Subject struct {
	Step,Log string
	EventData map[string]interface{}
	ProcessLabel map[string]string
	ProcessId string
}