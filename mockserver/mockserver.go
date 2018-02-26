package mockserver

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"github.com/bjohnson-va/pmcli/handlers"
	"github.com/fsnotify/fsnotify"
	"github.com/vendasta/gosdks/logging"
)

// TODO: Helper for generating initial config file

func BuildAndRun(mockServerPort int,
	allowedOrigin string, rootDir string, configFile string) error {
	ctx := context.Background()

	d := serverDetails{
		port:           mockServerPort,
		allowedOrigin:  allowedOrigin,
		rootDir:        rootDir,
		configFilePath: configFile,
		protofileNames: []string{"advertising/v1/api.proto"}, // TODO: Read from config file
	}
	err := runServerInBackgroundAndRestartOnConfigFileChanges(ctx, d)
	if err != nil {
		return fmt.Errorf("couldn't start server: %s", err.Error())
	}
	return nil
}

type serverDetails struct {
	port           int
	allowedOrigin  string
	rootDir        string
	configFilePath string
	protofileNames []string
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
	cfg, err := readConfigFile(ctx, d.configFilePath)
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

func readConfigFile(ctx context.Context, filename string) (map[string]interface{}, error) {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		logging.Warningf(ctx, "Config file did not exist: %s", filename)
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

	i, ok := data.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("couldn't turn config into map: %s")
	}
	return i, nil
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

func buildServerMux(ctx context.Context, d serverDetails, config map[string]interface{}) (
	*http.ServeMux, error) {

	mux := http.NewServeMux()
	for _, p := range d.protofileNames {
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
