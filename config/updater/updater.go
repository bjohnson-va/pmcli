package configUpdater

import (
	"encoding/json"
	"fmt"

	"github.com/bjohnson-va/pmcli/config"
)

type Updater struct {
	inMemoryConfig config.File
}

func (u *Updater) UpdateAndRestart(d config.File) error {
	s, err := json.Marshal(d)
	if err != nil {
		return fmt.Errorf("failed to marshal: %s", err.Error())
	}
	fmt.Println(string(s))
	return nil
}

func NewUpdater(c config.File) *Updater {
	return &Updater{
		inMemoryConfig: c,
	}
}
