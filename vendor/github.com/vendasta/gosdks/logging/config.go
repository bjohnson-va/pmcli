package logging

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"sync"

	gce_metadata "cloud.google.com/go/compute/metadata"
	"cloud.google.com/go/logging"
	"google.golang.org/api/option"
)

var configValue *config
var mut sync.Mutex

type config struct {
	ProjectID     string
	Namespace     string
	PodName       string
	AppName       string
	TracingClient *Client
}

func (c *config) BuildHeader(name string) string {
	if c == nil {
		return "local"
	}
	return fmt.Sprintf("x-%s-%s", strings.ToLower(c.AppName), name)
}

// Initialize must be called on app startup and must be done before any logging statements have been issued.
func Initialize(gkeNamespace, podName, appName string) error {
	if gkeNamespace == "" || podName == "" || appName == "" {
		return errors.New("gkeNamespace, podName and appName must be supplied")
	}
	mut.Lock()
	defer mut.Unlock()

	if configValue != nil {
		return errors.New("loggerInstance has already been initialized")
	}
	projectID := appName + "-local"
	if gce_metadata.OnGCE() {
		var err error
		projectID, err = gce_metadata.ProjectID()
		if err != nil {
			return err
		}
	}

	tracingClient, err := NewClient(context.Background(), projectID)
	if err != nil {
		return err
	}

	configValue = &config{
		Namespace:     gkeNamespace,
		PodName:       podName,
		AppName:       appName,
		ProjectID:     projectID,
		TracingClient: tracingClient,
	}
	ctx := context.Background()
	if gce_metadata.OnGCE() {
		client, err := logging.NewClient(ctx, projectID, option.WithGRPCConnectionPool(5))
		if err != nil {
			return err
		}
		loggerInstance, err = newGkeLogger(configValue, client)
		if err != nil {
			return err
		}
	} else {
		loggerInstance, _ = newStdErrLogger(configValue)
	}
	return nil
}
