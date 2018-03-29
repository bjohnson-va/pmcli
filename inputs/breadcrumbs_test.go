package inputs_test

import (
	"testing"
	"github.com/bjohnson-va/pmcli/inputs"
	"fmt"
)

type individualizerTest struct {
	caseNo int
	crumb inputs.BreadCrumb
	expected int
}

func TestGeneratesExpectedIndividualizers(t *testing.T) {
	tests := []individualizerTest{
		{
			caseNo: 1,
			crumb: inputs.InitialBreadCrumb(),
			expected: 0,
		},
		{
			caseNo: 2,
			crumb: inputs.InitialBreadCrumb().Indexed(1),
			expected: 0,
		},
		{
			caseNo: 3,
			crumb: inputs.InitialBreadCrumb().AndField("field"),
			expected: 1,
		},
		{
			caseNo: 4,
			crumb: inputs.InitialBreadCrumb().AndField("field").AndField("field"),
			expected: 11,
		},
		{
			caseNo: 5,
			crumb: inputs.InitialBreadCrumb().AndField("field").Indexed(5),
			expected: 5,
		},
		{
			caseNo: 6,
			crumb: inputs.InitialBreadCrumb().AndField("field").Indexed(5).AndField("field"),
			expected: 51,
		},
		{
			caseNo: 7,
			crumb: inputs.InitialBreadCrumb().AndField("field").AndField("field").Indexed(5),
			expected: 15,
		},{
			caseNo: 8,
			crumb: inputs.InitialBreadCrumb().AndField("field").Indexed(2).Indexed(5),
			expected: 5,
		},
	}
	for _, tst := range tests {
		actual := tst.crumb.GetIndividualizer()
		if actual != tst.expected {
			t.Error(fmt.Sprintf("For case: %d\nExpected %d but got %d", tst.caseNo, tst.expected, actual))
		}
	}
}
