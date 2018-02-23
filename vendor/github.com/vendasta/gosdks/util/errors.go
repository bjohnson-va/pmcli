package util

import (
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ErrorType is an enum encapsulating the spectrum of all possible types of errors raised by the application
type ErrorType int64

const (
	// NotFound corresponds to errors caused by missing entities
	NotFound ErrorType = 1 + iota
	// InvalidArgument corresponds to errors caused by missing or malformed arguments supplied by a client
	InvalidArgument
	// AlreadyExists corresponds to errors caused by an entity already existing
	AlreadyExists
	// PermissionDenied corresponds to a user not having permission to access a resource.
	PermissionDenied
	// Unauthenticated indicates the request does not have valid authentication credentials for the operation.
	Unauthenticated
	// Unimplemented corresponds to a function that is unimplemented
	Unimplemented
	// Unknown Error occurred
	Unknown
	// Internal Error
	Internal
	// Unavailable error occurred
	Unavailable
	// FailedPrecondition indicates operation was rejected because the
	// system is not in a state required for the operation's execution.
	// For example, directory to be deleted may be non-empty, an rmdir
	// operation is applied to a non-directory, etc.
	//
	// A litmus test that may help a service implementor in deciding
	// between FailedPrecondition, Aborted, and Unavailable:
	//  (a) Use Unavailable if the client can retry just the failing call.
	//  (b) Use Aborted if the client should retry at a higher-level
	//      (e.g., restarting a read-modify-write sequence).
	//  (c) Use FailedPrecondition if the client should not retry until
	//      the system state has been explicitly fixed.  E.g., if an "rmdir"
	//      fails because the directory is non-empty, FailedPrecondition
	//      should be returned since the client should not retry unless
	//      they have first fixed up the directory by deleting files from it.
	//  (d) Use FailedPrecondition if the client performs conditional
	//      REST Get/Update/Delete on a resource and the resource on the
	//      server does not match the condition. E.g., conflicting
	//      read-modify-write on the same resource.
	FailedPrecondition
)

// ServiceError is an error that can be translated to a GRPC-compliant error
type ServiceError struct {
	msg     string
	errType ErrorType
}

// Error returns the message associated with this error
func (v ServiceError) Error() string {
	return v.msg
}

// ErrorType returns the ErrorType associated with this error
func (v ServiceError) ErrorType() ErrorType {
	return v.errType
}

// GRPCError returns an error in a format such that it can be consumed by GRPC
func (v ServiceError) GRPCError() error {
	if v.errType == NotFound {
		return status.Errorf(codes.NotFound, v.msg)
	} else if v.errType == InvalidArgument {
		return status.Errorf(codes.InvalidArgument, v.msg)
	} else if v.errType == AlreadyExists {
		return status.Errorf(codes.AlreadyExists, v.msg)
	} else if v.errType == PermissionDenied {
		return status.Errorf(codes.PermissionDenied, v.msg)
	} else if v.errType == Unauthenticated {
		return status.Errorf(codes.Unauthenticated, v.msg)
	} else if v.errType == Unimplemented {
		return status.Errorf(codes.Unimplemented, v.msg)
	} else if v.errType == Internal {
		return status.Errorf(codes.Internal, v.msg)
	} else if v.errType == Unavailable {
		return status.Errorf(codes.Unavailable, v.msg)
	} else if v.errType == FailedPrecondition {
		return status.Errorf(codes.FailedPrecondition, v.msg)
	}
	return status.Errorf(codes.Unknown, "Unknown server error.")
}

// HTTPCode returns the corresponding http status code for a given error
func (v ServiceError) HTTPCode() int {
	if v.errType == NotFound {
		return 404
	} else if v.errType == InvalidArgument {
		return 400
	} else if v.errType == AlreadyExists {
		return 409
	} else if v.errType == PermissionDenied {
		return 403
	} else if v.errType == Unauthenticated {
		return 401
	} else if v.errType == Unimplemented {
		return 501
	} else if v.errType == FailedPrecondition {
		return 412
	}
	return 500
}

// Error returns a ServiceError
func Error(errorType ErrorType, format string, a ...interface{}) error {
	return ServiceError{msg: fmt.Sprintf(format, a...), errType: errorType}
}

// FromError given an error tries to return a proper ServiceError.
func FromError(err error) ServiceError {
	serviceError, ok := err.(ServiceError)
	if !ok {
		return Error(Unknown, "Unknown server error.").(ServiceError)
	}
	return serviceError
}

// IsError returns true/false if the given err matches the errorType type.
func IsError(errorType ErrorType, err error) bool {
	serviceError, ok := err.(ServiceError)
	if !ok {
		return false
	}
	return serviceError.errType == errorType
}

// ToGrpcError calculates the correct GRPC error code for a ServiceError or existing GRPC error and returns it
// All errors that are not GRPC errors or ServiceErrors will be interpreted as Unknown errors
func ToGrpcError(err error) error {
	// if this is already a GRPC error, pass through
	grpcErr, ok := status.FromError(err)
	if ok {
		return grpcErr.Err()
	}
	// otherwise map to ServiceError
	return FromError(err).GRPCError()
}

//Convert a http error into a grpc error
func StatusCodeToGRPCError(statusCode int) ErrorType {
	switch statusCode {
	case 400:
		return InvalidArgument
	case 401:
		return Unauthenticated
	case 403:
		return PermissionDenied
	case 404:
		return NotFound
	case 409:
		return AlreadyExists
	case 412:
		return FailedPrecondition
	case 501:
		return Unimplemented
	default:
		return Internal
	}
}
