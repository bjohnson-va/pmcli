package mockserver

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net"
	"os"
	"github.com/bjohnson-va/pmcli/handlers"
	"github.com/bjohnson-va/pmcli/protofiles"
	"github.com/fsnotify/fsnotify"
	"github.com/vendasta/gosdks/logging"
	"github.com/vendasta/mscli/pkg/spec"
)

// TODO: Helper for generating initial config file

func BuildAndRun(file spec.MicroserviceFile, mockServerPort int,
	allowedOrigin string, rootDir string, configFile string) error {
	ctx := context.Background()

	d := serverDetails{
		microServiceConfig: file.Microservice,
		port:               mockServerPort,
		allowedOrigin:      allowedOrigin,
		rootDir:            rootDir,
		configFilePath:     configFile,
	}
	runServerInBackgroundAndRestartOnConfigFileChanges(ctx, d)
	return nil
}

type serverDetails struct {
	microServiceConfig spec.MicroserviceConfig
	port               int
	allowedOrigin      string
	rootDir            string
	configFilePath     string
}

func runServerInBackgroundAndRestartOnConfigFileChanges(ctx context.Context, d serverDetails) error {

	srv, err := prepareServerFromConfig(ctx, d)
	if err != nil {
		return fmt.Errorf("problem preparing server: %s", err.Error())
	}
	go func() {
		logging.Infof(ctx, "Running HTTP server on port %d...", d.port)
		srv.ListenAndServe()
	}()
	startNewServerOnConfigFileChanges(ctx, srv, d)
	return nil
}

func prepareServerFromConfig(ctx context.Context, d serverDetails) (*http.Server, error) {
	cfg, err := readConfigFile(d.configFilePath)
	if err != nil {
		return nil, fmt.Errorf("error reading config: %s %s", d.configFilePath, err.Error())
	}

	mux, err := buildServerMux(ctx, d, cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to build server mux: %s", err.Error())
	}
	httpSrv := &http.Server{
		Addr:    fmt.Sprintf(":%d", d.port),
		Handler: mux,
	}
	return httpSrv, nil

}

func readConfigFile(filename string) (interface{}, error) {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return make(map[string]interface{}), nil
	}

	plan, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("unable to read config file: %s", err.Error())
	}

	var data interface{}
	err = json.Unmarshal(plan, &data)
	if err != nil {
		return nil, fmt.Errorf("unable to unmarshal config: %s", err.Error())
	}

	return data, nil
}

func startNewServerOnConfigFileChanges(ctx context.Context, srv *http.Server, d serverDetails) error {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return fmt.Errorf("unable to create new watcher: %s", err.Error())
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				logging.Infof(ctx, "Detected config file changes. Restarting server")
				if event.Op&fsnotify.Write == fsnotify.Write {
					logging.Infof(ctx, "modified file: %s", event.Name)
				}
				srv.Shutdown(ctx)
				err := runServerInBackgroundAndRestartOnConfigFileChanges(ctx, d)
				if err != nil {
					logging.Errorf(ctx, "Error starting new server after config change: %s", err.Error())
				}
				watcher.Close()

			case err := <-watcher.Errors:
				logging.Errorf(ctx, "Error watching for config file changes: %s", err.Error())
			}
		}
	}()

	err = watcher.Add(d.configFilePath)
	if err != nil {
		return fmt.Errorf("error binding watcher to file (%s): %s",
			d.configFilePath, err.Error())
	}
	<-done
	return nil
}

func buildServerMux(ctx context.Context, d serverDetails, config interface{}) (
	*http.ServeMux, error) {
	protofileNames, err := protofiles.GetNames(d.microServiceConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to get protofile names: %s", err.Error())
	}
	mux := http.NewServeMux()
	for _, p := range protofileNames {
		logging.Infof(ctx, "Building endpoints for: %s/%s", d.rootDir, p)
		h, err := handlers.FromProtofile(d.allowedOrigin, d.rootDir, p, config)
		if err != nil {
			return nil, fmt.Errorf("failed to make handlers: %s", err.Error())
		}

		for _, handler := range h {
			logging.Infof(ctx, "Created handler: %s\n", handler.Path)
			mux.HandleFunc(handler.Path, handler.HandlerFunc)
		}
	}
	return mux, nil
}

func runServer(port int, mux *http.ServeMux) error {
	httpSrv := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: mux,
	}
	lis, err := buildListener(port)
	if err != nil {
		return fmt.Errorf("unable to prepare net.listener: %s", err.Error())
	}
	err = httpSrv.Serve(*lis)
	if err != nil {
		return fmt.Errorf("unable to start HTTP server %s", err.Error())
	}
	return nil
}

func buildListener(port int) (*net.Listener, error) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return nil, err
	}
	return &lis, nil
}
