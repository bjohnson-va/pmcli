package inputs_test

import (
	"testing"
	"github.com/bjohnson-va/pmcli/inputs"
	"fmt"
	"github.com/bjohnson-va/pmcli/config"
	"github.com/emicklei/proto"
)

type getOverrideTest struct {
	description string
	crumb inputs.BreadCrumb
	expected interface{}
	mockserverFileName string
}

func TestGetsOverridesAsExpected(t *testing.T) {
	tests := []getOverrideTest{
		{
			description: "override outermost field",
			crumb: inputs.InitialBreadCrumb().AndField("nonRepeatedField"),
			expected: float64(123),
			mockserverFileName: "./test_data/top_level_override.json",
		},
		{
			description: "override nested field",
			crumb: inputs.InitialBreadCrumb().AndField("nonRepeatedField").AndField("innerField"),
			expected: float64(456),
			mockserverFileName: "./test_data/inner_override_on_nonrepeated.json",
		},
		{
			description: "override nested field on specific index of repeated field",
			crumb: inputs.InitialBreadCrumb().AndField("repeatedField").Indexed(5).AndField("innerField"),
			expected: float64(789),
			mockserverFileName: "./test_data/inner_override_on_repeated.json",
		},
		{
			description: "override nested field on 0th index of repeated field",
			crumb: inputs.InitialBreadCrumb().AndField("repeatedField").Indexed(0).AndField("innerField"),
			expected: float64(191),
			mockserverFileName: "./test_data/inner_override_on_repeated_0th.json",
		},
	}
	for _, tst := range tests {
		i := loadInputsFromFile(tst.mockserverFileName)
		actual := i.GetFieldOverride(tst.crumb, nil)
		if actual != tst.expected {
			t.Error(fmt.Sprintf("For case: %v\nExpected %v but got %v", tst.description, tst.expected, actual))
		}
	}
}

func loadInputsFromFile(filename string) inputs.Provider {
	file, err := config.ReadFile(filename)
	if err != nil {
		panic(fmt.Sprintf("Error during test setup: %s", err.Error()))
	}
	s := proto.Service{
		Name: "Service",
	}
	r := proto.RPC{
		Name: "RPC",
	}
	i, err := config.GetInputsForRPC(s, r, file.ConfigMap)
	if err != nil {
		panic(fmt.Sprintf("Error during test setup: %s", err.Error()))
	}
	return i
}


