package validation

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/vendasta/gosdks/util"
)

type StringNotEmptyValidateTestSuite struct {
	suite.Suite
}

func TestStringNotEmptyValidateTestSuite(t *testing.T) {
	suite.Run(t, new(StringNotEmptyValidateTestSuite))
}

func (suite *StringNotEmptyValidateTestSuite) Test_ShouldReturnNilWhenDataStringIsNotEmpty() {
	validator := StringNotEmpty("Not empty", util.InvalidArgument, "message")
	suite.Assert().Nil(validator.Validate())
}

func (suite *StringNotEmptyValidateTestSuite) Test_ShouldReturnErrorWhenDataStringIsEmpty() {
	validator := StringNotEmpty("", util.InvalidArgument, "message")
	suite.Assert().NotNil(validator.Validate())
}

func TestStringMaxLength(t *testing.T) {
	type test struct {
		input       string
		max         int
		expected    error
		description string
	}

	cases := []*test{
		{
			input:       "small",
			max:         200,
			expected:    nil,
			description: "'small' is less than 200 characters.",
		},
		{
			input:       "two",
			max:         3,
			expected:    nil,
			description: "'two' is exactly 3 characters.",
		},
		{
			input:       "long",
			max:         2,
			expected:    util.Error(util.InvalidArgument, "message"),
			description: "'long' is more than 2 characters long",
		},
		{
			input:       "世界",
			max:         2,
			expected:    nil,
			description: "'世界' is exactly two characters in length.",
		},
	}

	for _, c := range cases {
		err := StringMaxLength(c.input, c.max, util.InvalidArgument, "message").Validate()
		assert.Equal(t, err, c.expected, c.description)
	}
}

func TestAtLeastOneStringNotEmpty(t *testing.T) {
	type test struct {
		input       []string
		expected    error
		description string
	}

	cases := []*test{
		{
			input:    []string{},
			expected: util.Error(util.InvalidArgument, "message"),
		},
		{
			input:    []string{""},
			expected: util.Error(util.InvalidArgument, "message"),
		},
		{
			input:    []string{"one"},
			expected: nil,
		},
		{
			input:    []string{"", "two"},
			expected: nil,
		},
	}

	for _, c := range cases {
		err := AtLeastOneStringRequired(c.input, util.InvalidArgument, "message").Validate()
		assert.Equal(t, err, c.expected, c.description)
	}
}

func TestValidateURL(t *testing.T) {
	type test struct {
		input       string
		expected    error
		description string
	}

	cases := []*test{
		{
			input: "http://www.google.ca",
		},
		{
			input: "http://www.google.com",
		},
		{
			input: "https://www.google.com",
		},
		{
			input: "localhost:8080",
		},
		{
			input:    "asdf",
			expected: util.Error(util.InvalidArgument, "message"),
		},
		{
			input:    "google.ca",
			expected: util.Error(util.InvalidArgument, "message"),
		},
		{
			input:    "google.com",
			expected: util.Error(util.InvalidArgument, "message"),
		},
	}

	for _, c := range cases {
		err := ValidURL(c.input, util.InvalidArgument, "message").Validate()
		assert.Equal(t, c.expected, err, c.description)
	}
}

func TestUrlHostEquals(t *testing.T) {
	type testCase struct {
		input       string
		host        string
		expected    error
		description string
	}

	cases := []*testCase{
		{
			input:       "https://www.facebook.com",
			host:        "facebook.com",
			expected:    nil,
			description: "Test the simplest case passes.",
		},
		{
			input:       "https://www.facebook.com/McDonalds",
			host:        "facebook.com",
			expected:    nil,
			description: "Test a URL with a path.",
		},
		{
			input:       "https://www.facebook.com/McDonalds?var=value",
			host:        "facebook.com",
			expected:    nil,
			description: "Test a URL with a path and parameters.",
		},
		{
			input:       "https://www.facebook.ca",
			host:        "facebook.com",
			expected:    util.Error(util.InvalidArgument, "message"),
			description: "Test an error is thrown when end of hosts do not match.",
		},
		{
			input:       "http://www.google.com/McDonalds",
			host:        "facebook.com",
			expected:    util.Error(util.InvalidArgument, "message"),
			description: "Test an error is thrown when the URL hosts are not the same.",
		},
		{
			input:       "https://www.google.com/McDonalds?var=value",
			host:        "facebook.com",
			expected:    util.Error(util.InvalidArgument, "message"),
			description: "Test an error is thrown when hosts do not match in longer URLs.",
		},
		{
			input:    "www.google.com/McDonalds?var=value",
			host:     "google.com",
			expected: util.Error(util.InvalidArgument, "message"),
		},
		{
			input:       "https://google.com/McDonalds?var=value",
			host:        "google.com",
			expected:    nil,
			description: "Test that URLs without www. still pass.",
		},
		{
			input:       "https://www.google.com/www.facebook.com?var=value&host=facebook.com",
			host:        "facebook.com",
			expected:    util.Error(util.InvalidArgument, "message"),
			description: "Test an error is thrown when the desired host is further in the URL.",
		},
		{
			input:       "https://sillythings.facebook.com",
			host:        "facebook.com",
			expected:    util.Error(util.InvalidArgument, "message"),
			description: "Test an error is thrown when a silly URL is provided.",
		},
		{
			input:       "https://ww.facebook.com",
			host:        "facebook.com",
			expected:    util.Error(util.InvalidArgument, "message"),
			description: "Test an error is returned when the prefix does not match.",
		},
		{
			input:       "https://googlE.com",
			host:        "google.com",
			expected:    nil,
			description: "Hostnames that only differ in case are equal.",
		},
		{
			input:       "https://Www.google.com",
			host:        "google.com",
			expected:    nil,
			description: "Hostnames with www subdomain that only differ in case are equal.",
		},
	}

	for _, c := range cases {
		err := URLHostEquals(c.input, c.host, util.InvalidArgument, "message").Validate()
		assert.Equal(t, c.expected, err, c.description)
	}
}

