package mockserver

import (
	"context"

	"github.com/bjohnson-va/pmcli/config"
)

type Server interface {
	ListenAndServe(ctx context.Context) error
	Shutdown(ctx context.Context) error
	SetAssists(a config.AssistEnum)
}
