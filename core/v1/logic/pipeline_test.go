package logic

import (
	"fmt"
	v1 "github.com/klovercloud-ci-cd/event-bank/core/v1"
	"github.com/klovercloud-ci-cd/event-bank/enums"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestPipelineService_GetProcessStatusMapFromEvents(t *testing.T) {
	type testCase struct {
		data     []v1.ProcessLifeCycleEvent
		expected map[string]enums.PROCESS_STATUS
		actual   map[string]enums.PROCESS_STATUS
	}

	testcase := testCase{}
	testcase.data = []v1.ProcessLifeCycleEvent{
		{
			ProcessId: "01",
			Status:    enums.COMPLETED,
		},
		{
			ProcessId: "01",
			Status:    enums.COMPLETED,
		},

		{
			ProcessId: "02",
			Status:    enums.COMPLETED,
		},
		{
			ProcessId: "02",
			Status:    enums.NON_INITIALIZED,
		},
		{
			ProcessId: "02",
			Status:    enums.COMPLETED,
		},

		{
			ProcessId: "03",
			Status:    enums.COMPLETED,
		},
		{
			ProcessId: "03",
			Status:    enums.FAILED,
		},
		{
			ProcessId: "03",
			Status:    enums.ACTIVE,
		},

		{
			ProcessId: "04",
			Status:    enums.NON_INITIALIZED,
		},
		{
			ProcessId: "04",
			Status:    enums.COMPLETED,
		},

		{
			ProcessId: "05",
			Status:    enums.NON_INITIALIZED,
		},
		{
			ProcessId: "05",
			Status:    enums.ACTIVE,
		},
		{
			ProcessId: "05",
			Status:    enums.NON_INITIALIZED,
		},

		{
			ProcessId: "06",
			Status:    enums.COMPLETED,
		},
		{
			ProcessId: "06",
			Status:    enums.PAUSED,
		},
		{
			ProcessId: "06",
			Status:    enums.NON_INITIALIZED,
		},

		{
			ProcessId: "07",
			Status:    enums.ACTIVE,
		},
		{
			ProcessId: "07",
			Status:    enums.PAUSED,
		},
	}
	testcase.expected = map[string]enums.PROCESS_STATUS{
		"01": enums.COMPLETED,
		"02": enums.PAUSED,
		"03": enums.FAILED,
		"04": enums.NON_INITIALIZED,
		"05": enums.ACTIVE,
		"06": enums.PAUSED,
		"07": enums.ACTIVE,
	}
	testcase.actual = GetProcessStatusMapFromEvents(testcase.data)
	if !reflect.DeepEqual(testcase.expected, testcase.actual) {
		fmt.Println(testcase.actual)
		assert.ElementsMatch(t, testcase.expected, testcase.actual)
	}
}
