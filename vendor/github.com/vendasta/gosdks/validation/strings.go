package validation

import (
	"net/url"
	"strings"
	"unicode/utf8"

	"github.com/vendasta/gosdks/util"
)

type stringRequired struct {
	data      string
	errorType util.ErrorType
	message   string
}

func (r *stringRequired) Validate() error {
	if r.data == "" {
		return util.Error(r.errorType, r.message)
	}
	return nil
}

//StringNotEmpty validates that the provided string is not the empty string
func StringNotEmpty(data string, errorType util.ErrorType, message string) *stringRequired {
	return &stringRequired{data: data, errorType: errorType, message: message}
}

type maxLengthString struct {
	data      string
	errorType util.ErrorType
	message   string
	max       int
}

//StringMaxLength validates that that string does not contain more characters than the provided max
func StringMaxLength(data string, max int, errorType util.ErrorType, message string) *maxLengthString {
	return &maxLengthString{data: data, errorType: errorType, max: max, message: message}
}

func (r *maxLengthString) Validate() error {
	if utf8.RuneCountInString(r.data) > r.max {
		return util.Error(util.InvalidArgument, r.message)
	}
	return nil
}

type atLeastOneStringRequired struct {
	data      []string
	errorType util.ErrorType
	message   string
}

//AtLeastOneStringRequired validates that that at least one string in the slice is not empty
func AtLeastOneStringRequired(data []string, errorType util.ErrorType, message string) *atLeastOneStringRequired {
	return &atLeastOneStringRequired{data: data, errorType: errorType, message: message}
}

func (r *atLeastOneStringRequired) Validate() error {
	oneNotEmpty := false
	for _, str := range r.data {
		if str != "" {
			oneNotEmpty = true
		}
	}
	if !oneNotEmpty {
		return util.Error(r.errorType, r.message)
	}
	return nil
}

type optionalStringValidation struct {
	data string
	rule Rule
}

//OptionalStringValidation will run the validator on the string if the string is not empty
func OptionalStringValidation(data string, rule Rule) *optionalStringValidation {
	return &optionalStringValidation{data: data, rule: rule}
}

func (r *optionalStringValidation) Validate() error {
	if r.data != "" {
		return r.rule.Validate()
	}
	return nil
}

type validUrl struct {
	data      string
	errorType util.ErrorType
	message   string
}

func (r *validUrl) Validate() error {
	_, err := url.ParseRequestURI(r.data)
	if err != nil {
		return util.Error(r.errorType, r.message)
	}
	return nil
}

// ValidURL validates that the provided string is a valid url
func ValidURL(data string, errorType util.ErrorType, message string) *validUrl {
	return &validUrl{data: data, errorType: errorType, message: message}
}

type urlHostEquals struct {
	data      string
	host      string
	errorType util.ErrorType
	message   string
}

func (r *urlHostEquals) Validate() error {
	parsedUrl, err := url.Parse(r.data)

	if err != nil || strings.TrimPrefix(strings.ToLower(parsedUrl.Host), "www.") != strings.ToLower(r.host) {
		return util.Error(r.errorType, r.message)
	}

	return nil
}

// URLHostEquals validates that the url equals the host string.
func URLHostEquals(data string, host string, errorType util.ErrorType, message string) *urlHostEquals {
	return &urlHostEquals{data: data, host: host, errorType: errorType, message: message}
}

type urlHostEndsWith struct {
	data      string
	host      string
	errorType util.ErrorType
	message   string
}

func (r *urlHostEndsWith) Validate() error {
	parsedUrl, err := url.Parse(r.data)

	if err != nil || !strings.HasSuffix(strings.ToLower(parsedUrl.Host), strings.ToLower(r.host)) {
		return util.Error(r.errorType, r.message)
	}

	return nil
}

// URLHostEndsWith validates that the url ends with the host string.
func URLHostEndsWith(data string, host string, errorType util.ErrorType, message string) *urlHostEndsWith {
	return &urlHostEndsWith{data: data, host: host, errorType: errorType, message: message}
}

