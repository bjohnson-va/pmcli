package serverconfig

import (
	"fmt"
	"net"
	"net/http"

	"golang.org/x/net/context"

	"github.com/vendasta/gosdks/logging"
)

// StartHTTPServer bootstraps an http server and blocks
func StartHTTPServer(healthz http.HandlerFunc, port int, mux *http.ServeMux) error {
	httpSrv, err := buildServer(healthz, port, mux)
	if err != nil {
		return fmt.Errorf("Unable to prepare server prior to serving: %s", err.Error())
	}
	lis, err := buildListener(port)
	if err != nil {
		return fmt.Errorf("Unable to prepare net.listener prior to serving: %s", err.Error())
	}
	return httpSrv.Serve(*lis)
}

func buildServer(healthz http.HandlerFunc, port int, mux *http.ServeMux) (*http.Server, error) {
	httpSrv := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: logging.HTTPMiddleware(mux),
	}
	mux.Handle("/healthz", healthz)
	return httpSrv, nil
}

func buildListener(port int) (*net.Listener, error) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port));
	if err != nil {
		logging.Errorf(context.Background(), "Error creating HTTP listening socket: %s", err.Error())
		return nil, err
	}
	return &lis, nil
}


