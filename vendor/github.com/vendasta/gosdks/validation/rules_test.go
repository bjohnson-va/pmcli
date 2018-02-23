package validation

import (
	"testing"

	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/vendasta/gosdks/util"
)

type Int64Not0ValidateTestSuite struct {
	suite.Suite
}

func TestInt64Not0ValidateTestSuite(t *testing.T) {
	suite.Run(t, new(Int64Not0ValidateTestSuite))
}

func (suite *Int64Not0ValidateTestSuite) Test_ShouldReturnNilWhenDataIsGreaterThan0() {
	validator := Int64Not0(1, util.InvalidArgument, "message")
	suite.Assert().Nil(validator.Validate())
}

func (suite *Int64Not0ValidateTestSuite) Test_ShouldReturnNilWhenDataIsLessThan0() {
	validator := Int64Not0(-1, util.InvalidArgument, "message")
	suite.Assert().Nil(validator.Validate())
}

func (suite *Int64Not0ValidateTestSuite) Test_ShouldReturnErrorWhenDataIs0() {
	validator := Int64Not0(0, util.InvalidArgument, "message")
	suite.Assert().NotNil(validator.Validate())
}

type Float64Not0ValidateTestSuite struct {
	suite.Suite
}

func TestFloat64Not0ValidateTestSuite(t *testing.T) {
	suite.Run(t, new(Float64Not0ValidateTestSuite))
}

func (suite *Float64Not0ValidateTestSuite) Test_ShouldReturnNilWhenDataIsGreaterThan0() {
	validator := Float64Not0(1, util.InvalidArgument, "message")
	suite.Assert().Nil(validator.Validate())
}

func (suite *Float64Not0ValidateTestSuite) Test_ShouldReturnNilWhenDataIsLessThan0() {
	validator := Float64Not0(-1, util.InvalidArgument, "message")
	suite.Assert().Nil(validator.Validate())
}

func (suite *Float64Not0ValidateTestSuite) Test_ShouldReturnErrorWhenDataIs0() {
	validator := Float64Not0(0, util.InvalidArgument, "message")
	suite.Assert().NotNil(validator.Validate())
}

type TimeNot0ValidateTestSuite struct {
	suite.Suite
}

func TestTimeNot0ValidateTestSuite(t *testing.T) {
	suite.Run(t, new(TimeNot0ValidateTestSuite))
}

func (suite *TimeNot0ValidateTestSuite) Test_ShouldReturnNilWhenDataIsGreaterThan0() {
	validator := TimeNot0(time.Now(), util.InvalidArgument, "message")
	suite.Assert().Nil(validator.Validate())
}

func (suite *TimeNot0ValidateTestSuite) Test_ShouldReturnErrorWhenDataIs0() {
	validator := TimeNot0(time.Time{}, util.InvalidArgument, "message")
	suite.Assert().NotNil(validator.Validate())
}

type TimeAfterValidateTestSuite struct {
	suite.Suite
}

func TestTimeAfterValidateTestSuite(t *testing.T) {
	suite.Run(t, new(TimeAfterValidateTestSuite))
}

func (suite *TimeAfterValidateTestSuite) Test_ShouldReturnNilWhenAfter() {
	validator := TimeAfter(time.Now().Add(time.Hour), time.Now(), true, util.InvalidArgument, "message")
	suite.Assert().Nil(validator.Validate())
}

func (suite *TimeAfterValidateTestSuite) Test_ShouldReturnErrorWhenNotAfter() {
	validator := TimeAfter(time.Now().Add(-time.Hour), time.Now(), true, util.InvalidArgument, "message")
	suite.Assert().NotNil(validator.Validate())
}

func (suite *TimeAfterValidateTestSuite) Test_ShouldReturnNilWhenEqual() {
	now := time.Now()
	validator := TimeAfter(now, now, true, util.InvalidArgument, "message")
	suite.Assert().Nil(validator.Validate())
}

func (suite *TimeAfterValidateTestSuite) Test_ShouldReturnErrorWhenEqualButNotInclusive() {
	validator := TimeAfter(time.Now(), time.Now(), false, util.InvalidArgument, "message")
	suite.Assert().NotNil(validator.Validate())
}

func TestMaxInt(t *testing.T) {
	type test struct {
		input       int
		max         int
		expected    error
		description string
	}

	cases := []*test{
		{
			input:       4,
			max:         5,
			expected:    nil,
			description: "4 is less than 5.",
		},
		{
			input:       5,
			max:         5,
			expected:    nil,
			description: "5 is less than or equal to 5.",
		},
		{
			input:       6,
			max:         5,
			expected:    util.Error(util.InvalidArgument, "message"),
			description: "6 is more than 5.",
		},
	}

	for _, c := range cases {
		err := MaxInt(c.input, c.max, util.InvalidArgument, "message").Validate()
		assert.Equal(t, err, c.expected, c.description)
	}
}

func TestIntBetween(t *testing.T) {
	type test struct {
		input       int
		max         int
		min         int
		expected    error
		description string
	}

	cases := []*test{
		{
			input:       4,
			min:         3,
			max:         5,
			expected:    nil,
			description: "4 is less than 5.",
		},
		{
			input:       5,
			min:         3,
			max:         5,
			expected:    nil,
			description: "5 is less than or equal to 5.",
		},
		{
			input:       3,
			min:         3,
			max:         5,
			expected:    nil,
			description: "3 is greater than or equal to 3.",
		},
		{
			input:       6,
			min:         3,
			max:         5,
			expected:    util.Error(util.InvalidArgument, "message"),
			description: "6 is more than 5.",
		},
		{
			input:       2,
			min:         3,
			max:         5,
			expected:    util.Error(util.InvalidArgument, "message"),
			description: "2 is less than 3.",
		},
	}

	for _, c := range cases {
		err := IntBetween(c.input, c.min, c.max, util.InvalidArgument, "message").Validate()
		assert.Equal(t, c.expected, err, c.description)
	}
}