type urlHostEqualsOneOf struct {
	data      string
	hosts     []string
	errorType util.ErrorType
	message   string
}

func (r *urlHostEqualsOneOf) Validate() error {
	parsedUrl, err := url.Parse(r.data)
	if err != nil {
		return util.Error(r.errorType, r.message)
	}

	for _, host := range r.hosts {
		if strings.TrimPrefix(parsedUrl.Host, "www.") == host {
			return nil
		}
	}
	return util.Error(r.errorType, r.message)
}

// URLHostEqualsOneOf validates that the url contains one of the host strings.
func URLHostEqualsOneOf(data string, hosts []string, errorType util.ErrorType, message string) *urlHostEqualsOneOf {
	return &urlHostEqualsOneOf{data: data, hosts: hosts, errorType: errorType, message: message}
}

type urlPathNotEmpty struct {
	data      string
	errorType util.ErrorType
	message   string
}

func (r *urlPathNotEmpty) Validate() error {
	parsedUrl, err := url.Parse(r.data)

	if err != nil || parsedUrl.Path == "" || (parsedUrl.Path == "/" && parsedUrl.Fragment == "") {
		return util.Error(r.errorType, r.message)
	}

	return nil
}

// URLPathNotEmpty validates that the url contains a nonempty path
func URLPathNotEmpty(data string, errorType util.ErrorType, message string) *urlPathNotEmpty {
	return &urlPathNotEmpty{data: data, errorType: errorType, message: message}
}

type validEmail struct {
	data      string
	errorType util.ErrorType
	message   string
}

// Validate that the email has the proper format of an email
func (r *validEmail) Validate() error {
	indexOfAt := strings.IndexByte(r.data, '@')

	if indexOfAt == 0 || indexOfAt == -1 {
		return util.Error(r.errorType, r.message)
	}
	dataAfterAt := r.data[indexOfAt+1 : len(r.data)]
	if len(dataAfterAt) <= 2 {
		return util.Error(r.errorType, r.message)
	}
	indexOfDot := strings.IndexByte(dataAfterAt, '.')
	if indexOfDot == -1 || indexOfDot == len(dataAfterAt)-1 {
		return util.Error(r.errorType, r.message)
	}
	return nil
}

// ValidEmail validates that the provided string is in a valid email format.
func ValidEmail(data string, errorType util.ErrorType, message string) *validEmail {
	return &validEmail{data: data, errorType: errorType, message: message}
}

type urlHostEndsWithOneOf struct {
	data      string
	hosts     []string
	errorType util.ErrorType
	message   string
}

func (r *urlHostEndsWithOneOf) Validate() error {
	parsedUrl, err := url.Parse(r.data)
	if err != nil {
		return util.Error(r.errorType, r.message)
	}
	for _, host := range r.hosts {
		if strings.HasSuffix(strings.ToLower(parsedUrl.Host), strings.ToLower(host)) {
			return nil
		}
	}
	return util.Error(r.errorType, r.message)
}

// URLHostEndsWithOneOf validates that the url ends with one of the hosts string.
func URLHostEndsWithOneOf(data string, hosts []string, errorType util.ErrorType, message string) *urlHostEndsWithOneOf {
	return &urlHostEndsWithOneOf{data: data, hosts: hosts, errorType: errorType, message: message}
}

type stringInSlice struct {
	data      string
	slice     []string
	errorType util.ErrorType
	message   string
}

func (r *stringInSlice) Validate() error {
	if !util.StringInSlice(r.data, r.slice) {
		return util.Error(r.errorType, r.message)
	}
	return nil
}

// StringInSlice validates that the string is in the provided slice.
func StringInSlice(data string, slice []string, errorType util.ErrorType, message string) *stringInSlice {
	return &stringInSlice{data: data, slice: slice, errorType: errorType, message: message}
}
