package configUpdater

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/bjohnson-va/pmcli/config"
	"github.com/bjohnson-va/pmcli/server"
)

type ServerHandle interface {
	CreateNewServer(ctx context.Context, d server.Details, cfg config.File) (server.Definition, error)
}

type Updater struct {
	inMemoryConfig *config.File
	unsavedChanges bool
	activeServer   server.Definition
	servers        ServerHandle
	serverDetails  server.Details
}

func (u *Updater) hasUnsavedChanged() bool {
	return u.unsavedChanges
}

func (u *Updater) UpdateAndRestart(ctx context.Context, m config.Mutation) error {
	u.inMemoryConfig = applyMutation(*u.inMemoryConfig, m)
	u.unsavedChanges = true
	err := u.activeServer.Shutdown(ctx)
	if err != nil {
		return fmt.Errorf("unable to stop existing servers: %s", err.Error())
	}

	u.activeServer, err = u.servers.CreateNewServer(
		ctx, u.serverDetails, *u.inMemoryConfig,
	)
	if err != nil {
		return fmt.Errorf("could not create new server: %s", err.Error())
	}

	go func() {
		err := u.activeServer.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			fmt.Printf("Could not start servers: %s", err.Error())
			os.Exit(1)
		}
	}()

	return nil
}

func (u *Updater) SaveChangesToDisk() error {
	_, err := json.Marshal(u.inMemoryConfig)
	if err != nil {
		return fmt.Errorf("failed to marshal: %s", err.Error())
	}
	// TODO: Implement "save config"
	return fmt.Errorf("not implemented")
}

func NewUpdater(d server.Details, initialServer server.Definition, c config.File, srv ServerHandle) *Updater {
	return &Updater{
		activeServer:   initialServer,
		inMemoryConfig: &c,
		servers:        srv,
		serverDetails:  d,
	}
}
