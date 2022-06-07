package mongo

import (
	v1 "github.com/klovercloud-ci-cd/event-bank/core/v1"
	"github.com/klovercloud-ci-cd/event-bank/enums"
	"log"
	"testing"
	"time"
)

func TestNewProcessLifeCycleRepository_GetByProcessId(t *testing.T) {
	err := loadEnv(t)
	if err != nil {
		log.Println(err.Error())
	}
	type TestCase struct {
		data     []v1.ProcessLifeCycleEvent
		expected []v1.ProcessLifeCycleEvent
		actual   []v1.ProcessLifeCycleEvent
	}
	data := InitProcessLifeCycleEventData()
	var testCases TestCase
	testCases = TestCase{
		data:     data,
		expected: data,
	}
	p := NewMockProcessLifeCycleEventRepository()
	p.Store(data)
	for _, each := range testCases.data {
		testCases.actual = append(testCases.actual, p.GetByProcessId(each.ProcessId)...)
	}
	if len(testCases.actual) != len(testCases.expected) {
		t.Errorf("expected: %v, actual: %v", len(testCases.expected), len(testCases.actual))
	}
}

func TestProcessLifeCycleRepository_UpdateStatusesByTime(t *testing.T) {
	err := loadEnv(t)
	if err != nil {
		log.Println(err.Error())
	}
	type TestCase struct {
		time     time.Time
		expected []string
		actual   []string
	}
	data := InitProcessLifeCycleEventData()
	var testCases TestCase
	testCases = TestCase{
		time:     time.Now().UTC().Add(time.Minute * -20),
		expected: []string{"1", "2"},
	}
	p := NewMockProcessLifeCycleEventRepository()
	p.Store(data)
	err = p.UpdateStatusesByTime(testCases.time)
	if err != nil {
		log.Println(err.Error())
	}
	processes := p.Get(8)
	for _, each := range processes {
		if each.Status == enums.ACTIVE {
			testCases.actual = append(testCases.actual, each.ProcessId)
		}
	}
	if len(testCases.actual) != len(testCases.expected) {
		t.Errorf("expected: %v, actual: %v", len(testCases.expected), len(testCases.actual))
	}
	testCases = TestCase{
		time:     time.Now().UTC(),
		expected: []string{},
	}
	err = p.UpdateStatusesByTime(testCases.time)
	processes = p.Get(8)
	for _, each := range processes {
		if each.Status == enums.ACTIVE {
			testCases.actual = append(testCases.actual, each.ProcessId)
		}
	}
	if len(testCases.actual) != len(testCases.expected) {
		t.Errorf("expected: %v, actual: %v", len(testCases.expected), len(testCases.actual))
	}
}
