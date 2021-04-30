package configUpdater

import (
	"fmt"

	"github.com/bjohnson-va/pmcli/config"
)

func applyMutation(cfgCopy config.File, m config.Mutation) *config.File {
	for key, value := range m.ConfigMap.Instructions {
		fmt.Printf("Setting inMemoryConfig.ConfigMap.Instructions[%s] to %s\n", key, value)
		old := cfgCopy.ConfigMap.Instructions[key]
		if value.StatusCode > 0 {
			old.StatusCode = value.StatusCode
		}
		if value.DelaySecs > 0 {
			old.DelaySecs = value.DelaySecs
		}
		if value.EmptyBody > config.EmptyBody_Unset {
			old.EmptyBody = value.EmptyBody
		}
		cfgCopy.ConfigMap.Instructions[key] = old
	}
	for key, value := range m.ConfigMap.Exclusions {
		cfgCopy.ConfigMap.Exclusions[key] = value
	}
	for key, value := range m.ConfigMap.Overrides {
		cfgCopy.ConfigMap.Overrides[key] = value
	}
	return &cfgCopy
}
