package tinyid

import "golang.org/x/net/context"

// Interface defines the methods for generating tiny ids.
type Interface interface {
	GenerateTinyID(ctx context.Context) (string, error)
}