func TestUrlHostEndsWith(t *testing.T) {
	type testCase struct {
		input       string
		host        string
		expected    error
		description string
	}

	cases := []*testCase{
		{
			input:       "https://www.facebook.com",
			host:        "facebook.com",
			expected:    nil,
			description: "Test that an exact match passes.",
		},
		{
			input:       "https://www.facebook.ca",
			host:        "facebook.com",
			expected:    util.Error(util.InvalidArgument, "message"),
			description: "Test that an error is thrown when host suffix does not match.",
		},
		{
			input:       "http://ca.linkedin.com/company/124673",
			host:        "linkedin.com",
			expected:    nil,
			description: "Test that a subdomain matches (regression test for AA-3620).",
		},
		{
			input:       "http://CA.LinkedIn.com/company/124673",
			host:        "linkeDin.com",
			expected:    nil,
			description: "Test that a subdomain matches case-insensitively (regression test for AA-3621).",
		},
	}

	for _, c := range cases {
		err := URLHostEndsWith(c.input, c.host, util.InvalidArgument, "message").Validate()
		assert.Equal(t, c.expected, err, c.description)
	}
}

func TestUrlHostEqualsOneOf(t *testing.T) {
	type testCase struct {
		input       string
		hosts       []string
		expected    error
		description string
	}
	cases := []*testCase{
		{
			input:       "https://www.facebook.com",
			hosts:       []string{"facebook.com"},
			expected:    nil,
			description: "Test the simplest case passes.",
		},
		{
			input:       "https://www.fredbook.com/McDonalds",
			hosts:       []string{"facebook.com", "fredbook.com"},
			expected:    nil,
			description: "Test an alternate URL also passes.",
		},
		{
			input:       "https://www.barnybook.com/McDonalds",
			hosts:       []string{"facebook.com", "fredbook.com"},
			expected:    util.Error(util.InvalidArgument, "message"),
			description: "Test it fails if no URLs match.",
		},
	}

	for _, c := range cases {
		err := URLHostEqualsOneOf(c.input, c.hosts, util.InvalidArgument, "message").Validate()
		assert.Equal(t, c.expected, err, c.description)
	}
}

