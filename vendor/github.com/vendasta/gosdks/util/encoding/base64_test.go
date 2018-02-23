package encoding

import (
	"testing"
	"math"
	"github.com/stretchr/testify/assert"
	"errors"
)

func TestEncodeIntegerToBase64(t *testing.T) {
	type testCase struct {
		in int64
		out string
		name string
	}

	cases := []*testCase {
		{
			in: 0,
			out: "MA==",
			name: "0",
		},
		{
			in: 10,
			out: "MTA=",
			name: "10 (common input value)",
		},
		{
			in: -10,
			out: "LTEw",
			name: "Negative input value",
		},
		{
			in: math.MaxInt64,
			out: "OTIyMzM3MjAzNjg1NDc3NTgwNw==",
			name: "Max input value",
		},
		{
			in: math.MinInt64,
			out: "LTkyMjMzNzIwMzY4NTQ3NzU4MDg=",
			name: "Min input value",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T){
			assert.Equal(t, c.out, EncodeIntegerToBase64(c.in))
		})
	}
}

func TestDecodeBase64ToInt(t *testing.T) {
	type testCase struct {
		in string
		out int64
		err error
		name string
	}

	cases := []*testCase {
		{
			in: "",
			out: 0,
			name: "0 is Empty string",
		},
		{
			in: "MA==",
			out: 0,
			name: "0",
		},
		{
			in: "MTA=",
			out: 10,
			name: "10 (common input value)",
		},
		{
			in: "LTEw",
			out: -10,
			name: "Negative input value",
		},
		{
			in: "OTIyMzM3MjAzNjg1NDc3NTgwNw==",
			out: math.MaxInt64,
			name: "Max input value",
		},
		{
			in: "LTkyMjMzNzIwMzY4NTQ3NzU4MDg=",
			out: math.MinInt64,
			name: "Min input value",
		},
		{
			in: "Ecng",
			err: errors.New(DecodingFailure),
			name: "Incorrect encoding format (500 in base62)",
		},
		{
			in: "potatoessss",
			err: errors.New(DecodingFailure),
			name: "Incorrect encoding format (idiot format)",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T){
			actual, err := DecodeBase64ToInt(c.in)
			assert.Equal(t, c.out, actual)
			assert.Equal(t, c.err, err)
		})
	}
}
