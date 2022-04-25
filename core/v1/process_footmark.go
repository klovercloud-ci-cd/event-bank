package v1

// ProcessFootmark ProcessFootmark struct
type ProcessFootmark struct {
	ProcessId string `bson:"process_id"`
	Step      string `bson:"step"`
	Footmark  string `bson:"footmark"`
}

func (ProcessFootmark) GetFootMarks(footmarks []ProcessFootmark) []string {
	footmarkMap := make(map[string]bool)
	for _, each := range footmarks {
		if _, ok := footmarkMap[each.Footmark]; ok {
			continue
		}
		footmarkMap[each.Footmark] = true
	}
	var footmarkStrs []string

	for key, _ := range footmarkMap {
		footmarkStrs = append(footmarkStrs, key)
	}
	return footmarkStrs
}
