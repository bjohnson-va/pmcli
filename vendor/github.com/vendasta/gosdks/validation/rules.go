package validation

import (
	"time"

	"github.com/vendasta/gosdks/util"
)

type int64Required struct {
	data      int64
	errorType util.ErrorType
	message   string
}

func (r *int64Required) Validate() error {
	if r.data == 0 {
		return util.Error(r.errorType, r.message)
	}
	return nil
}

//Int64Not0 validates that the provided int64 is not 0
func Int64Not0(data int64, errorType util.ErrorType, message string) *int64Required {
	return &int64Required{data: data, errorType: errorType, message: message}
}

type float64Required struct {
	data      float64
	errorType util.ErrorType
	message   string
}

func (r *float64Required) Validate() error {
	if r.data == 0 {
		return util.Error(r.errorType, r.message)
	}
	return nil
}

//Float64Not0 validates that the provided float64 is not 0
func Float64Not0(data float64, errorType util.ErrorType, message string) *float64Required {
	return &float64Required{data: data, errorType: errorType, message: message}
}

type timeRequired struct {
	data      time.Time
	errorType util.ErrorType
	message   string
}

func (r *timeRequired) Validate() error {
	if r.data.IsZero() {
		return util.Error(r.errorType, r.message)
	}
	return nil
}

//TimeNot0 validates that the provided time.Time is not the zero value
func TimeNot0(data time.Time, errorType util.ErrorType, message string) *timeRequired {
	return &timeRequired{data: data, errorType: errorType, message: message}
}

type timeAfter struct {
	laterTime   time.Time
	earlierTime time.Time
	errorType   util.ErrorType
	message     string
	inclusive   bool
}

func (r *timeAfter) Validate() error {
	if r.inclusive {
		if r.laterTime.Equal(r.earlierTime) {
			return nil
		}
	}
	if !r.laterTime.After(r.earlierTime) {
		return util.Error(r.errorType, r.message)
	}
	return nil
}

// TimeAfter validates that the first time is after the second time. If inclusive is True, >= comparison operator is
// used, otherwise >
func TimeAfter(laterTime time.Time, earlierTime time.Time, inclusive bool, errorType util.ErrorType, message string) *timeAfter {
	return &timeAfter{
		laterTime:   laterTime,
		earlierTime: earlierTime,
		errorType:   errorType,
		message:     message,
		inclusive:   inclusive,
	}
}

type maxInt struct {
	data      int
	errorType util.ErrorType
	message   string
	max       int
}

//MaxInt validates that the int is less than or equal to the provided max
func MaxInt(data int, max int, errorType util.ErrorType, message string) *maxInt {
	return &maxInt{data: data, errorType: errorType, max: max, message: message}
}

func (r *maxInt) Validate() error {
	if r.data > r.max {
		return util.Error(util.InvalidArgument, r.message)
	}
	return nil
}

type intGreaterThan struct {
	data      int64
	bound     int64
	errorType util.ErrorType
	message   string
}

func (i intGreaterThan) Validate() error {
	if i.data <= i.bound {
		return util.Error(i.errorType, i.message)
	}
	return nil
}

// IntGreaterThan validates that the int64 data is greater than the bound
func IntGreaterThan(data int64, bound int64, errorType util.ErrorType, message string) *intGreaterThan {
	return &intGreaterThan{data: data, bound: bound, errorType: errorType, message: message}
}

type intLessThan struct {
	data      int64
	bound     int64
	errorType util.ErrorType
	message   string
}

func (i intLessThan) Validate() error {
	if i.data >= i.bound {
		return util.Error(i.errorType, i.message)
	}
	return nil
}

// IntLessThan validates that the int64 data is less than the bound
func IntLessThan(data int64, bound int64, errorType util.ErrorType, message string) *intLessThan {
	return &intLessThan{data: data, bound: bound, errorType: errorType, message: message}
}

type intBetween struct {
	data      int
	errorType util.ErrorType
	message   string
	max       int
	min       int
}

//IntBetween validates that the int is between min and max
func IntBetween(data int, min int, max int, errorType util.ErrorType, message string) *intBetween {
	return &intBetween{data: data, errorType: errorType, max: max, min: min, message: message}
}

func (r *intBetween) Validate() error {
	if r.data > r.max || r.data < r.min {
		return util.Error(util.InvalidArgument, r.message)
	}
	return nil
}

type floatBetween struct {
	data      float64
	errorType util.ErrorType
	message   string
	max       float64
	min       float64
}

//FloatBetween validates that the float is between min and max
func FloatBetween(data float64, min float64, max float64, errorType util.ErrorType, message string) *floatBetween {
	return &floatBetween{data: data, errorType: errorType, max: max, min: min, message: message}
}

func (r *floatBetween) Validate() error {
	if r.data > r.max || r.data < r.min {
		return util.Error(util.InvalidArgument, r.message)
	}
	return nil
}

type valueNotNil struct {
	data      interface{}
	errorType util.ErrorType
	message   string
}

//ValueNotNil validates that the struct pointer value provided is not nil
//ex: ValueNotNil(&myStruct{}, util.InvalidArgument, "error message")
func ValueNotNil(data interface{}, errorType util.ErrorType, message string) *valueNotNil {
	return &valueNotNil{data: data, errorType: errorType, message: message}
}

func (r *valueNotNil) Validate() error {
	if r.data == nil {
		return util.Error(util.InvalidArgument, r.message)
	}
	return nil
}

type boolTrue struct {
	value     bool
	errorType util.ErrorType
	message   string
}

//BoolTrue returns an error if the value is false
func BoolTrue(value bool, errorType util.ErrorType, message string) *boolTrue {
	return &boolTrue{value: value, errorType: errorType, message: message}
}

func (r *boolTrue) Validate() error {
	if r.value == false {
		return util.Error(r.errorType, r.message)
	}
	return nil
}

// BoolNotFalse returns an error if the value is false.
// Deprecated: Use BoolTrue() instead.
func BoolNotFalse(value bool, errorType util.ErrorType, message string) *boolTrue {
	return BoolTrue(value, errorType, message)
}

type boolFalse struct {
	value     bool
	errorType util.ErrorType
	message   string
}

//BoolFalse returns an error if the value is true
func BoolFalse(value bool, errorType util.ErrorType, message string) *boolFalse {
	return &boolFalse{value: value, errorType: errorType, message: message}
}

func (r *boolFalse) Validate() error {
	if r.value {
		return util.Error(r.errorType, r.message)
	}
	return nil
}
