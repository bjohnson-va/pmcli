package inputs

import (
	"testing"
	"fmt"
)

type buildIndexedKeyTest struct {
	description string
	crumb BreadCrumb
	expected interface{}
}

func TestBuildIndexedKeyReturnsExpectedResults(t *testing.T) {
	tests := []buildIndexedKeyTest{
		{
			description: "initial",
			crumb: InitialBreadCrumb(),
			expected: "",
		},
		{
			description: "nested",
			crumb: InitialBreadCrumb().AndField("field"),
			expected: "field",
		},
		{
			description: "double-nested",
			crumb: InitialBreadCrumb().AndField("field").AndField("inner"),
			expected: "field.inner",
		},
		{
			description: "repeated",
			crumb: InitialBreadCrumb().AndField("field").Indexed(5),
			expected: "field#5",
		},
		{
			description: "nested within repeated",
			crumb: InitialBreadCrumb().AndField("field").Indexed(5).AndField("inner"),
			expected: "field#5.inner",
		},
		{
			description: "repeated within repeated",
			crumb: InitialBreadCrumb().AndField("field").Indexed(5).AndField("inner").Indexed(9),
			expected: "field#5.inner#9",
		},
		{
			description: "0th",
			crumb: InitialBreadCrumb().AndField("field").Indexed(0),
			expected: "field#0",
		},
	}
	for _, tst := range tests {
		actual := buildIndexedKey(tst.crumb)
		if actual != tst.expected {
			t.Error(fmt.Sprintf("For case: %v\nExpected %v but got %v", tst.description, tst.expected, actual))
		}
	}
}

