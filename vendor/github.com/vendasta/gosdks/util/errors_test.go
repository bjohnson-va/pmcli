package util

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"errors"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/codes"
)

func TestToGrpcError(t *testing.T) {
	type testCase struct {
		in error
		out error
		name string
	}

	cases := []*testCase{
		{
			in: errors.New("Unknown!"),
			out: status.Error(codes.Unknown, "Unknown server error."),
			name: "Raw error is unknown",
		},
		{
			in: nil,
			out: nil,
			name: "nil is just nil",
		},
		{
			in: Error(AlreadyExists, "im a proprietary error"),
			out: status.Error(codes.AlreadyExists, "im a proprietary error"),
			name: "Service error is mapped to GRPC error with custom message",
		},
		{
			in: status.Error(codes.DataLoss, "im a GRPC error"),
			out: status.Error(codes.DataLoss, "im a GRPC error"),
			name: "GRPC error is mapped to GRPC error with custom message",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T){
			assert.Equal(t, c.out, ToGrpcError(c.in))
		})
	}
}