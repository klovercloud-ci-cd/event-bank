package service

import (
	v1 "github.com/klovercloud-ci-cd/event-bank/core/v1"
)

// LogEvent Log Event operations.
type LogEvent interface {
	Store(log v1.LogEvent)
	GetByProcessId(processId string, option v1.LogEventQueryOption) ([]string, int64)
	GetByProcessIdAndStepAndFootmark(processId string, step string, footmark string, claim int, option v1.LogEventQueryOption) ([]string, int64)
	GetByProcessIdAndStepAndClaim(processId string, step string, claim int, option v1.LogEventQueryOption) ([]v1.LogEvent, int64)
}
