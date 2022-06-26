package logic

import (
	v1 "github.com/klovercloud-ci-cd/event-bank/core/v1"
	"github.com/klovercloud-ci-cd/event-bank/repository/v1/inmemory"
	"testing"
)

func TestProcessEventService_Store(t *testing.T) {
	type testCase struct {
		data     []v1.PipelineProcessEvent
		expected map[string]interface{}
		actual   map[string]interface{}
	}

	testCases := testCase{}
	testCases.data = []v1.PipelineProcessEvent{
		{
			ProcessId: "01",
			CompanyId: "1",
			Data:      nil,
				},
		{
			ProcessId: "02",
			CompanyId: "2",
			Data:      nil,
		},
		{
			ProcessId: "03",
			CompanyId: "3",
			Data:      nil,
		},
		{
			ProcessId: "04",
			CompanyId: "4",
			Data:      nil,
		},
	}
	testCases.expected = map[string]interface{}{
		"1": testCases.data[0],
		"2": testCases.data[1],
		"3": testCases.data[2],
		"4": testCases.data[3],
	}
	processEventService := NewProcessEventService(inmemory.NewProcessEventRepository())
	for _, data := range testCases.data {
		processEventService.Store(data)
	}
	testCases.actual = processEventService.GetByCompanyId("1")
	//if len(testCases.actual) != 1 {
	//	t.Errorf("Expected 1, got %d", len(testCases.actual))
	//}
}
