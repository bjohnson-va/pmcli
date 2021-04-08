package mockserver

import (
	"context"
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"os/exec"

	"github.com/vendasta/gosdks/logging"

	"github.com/bjohnson-va/pmcli/config"
)

type InsecureServer struct {
	delegate *http.Server
	assists  config.AssistEnum
}

func (i *InsecureServer) SetAssists(a config.AssistEnum) {
	i.assists = a
}

func (i *InsecureServer) ListenAndServe(ctx context.Context) error {
	i.logAssists(ctx)
	return i.delegate.ListenAndServe()
}

func (i *InsecureServer) Shutdown(ctx context.Context) error {
	return i.delegate.Shutdown(ctx)
}

func (i *InsecureServer) Serve(ctx context.Context, listener net.Listener) error {
	i.logAssists(ctx)
	return i.delegate.Serve(listener)
}

func (i *InsecureServer) GetTLSConfig() *tls.Config {
	return i.delegate.TLSConfig
}

func (i *InsecureServer) logAssists(ctx context.Context) {
	switch i.assists {
	case config.AssistUnset:
	case config.AssistAngular:
		mime, err := exec.LookPath("xdg-mime")
		if err != nil {
			logAssistsFailure(ctx, err)
			break
		}
		cmd := exec.Command(fmt.Sprintf("%s query default text/html", mime))
		output, err := cmd.Output()
		if err != nil {
			logAssistsFailure(ctx, err)
			break
		}
		logging.Infof(ctx, "Your browser is %s", output)
	}
}

func logAssistsFailure(ctx context.Context, err error) {
	logging.Errorf(
		ctx, "There was a problem generating helper documentation. Angular assistance will not function: %s",
		err.Error(),
	)
}

func NewInsecureServer(ctx context.Context, d serverDetails, cfg *config.File) (*InsecureServer, error) {
	port := determinePortNumber(d, cfg)
	mux, err := buildServerMux(ctx, d, cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to build server mux: %s", err.Error())
	}
	s := "http"
	if cfg.Https {
		s = "https"
	}
	logging.Infof(ctx, "Ready to serve on %s://localhost:%d...", s, port)
	return &InsecureServer{
		delegate: &http.Server{
			Addr:    fmt.Sprintf(":%d", cfg.Port),
			Handler: mux,
		},
	}, nil
}
