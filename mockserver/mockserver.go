package mockserver

import (
	"context"
	"fmt"
	"net/http"
	"github.com/bjohnson-va/pmcli/config"
	"github.com/bjohnson-va/pmcli/handlers"
	"github.com/fsnotify/fsnotify"
	"github.com/vendasta/gosdks/logging"
	"os"
	"github.com/bjohnson-va/pmcli/random"
)

// TODO: Helper for generating initial config file

func BuildAndRun(mockServerPort int64, allowedOrigin string,
	rootDir string, configFile string, randomSeed string) error {
	ctx := context.Background()

	d := serverDetails{
		port:           mockServerPort,
		allowedOrigin:  allowedOrigin,
		rootDir:        rootDir,
		configFilePath: configFile,
		randomSeed: randomSeed,
	}
	err := runServerInBackgroundAndRestartOnConfigFileChanges(ctx, d)
	if err != nil {
		return fmt.Errorf("couldn't start server: %s", err.Error())
	}
	return nil
}

type serverDetails struct {
	port           int64
	allowedOrigin  string
	rootDir        string
	configFilePath string
	randomSeed     string
}

func runServerInBackgroundAndRestartOnConfigFileChanges(ctx context.Context, d serverDetails) error {

	srv, err := prepareServerFromConfig(ctx, d)
	if err != nil {
		return fmt.Errorf("problem preparing server: %s", err.Error())
	}
	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			logging.Errorf(ctx, "Error on ListenAndServe: %s", err.Error())
		}
	}()
	err = startNewServerOnConfigFileChanges(ctx, srv, d)
	if err != nil {
		return fmt.Errorf("unable to start server: %s", err.Error())
	}
	return nil
}

func prepareServerFromConfig(ctx context.Context, d serverDetails) (*http.Server, error) {
	cfg, err := config.ReadFile(d.configFilePath)
	if err != nil {
		return nil, fmt.Errorf("error reading config [%s]: %s", d.configFilePath, err.Error())
	}
	port := determinePortNumber(d, cfg)
	mux, err := buildServerMux(ctx, d, cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to build server mux: %s", err.Error())
	}
	httpSrv := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: mux,
	}
	logging.Infof(ctx, "Ready to serve on port %d...", port)
	return httpSrv, nil
}

func determinePortNumber(d serverDetails, cfg *config.File) int64 {
	port := d.port
	if port == -1 {
		port = cfg.Port
		if port == -1 {
			port = 29000
		}
	}
	return port
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
				if srv != nil {
					srv.Shutdown(ctx)
				}
				err := runServerInBackgroundAndRestartOnConfigFileChanges(ctx, d)
				watcher.Close()
				if err != nil {
					logging.Infof(ctx, "Error starting new server after config change: %s", err.Error())
					logging.Infof(ctx, "Will retry on further changes")
					e := startNewServerOnConfigFileChanges(ctx, nil, d)
					if e != nil {
						logging.Errorf(ctx, "Unable to start file watcher: %s", e.Error())
					}
				}

			case err := <-watcher.Errors:
				logging.Errorf(ctx, "Error watching for config file changes: %s", err.Error())
			}
		}
	}()

	err = watcher.Add(d.configFilePath)
	if err != nil {
		if _, err := os.Stat(d.configFilePath); os.IsNotExist(err) {
			lm := "Cannot watch file because it doesn't exist: %s"
			logging.Warningf(ctx, lm, d.configFilePath)
		} else {
			return fmt.Errorf("error binding watcher to file (%s): %s",
				d.configFilePath, err.Error())
		}
	}
	<-done
	return nil
}

func buildServerMux(ctx context.Context, d serverDetails, configs *config.File) (
	*http.ServeMux, error) {

	r := initializeRandomProvider(ctx, d.randomSeed)

	ao := configs.AllowedOrigin
	if ao == "" {
		ao = d.allowedOrigin
	}

	mux := http.NewServeMux()
	hbc := handlers.HandlerBuildingConfig{
		AllowedOrigin:     ao,
		ProtofileRootPath: d.rootDir,
		AllConfig:         configs.ConfigMap,
		RandomProvider:    &r,
	}
	for _, p := range configs.ProtofileNames {
		h, err := handlers.FromProtofile(hbc, p)
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

func initializeRandomProvider(ctx context.Context, randomSeed string) random.FieldProvider {
	fallback := "breadcrumb"
	switch randomSeed {
	case "breadcrumb":
		return random.BreadcrumbBasedFieldProvider()
	case "time":
		return random.TimeBasedFieldProvider()
	default:
		w := "Unexpected random seed type \"%s\". Falling back to \"%s\""
		logging.Warningf(ctx, w, randomSeed, fallback)
		return random.BreadcrumbBasedFieldProvider()
	}
}
