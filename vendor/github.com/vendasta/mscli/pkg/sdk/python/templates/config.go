package templates

import (
	"fmt"
	"strings"
)

// Config Stores the data used by the ConfigTemplate
type Config struct {
	Environments []Environment
}

// Environment information necessary to call the correct environment
type Environment struct {
	Host           string
	Scope          string
	ServiceAccount string
	URL            string
	Environment    string
}

// GetVaxEnv gets the vax environment
func (e *Environment) GetVaxEnv() string {
	return fmt.Sprintf("Environment.%s", strings.ToUpper(e.Environment))
}

// ConfigTemplate is used to generate the Python Config file
const ConfigTemplate = `""" vax environments """
from vax.environment import Environment

ENVIRONMENT_PARAMS = {
    {{ range .Environments }}{{.GetVaxEnv}}: {
        'host': '{{.Host}}',
        'scope': '{{.Scope}}',
        'service_account': '{{.ServiceAccount}}',
        'url': '{{.URL}}',
    },
    {{end}}
}
`
