package mockserver

import (
	"bufio"
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/bjohnson-va/pmcli/certs"
	"github.com/bjohnson-va/pmcli/config"
	configUpdater "github.com/bjohnson-va/pmcli/config/updater"
	"github.com/bjohnson-va/pmcli/handlers"
	"github.com/bjohnson-va/pmcli/random"
	"github.com/fsnotify/fsnotify"
	"github.com/vendasta/gosdks/logging"
)

// TODO: Helper for generating initial config file

func BuildAndRun(
	mockServerPort int64, allowedOrigin string,
	rootDir string, configFile string, randomSeed string, interactive bool,
) error {
	ctx := context.Background()

	d := serverDetails{
		port:           mockServerPort,
		allowedOrigin:  allowedOrigin,
		rootDir:        rootDir,
		configFilePath: configFile,
		randomSeed:     randomSeed,
		interactive:    interactive,
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
	interactive    bool
}

func runServerInBackgroundAndRestartOnConfigFileChanges(ctx context.Context, d serverDetails) error {
	srvRes, err := prepareServerFromConfig(ctx, d)
	if err != nil {
		return fmt.Errorf("problem preparing server: %s", err.Error())
	}
	srv := srvRes.server
	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			logging.Errorf(ctx, "Error on ListenAndServe: %s", err.Error())
		}
	}()

	logging.Infof(ctx, "ineractive %t", d.interactive)
	if d.interactive {
		cfg, err := config.ReadFile(d.configFilePath)
		if err != nil {
			return fmt.Errorf("error reading config [%s]: %s", d.configFilePath, err.Error())
		}
		u := configUpdater.NewUpdater(*cfg)
		err = showInteractivePrompts(ctx, srvRes.endpoints, u, d)
	} else {
		err = startNewServerOnConfigFileChanges(ctx, srv, d)
	}
	if err != nil {
		return fmt.Errorf("unable to start server: %s", err.Error())
	}
	return nil
}

type ServerResult struct {
	server    Server
	endpoints []string
}

func prepareServerFromConfig(ctx context.Context, d serverDetails) (ServerResult, error) {
	cfg, err := config.ReadFile(d.configFilePath)
	if err != nil {
		return ServerResult{}, fmt.Errorf("error reading config [%s]: %s", d.configFilePath, err.Error())
	}
	port := determinePortNumber(d, cfg)
	muxRes, err := buildServerMux(ctx, d, cfg)
	if err != nil {
		return ServerResult{}, fmt.Errorf("failed to build server mux: %s", err.Error())
	}
	s := "http"
	if cfg.Https {
		s = "https"
	}
	logging.Infof(ctx, "Ready to serve on %s://localhost:%d...", s, port)

	insecureSrv := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: muxRes.mux,
	}
	var httpSrv Server
	httpSrv = insecureSrv

	schema := "insecure HTTP"
	if cfg.Https {
		schema = "HTTPS"
		httpSrv = AddCertsToServer(insecureSrv, cfg.Port, certs.Certificate, certs.PrivateKey)
	}
	logging.Infof(ctx, "Server will use %s (From %s in %s)", schema, config.UseHTTPSToken, config.FILENAME)

	return ServerResult{
		server:    httpSrv,
		endpoints: muxRes.endpoints,
	}, nil
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

func startNewServerOnConfigFileChanges(ctx context.Context, srv Server, d serverDetails) error {
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

type ServeMuxResult struct {
	mux       *http.ServeMux
	endpoints []string
}

var nilResult = ServeMuxResult{}

func buildServerMux(
	ctx context.Context, d serverDetails, configs *config.File,
) (ServeMuxResult, error) {

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
	logging.Infof(ctx, "Handlers will allow requests from origin: \"%s\"", ao)

	var mockedPaths []string

	for _, p := range configs.ProtofileNames {
		h, err := handlers.FromProtofile(hbc, p)
		if err != nil {
			return nilResult, fmt.Errorf("failed to make handlers: %s", err.Error())
		}

		for _, handler := range h {
			logging.Infof(ctx, "Created handler: %s\n", handler.Path)
			mux.HandleFunc(handler.Path, handler.HandlerFunc)
			mockedPaths = append(mockedPaths, handler.Path)
		}
	}

	buildRootHandler(ctx, mux, mockedPaths)

	return ServeMuxResult{
		mux:       mux,
		endpoints: mockedPaths,
	}, nil
}

func buildRootHandler(ctx context.Context, mux *http.ServeMux, mockedPaths []string) {
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		bw := bufio.NewWriter(w)
		_, err := bw.WriteString("<html><body>This is a mock server. The following RPCs are available:")
		if err != nil {
			logging.Errorf(ctx, "Failed to write to response: %s", err.Error())
		}
		for _, p := range mockedPaths {
			_, err := bw.WriteString(fmt.Sprintf(`<p><a href="%s">%s</a><p>`, p, p))
			if err != nil {
				logging.Errorf(ctx, "Failed to write path to response: %s", err.Error())
			}
		}
		_, err = bw.WriteString("</body></html>")
		if err != nil {
			logging.Errorf(ctx, "Failed to write to response: %s", err.Error())
		}
		err = bw.Flush()
		if err != nil {
			logging.Errorf(ctx, "Failed to flush writer", err.Error())
		}
	})
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
