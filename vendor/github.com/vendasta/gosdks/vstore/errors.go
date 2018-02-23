package vstore

import (
	"errors"
	"fmt"
)

var (
	//ErrTooMuchContention is raised when the client detects there is too much contention on an entity being operated on
	ErrTooMuchContention = errors.New("Too much contention on the entity.")
	//ErrUnregisteredModel is raised when the client can not find a vstore.Model to deserialize an entity into
	ErrUnregisteredModel = errors.New("Unable to find model for given namespace. Have you called `vstore.RegisterModel`?")
)

type invalidFieldError struct {
	reason string
}

func (i invalidFieldError) Error() string {
	return i.reason
}
func newInvalidFieldError(format string, a ...interface{}) error {
	return invalidFieldError{reason: fmt.Sprintf(format, a...)}
}
