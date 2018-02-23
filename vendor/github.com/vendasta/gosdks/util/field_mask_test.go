package util

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"fmt"
)

func Test_FieldMaskProtoToFieldMask_ReturnsNilOnEmptyPath(t *testing.T) {
	assert.Nil(t, FieldMaskProtoPathsToFieldMask([]string{}))
}

func Test_FieldMaskProtoToFieldMask_ReturnsNilOnPathWithEmptyString(t *testing.T) {
	assert.Nil(t, FieldMaskProtoPathsToFieldMask([]string{""}))
}

func Test_FieldMaskProtoToFieldMask_SupportsMultipleTextCases(t *testing.T) {
	fm := FieldMaskProtoPathsToFieldMask([]string{"testCase", "TestCase2", "test_case_three"})

	cases := []string{
		"testcase",
		"test_case",
		"testcase2",
		"test_case2",
		"test_case_three",
		"testcasethree",
	}

	for _, c := range cases {
		_, ok := (*fm)[c]
		assert.True(t, ok, fmt.Sprintf(`fm["%s"] should be true`, c))
	}
}

func Test_IsPropertyPresent_SupportsManyDifferentTextCases(t *testing.T) {
	fm := FieldMaskProtoPathsToFieldMask([]string{"testCase", "TestCase2", "test_case_three"})

	cases := []string{
		"testCase",
		"test_case",
		"TestCase",
		"testCase2",
		"TestCase2",
		"test_case2",
		"test_case_three",
		"TestCaseThree",
		"testCaseThree",
	}

	for _, c := range cases {
		assert.True(t, fm.IsPropertyPresent(c), fmt.Sprintf(`fm.IsPropertyPresent("%s") should be true`, c))
	}
}

func Test_AllPropertiesPresent_IsTrueIfAllPresent(t *testing.T) {
	type test struct {
		fieldMask   *FieldMask
		properties  []string
		expected    bool
		description string
	}

	fm := FieldMaskProtoPathsToFieldMask([]string{"propertyA", "propertyB"})


	cases := []*test{
		{
			fieldMask: fm,
			properties: []string{"propertyA"},
			expected: true,
			description: "should be true, subset",
		},
		{
			fieldMask: fm,
			properties: []string{"propertyA", "propertyB"},
			expected: true,
			description: "should be true, same set",
		},
		{
			fieldMask: nil,
			properties: []string{"propertyA", "propertyB"},
			expected: true,
			description: "should be false, no field mask",
		},
		{
			fieldMask: fm,
			properties: []string{"propertyA", "propertyB", "propertyC"},
			expected: false,
			description: "should be false, superset",
		},
		{
			fieldMask: fm,
			properties: []string{"propertyX"},
			expected: false,
			description: "should be false, no shared between set",
		},
	}

	for _, c := range cases {
		assert.Equal(t, c.expected, c.fieldMask.AllPropertiesPresent(c.properties...), c.description)
	}
}
