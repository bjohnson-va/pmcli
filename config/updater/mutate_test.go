package configUpdater

import (
	"testing"

	"github.com/bjohnson-va/pmcli/config"
	"github.com/stretchr/testify/assert"
)

func TestUpdater_ApplyMutation_ShouldApplyMutation(t *testing.T) {
	initialConfig := config.File{
		ConfigMap: config.Map{
			Instructions: map[string]config.RPCInstructions{
				"MyRPC": {
					DelaySecs: 10,
				},
			},
		},
	}
	mutation := config.Mutation{
		ConfigMap: config.Map{
			Instructions: map[string]config.RPCInstructions{
				"MyRPC": {
					DelaySecs: 20,
				},
			},
		},
	}
	newCfg := applyMutation(initialConfig, mutation)
	assert.Equal(t, newCfg.ConfigMap.Instructions["MyRPC"].DelaySecs, 20)
}
func TestUpdater_ApplyMutation_ShouldNotChangeValues_ThatAreNotInMutation(t *testing.T) {
	initialStatusCode := 200
	initialConfig := config.File{
		ConfigMap: config.Map{
			Instructions: map[string]config.RPCInstructions{
				"MyRPC": {
					DelaySecs:  10,
					StatusCode: initialStatusCode,
				},
			},
		},
	}
	mutation := config.Mutation{
		ConfigMap: config.Map{
			Instructions: map[string]config.RPCInstructions{
				"MyRPC": {
					DelaySecs: 20,
				},
			},
		},
	}
	newCfg := applyMutation(initialConfig, mutation)
	assert.Equal(t, newCfg.ConfigMap.Instructions["MyRPC"].StatusCode, initialStatusCode)
}