func TestUrlHostEndsWithOneOf(t *testing.T) {
	type testCase struct {
		input       string
		hosts       []string
		expected    error
		description string
	}
	cases := []*testCase{
		{
			input:       "https://www.facebook.com",
			hosts:       []string{"facebook.com"},
			expected:    nil,
			description: "Test the simplest case passes.",
		},
		{
			input:       "https://www.fredbook.com/McDonalds",
			hosts:       []string{"facebook.com", "fredbook.com"},
			expected:    nil,
			description: "Test an alternate URL also passes.",
		},
		{
			input:       "https://www.barnybook.com/McDonalds",
			hosts:       []string{"facebook.com", "fredbook.com"},
			expected:    util.Error(util.InvalidArgument, "message"),
			description: "Test it fails if no URLs match.",
		},
		{
			input:       "https://www.au.fredbook.com/McDonalds",
			hosts:       []string{"fredbook.com"},
			expected:    nil,
			description: "Test a url ends with proper host passes",
		},
		{
			input:       "https://www.au.fredbook.com/McDonalds",
			hosts:       []string{"facebook.com", "fredbook.com"},
			expected:    nil,
			description: "Test an alternate URL also passes on suffix match",
		},
		{
			input:       "https://www.au.fredbook.com/McDonalds",
			hosts:       []string{"facebook.com", "myspace.com", "aol.com", "msn.com", "vendastatus.com", "fredbook.com"},
			expected:    nil,
			description: "Test an alternate URL also passes on suffix match",
		},
		{
			input:       "https://www.au.fredbook.com/McDonalds",
			hosts:       []string{"facebook.com", "fredbook.com", "aol.com", "msn.com", "vendastatus.com"},
			expected:    nil,
			description: "Tests multiple hosts after correct one still passes",
		},
		{
			input:       "https://www.fredbook.co.uk/McDonalds",
			hosts:       []string{"fredbook.co.uk"},
			expected:    nil,
			description: "Test simple single co.uk host domain passes",
		},
		{
			input:       "https://www.fredbook.co.uk/McDonalds",
			hosts:       []string{"facebook.com", "fredbook.co.uk", "aol.com", "msn.com", "vendastatus.com"},
			expected:    nil,
			description: "Test co.uk host domain passes with multiple hosts",
		},
		{
			input:       "https://www.au.fredbook.co.uk/McDonalds",
			hosts:       []string{"facebook.com", "fredbook.co.uk", "aol.com", "msn.com", "vendastatus.com"},
			expected:    nil,
			description: "Test url ends with co.uk host domain passes with multiple hosts",
		},
	}

	for _, c := range cases {
		err := URLHostEndsWithOneOf(c.input, c.hosts, util.InvalidArgument, "message").Validate()
		assert.Equal(t, c.expected, err, c.description)
	}
}

func TestURLPathNotEmpty(t *testing.T) {
	type testCase struct {
		input       string
		host        string
		expected    error
		description string
	}

	cases := []*testCase{
		{
			input:       "https://www.facebook.com/theGoldenItch",
			host:        "facebook.com",
			expected:    nil,
			description: "Test an expected facebook url profile with path",
		},
		{
			input:       "https://www.facebook.com/",
			host:        "facebook.com",
			expected:    util.Error(util.InvalidArgument, "message"),
			description: "Test an expected facebook url profile with a slash for path throws error",
		},
		{
			input:       "https://www.facebook.com",
			host:        "facebook.com",
			expected:    util.Error(util.InvalidArgument, "message"),
			description: "Test an expected facebook url profile with no path throws error",
		},
		{
			input:       "https://www.twitter.com/theGoldenItch",
			expected:    nil,
			description: "Test an expected twitter url profile with path",
		},
		{
			input:       "https://www.twitter.com/#!/theGoldenItch",
			expected:    nil,
			description: "Test an expected twitter url profile with a hashbang in the path",
		},
		{
			input:       "https://www.twitter.com/",
			expected:    util.Error(util.InvalidArgument, "message"),
			description: "Test an expected twitter url profile with a slash for path throws error",
		},
		{
			input:       "https://www.twitter.com",
			expected:    util.Error(util.InvalidArgument, "message"),
			description: "Test an expected twitter url profile with no path throws error",
		},
		{
			input:       "https://www.pinterest.com/theGoldenItch",
			expected:    nil,
			description: "Test an expected pinterest url profile with path",
		},
		{
			input:       "https://www.pinterest.com/",
			expected:    util.Error(util.InvalidArgument, "message"),
			description: "Test an expected pinterest url profile with a slash for path throws error",
		},
		{
			input:       "https://www.pinterest.com",
			expected:    util.Error(util.InvalidArgument, "message"),
			description: "Test an expected pinterest url profile with no path throws error",
		},
		{
			input:       "https://www.instagram.com/theGoldenItch",
			expected:    nil,
			description: "Test an expected instagram url profile with path",
		},
		{
			input:       "https://www.instagram.com/",
			expected:    util.Error(util.InvalidArgument, "message"),
			description: "Test an expected instagram url profile with a slash for path throws error",
		},
		{
			input:       "https://www.instagram.com",
			expected:    util.Error(util.InvalidArgument, "message"),
			description: "Test an expected instagram url profile with no path throws error",
		},
		{
			input:       "https://www.youtube.com/theGoldenItch",
			expected:    nil,
			description: "Test an expected youtube url profile with path",
		},
		{
			input:       "https://www.youtube.com/",
			expected:    util.Error(util.InvalidArgument, "message"),
			description: "Test an expected youtube url profile with a slash for path throws error",
		},
		{
			input:       "https://www.youtube.com",
			expected:    util.Error(util.InvalidArgument, "message"),
			description: "Test an expected youtube url profile with no path throws error",
		},
		{
			input:       "https://www.plus.google.com/+theGoldenItch",
			expected:    nil,
			description: "Test an expected googleplus url profile with path",
		},
		{
			input:       "https://www.plus.google.com/",
			expected:    util.Error(util.InvalidArgument, "message"),
			description: "Test an expected googleplus url profile with a slash for path",
		},
		{
			input:       "https://www.plus.google.com",
			expected:    util.Error(util.InvalidArgument, "message"),
			description: "Test an expected googleplus url profile with no path",
		},
		{
			input:       "https://www.linkedin.com/company/theGoldenItch",
			expected:    nil,
			description: "Test an expected linkedin url profile with path",
		},
		{
			input:       "https://www.linkedin.com/",
			expected:    util.Error(util.InvalidArgument, "message"),
			description: "Test an expected linkedin url profile with a slash for path throws error",
		},
		{
			input:       "https://www.linkedin.com",
			expected:    util.Error(util.InvalidArgument, "message"),
			description: "Test an expected linkedin url profile with no path throws error",
		},
		{
			input:       "https://www.foursquare.com/theGoldenItch",
			expected:    nil,
			description: "Test an expected foursquare url profile with path",
		},
		{
			input:       "https://www.foursquare.com/",
			expected:    util.Error(util.InvalidArgument, "message"),
			description: "Test an expected foursquare url profile with slash for path throws error",
		},
		{
			input:       "https://www.foursquare.com",
			expected:    util.Error(util.InvalidArgument, "message"),
			description: "Test an expected foursquare url profile with no path throws error",
		},
	}

	for _, c := range cases {
		err := URLPathNotEmpty(c.input, util.InvalidArgument, "message").Validate()
		assert.Equal(t, c.expected, err, c.description)
	}
}

