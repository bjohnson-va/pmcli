package group

import (
	"errors"
)

var (
	// ErrPathNodesRequired is used when path nodes are required but not provided
	ErrPathNodesRequired = errors.New("path nodes are required")
	// ErrPartnerIDRequired is used when the foreign key partner id is required but not provided
	ErrPartnerIDRequired = errors.New("foreign key partner id is required")
)
