package v1

import "time"

// ProcessFootmark ProcessFootmark struct
type ProcessFootmark struct {
	ProcessId string    `bson:"process_id"`
	Step      string    `bson:"step"`
	Footmark  string    `bson:"footmark"`
	Claim     int       `bson:"claim" json:"claim"`
	Time      time.Time `bson:"time" json:"time"`
}

func (ProcessFootmark) GetFootMarks(footmarks []ProcessFootmark) []string {
	footmarkMap := make(map[string]bool)
	var footmarkStrs []string
	for _, each := range footmarks {
		if _, ok := footmarkMap[each.Footmark]; ok {
			continue
		}
		footmarkMap[each.Footmark] = true
		footmarkStrs = append(footmarkStrs, each.Footmark)
	}
	return footmarkStrs
}
