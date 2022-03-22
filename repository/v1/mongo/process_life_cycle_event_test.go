package mongo

import (
	v1 "github.com/klovercloud-ci-cd/event-bank/core/v1"
	"github.com/stretchr/testify/assert"
	"log"
	"reflect"
	"testing"
)

func TestNewProcessLifeCycleRepository_GetByProcessId(t *testing.T) {
	err := loadEnv(t)
	if err != nil {
		log.Println(err.Error())
	}
	type TestCase struct {
		processId string
		expected []v1.ProcessLifeCycleEvent
		actual   []v1.ProcessLifeCycleEvent
	}
	data := InitProcessLifeCycleEventData()
	var testCases []TestCase
	testCases = append(testCases, TestCase{
		processId: "24033301-5186-4de3-8ab2-469b8d717e45",
		expected: data[:4],
	})
	p := NewMockProcessLifeCycleEventRepository()
	p.Store(data)
	for _, each := range testCases {
		each.actual = p.GetByProcessId(each.processId)
		if !reflect.DeepEqual(each.expected, each.actual) {
			assert.ElementsMatch(t, each.expected, each.actual)
		}
	}
}