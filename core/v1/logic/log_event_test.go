package logic

import (
	v1 "github.com/klovercloud-ci/core/v1"
	inmemory "github.com/klovercloud-ci/repository/v1/in-memory"
	"github.com/stretchr/testify/assert"
	"reflect"
	"strconv"
	"testing"
	"time"
)

func Test_GetByProcessId(t *testing.T) {
	type TestData struct {
		data     v1.LogEvent
		expected int
		actual   int
		option   v1.LogEventQueryOption
	}
	var testCases []TestData

	for i := 0; i < 50; i++ {
		testCases = append(testCases, TestData{
			data: v1.LogEvent{
				ProcessId: "01",
				Log:       "log-" + strconv.Itoa(i),
				Step:      "BUILD",
				CreatedAt: time.Time{}.UTC(),
			},
			expected: i + 1,
			option: v1.LogEventQueryOption{
				Pagination: struct {
					Page  int64
					Limit int64
				}{
					Page:  0,
					Limit: int64(i + 1),
				},
			},
		})
	}
	logEventService := NewLogEventService(inmemory.NewLogEventRepository())
	for _, each := range testCases {
		logEventService.Store(each.data)
		logs, _ := logEventService.GetByProcessId(each.data.ProcessId, each.option)
		if !reflect.DeepEqual(len(logs), each.expected) {
			assert.ElementsMatch(t, len(logs), each.expected)
		}
	}
	testCases = append(testCases, TestData{
		data: v1.LogEvent{
			ProcessId: "01",
			Log:       "log-99",
			Step:      "DEVELOP",
			CreatedAt: time.Time{}.UTC(),
		},
		expected: 1,
		option: v1.LogEventQueryOption{
			Pagination: struct {
				Page  int64
				Limit int64
			}{
				Page:  0,
				Limit: 10,
			},
			Step: "DEVELOP",
		},
	})
	logEventService = NewLogEventService(inmemory.NewLogEventRepository())
	logEventService.Store(testCases[50].data)
	logs, _ := logEventService.GetByProcessId(testCases[50].data.ProcessId, testCases[50].option)
	if !reflect.DeepEqual(len(logs), testCases[50].expected) {
		assert.ElementsMatch(t, len(logs), testCases[50].expected)
	}
}

func Test_Store(t *testing.T) {
	type TestData struct {
		data     v1.LogEvent
		expected int64
		actual   int
		option   v1.LogEventQueryOption
	}
	var testCases []TestData

	for i := 0; i < 50; i++ {
		testCases = append(testCases, TestData{
			data: v1.LogEvent{
				ProcessId: "01",
				Log:       "log-" + strconv.Itoa(i),
				Step:      "BUILD",
				CreatedAt: time.Time{}.UTC(),
			},
			expected: int64(i + 52),
			option: v1.LogEventQueryOption{
				Pagination: struct {
					Page  int64
					Limit int64
				}{
					Page:  0,
					Limit: 100,
				},
			},
		})
	}
	logEventService := NewLogEventService(inmemory.NewLogEventRepository())
	for _, each := range testCases {
		logEventService.Store(each.data)
		_, size := logEventService.GetByProcessId(each.data.ProcessId, each.option)
		if size != each.expected {
			assert.ElementsMatch(t, size, each.expected)
		}
	}
}