func TestFloatBetween(t *testing.T) {
	type test struct {
		input       float64
		max         float64
		min         float64
		expected    error
		description string
	}

	cases := []*test{
		{
			input:       4,
			min:         3,
			max:         5,
			expected:    nil,
			description: "4 is less than 5.",
		},
		{
			input:       5,
			min:         3,
			max:         5,
			expected:    nil,
			description: "5 is less than or equal to 5.",
		},
		{
			input:       3,
			min:         3,
			max:         5,
			expected:    nil,
			description: "3 is greater than or equal to 3.",
		},
		{
			input:       5.1,
			min:         3,
			max:         5,
			expected:    util.Error(util.InvalidArgument, "message"),
			description: "5.1 is more than 5.",
		},
		{
			input:       2.9,
			min:         3,
			max:         5,
			expected:    util.Error(util.InvalidArgument, "message"),
			description: "2.9 is less than 3.",
		},
	}

	for _, c := range cases {
		err := FloatBetween(c.input, c.min, c.max, util.InvalidArgument, "message").Validate()
		assert.Equal(t, c.expected, err, c.description)
	}
}

func TestIntGreaterThan(t *testing.T) {
	type test struct {
		data        int64
		bound       int64
		expected    error
		description string
	}

	cases := []*test{
		{
			data:        100,
			bound:       80,
			expected:    nil,
			description: "Data is greater than given bound",
		},
		{
			data:        100,
			bound:       120,
			expected:    util.Error(util.InvalidArgument, "message"),
			description: "Data is less than given bound",
		},
		{
			data:        100,
			bound:       100,
			expected:    util.Error(util.InvalidArgument, "message"),
			description: "Data is equal to given bound",
		},
	}

	for _, c := range cases {
		err := IntGreaterThan(c.data, c.bound, util.InvalidArgument, "message").Validate()
		assert.Equal(t, c.expected, err, c.description)
	}
}

func TestIntLessThan(t *testing.T) {
	type test struct {
		data        int64
		bound       int64
		expected    error
		description string
	}

	cases := []*test{
		{
			data:        80,
			bound:       100,
			expected:    nil,
			description: "Data is less than given bound",
		},
		{
			data:        120,
			bound:       100,
			expected:    util.Error(util.InvalidArgument, "message"),
			description: "Data is greater than given bound",
		},
		{
			data:        100,
			bound:       100,
			expected:    util.Error(util.InvalidArgument, "message"),
			description: "Data is equal to given bound",
		},
	}

	for _, c := range cases {
		err := IntLessThan(c.data, c.bound, util.InvalidArgument, "message").Validate()
		assert.Equal(t, c.expected, err, c.description)
	}
}

func TestValueNotNil(t *testing.T) {
	type test struct {
		input       interface{}
		expected    error
		description string
	}

	cases := []*test{
		{
			input:       &struct{}{},
			expected:    nil,
			description: "Non-nil struct pointer is valid",
		},
		{
			input:       struct{}{},
			expected:    nil,
			description: "Non-nil struct value is valid",
		},
		{
			input:       "",
			expected:    nil,
			description: "Empty string value is valid",
		},

		{
			input:       "Meow",
			expected:    nil,
			description: "Non-empty string value is valid",
		},
		{
			input:       []interface{}{},
			expected:    nil,
			description: "Empty slice is valid",
		},

		{
			input:       false,
			expected:    nil,
			description: "Zero-value boolean is valid",
		},
		{
			input:       nil,
			expected:    util.Error(util.InvalidArgument, "message"),
			description: "Nil value is invalid",
		},
		{
			input:       1,
			expected:    nil,
			description: "Non-zero integer is valid",
		},
		{
			input:       0,
			expected:    nil,
			description: "Zero integer is valid",
		},
	}

	for _, c := range cases {
		err := ValueNotNil(c.input, util.InvalidArgument, "message").Validate()
		assert.Equal(t, c.expected, err, c.description)
	}
}

func TestBoolNotFalse(t *testing.T) {
	type test struct {
		data        bool
		expected    error
		description string
	}

	cases := []*test{
		{
			data:        (1 == 1),
			expected:    nil,
			description: "True",
		},
		{
			data:        (1 == 2),
			expected:    util.Error(util.InvalidArgument, "message"),
			description: "False",
		},
	}

	for _, c := range cases {
		err := BoolNotFalse(c.data, util.InvalidArgument, "message").Validate()
		assert.Equal(t, c.expected, err, c.description)
	}
}

func TestBoolTrue(t *testing.T) {
	type test struct {
		data        bool
		expected    error
		description string
	}

	cases := []*test{
		{
			data:        (1 == 1),
			expected:    nil,
			description: "True",
		},
		{
			data:        (1 == 2),
			expected:    util.Error(util.InvalidArgument, "message"),
			description: "False",
		},
	}

	for _, c := range cases {
		err := BoolTrue(c.data, util.InvalidArgument, "message").Validate()
		assert.Equal(t, c.expected, err, c.description)
	}
}

func TestBoolFalse(t *testing.T) {
	type test struct {
		data        bool
		expected    error
		description string
	}

	cases := []*test{
		{
			data:        false,
			expected:    nil,
			description: "False",
		},
		{
			data:        true,
			expected:    util.Error(util.InvalidArgument, "message"),
			description: "True",
		},
	}

	for _, c := range cases {
		err := BoolFalse(c.data, util.InvalidArgument, "message").Validate()
		assert.Equal(t, c.expected, err, c.description)
	}
}
