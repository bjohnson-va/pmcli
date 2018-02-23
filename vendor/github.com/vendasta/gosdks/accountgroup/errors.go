package accountgroup

import (
	"errors"
)

var (
	// ErrInvalidCreate is used when invalid data has been passed to create.
	ErrInvalidCreate = errors.New("NAP and External Identifiers must be provided")
)