func TestValidateEmail(t *testing.T) {
	type test struct {
		input       string
		expected    error
		description string
	}

	cases := []*test{
		{
			input: "cwalker@vendasta.com",
		},
		{
			input: "1@a.b",
		},
		{
			input:    "cwalker",
			expected: util.Error(util.InvalidArgument, "message"),
		},
		{
			input:    "cwalker@.",
			expected: util.Error(util.InvalidArgument, "message"),
		},
		{
			input:    "cwalker@a.",
			expected: util.Error(util.InvalidArgument, "message"),
		},
		{
			input:    "cwalker@.a",
			expected: util.Error(util.InvalidArgument, "message"),
		},
		{
			input:    "@a.b",
			expected: util.Error(util.InvalidArgument, "message"),
		},
		{
			input:    "a.b",
			expected: util.Error(util.InvalidArgument, "message"),
		},
		{
			input:    "a@a.",
			expected: util.Error(util.InvalidArgument, "message"),
		},
		{
			input:    "abc123@234.",
			expected: util.Error(util.InvalidArgument, "message"),
		},
	}

	for _, c := range cases {
		err := ValidEmail(c.input, util.InvalidArgument, "message").Validate()
		assert.Equal(t, c.expected, err, c.description)
	}
}

// 12.8 ns/op	       0 B/op	       0 allocs/op
func BenchmarkValidEmail_Validate(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		ValidEmail("bbass@vendasta.com", util.InvalidArgument, "message").Validate()
	}
}

func TestOptionalStringValidation(t *testing.T) {
	type test struct {
		input       string
		rule        Rule
		expected    error
		description string
	}

	cases := []*test{
		{
			input:       "",
			rule:        StringNotEmpty("", util.InvalidArgument, "This is required"),
			expected:    nil,
			description: "The rule should not run if the input is empty",
		},
		{
			input:       "not-valid",
			rule:        ValidEmail("not-valid", util.InvalidArgument, "That's not a valid email address"),
			expected:    util.Error(util.InvalidArgument, "That's not a valid email address"),
			description: "The rule should not run if the input is empty",
		},
		{
			input:       "valid@email.com",
			rule:        ValidEmail("valid@email.com", util.InvalidArgument, "That's not a valid email address"),
			expected:    nil,
			description: "The rule should pass if the data is valid",
		},
	}

	for _, c := range cases {
		err := OptionalStringValidation(c.input, c.rule).Validate()
		assert.Equal(t, c.expected, err, c.description)
	}
}

func TestStringInSlice(t *testing.T) {
	type testCase struct {
		value       string
		slice       []string
		expected    error
		description string
	}
	cases := []*testCase{
		{
			value:       "one",
			slice:       []string{"one", "two"},
			expected:    nil,
			description: "When the string is in the slice it passes.",
		},
		{
			value:       "one",
			slice:       []string{"two", "three"},
			expected:    util.Error(util.InvalidArgument, "message"),
			description: "When the string is not in the slice it fails",
		},
	}

	for _, c := range cases {
		err := StringInSlice(c.value, c.slice, util.InvalidArgument, "message").Validate()
		assert.Equal(t, c.expected, err, c.description)
	}
}
