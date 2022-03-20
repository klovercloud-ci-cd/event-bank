package service

import (
	v1 "github.com/klovercloud-ci-cd/event-bank/core/v1"
)

// Pipeline Pipeline operations
type Pipeline interface {
	GetByProcessId(processId string) v1.Pipeline
}