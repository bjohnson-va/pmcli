package datadog

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"strings"

	"github.com/spf13/viper"
	"github.com/vendasta/mscli/pkg/spec"
	"github.com/zorkian/go-datadog-api"
)

var client *datadog.Client

type specForDatadog struct {
	spec.MicroserviceConfig
}

// DatadogName returns the name normalized for datadog metrics (dashes replaced by underscores)
func (s specForDatadog) DatadogName() string {
	return strings.Replace(s.Name, "-", "_", -1)
}

// EnsureDashboard for microservice spec
func EnsureDashboard(spec spec.MicroserviceFile, force bool) error {
	d, err := client.GetDashboard(spec.Microservice.DatadogDashboardID)

	// Errors except 404s
	if err != nil && !strings.Contains(err.Error(), "404 Not Found") {
		return fmt.Errorf("Failed to get Dashboard: %s", err.Error())
	}

	if d != nil {
		if !force {
			return fmt.Errorf("Dashboard already exists. Use --force if you want to overwrite")
		} else {
			fmt.Println("Force overwriting dashboard")
		}
	}

	dToCreate := dashboard(specForDatadog{spec.Microservice})
	if spec.Microservice.Debug {
		bs, err := json.Marshal(dToCreate)
		if err != nil {
			return fmt.Errorf("Failed to marshal Dashboard to create: %s", err.Error())
		}
		fmt.Println("Creating Dashboard with JSON:")
		fmt.Println(string(bs))
	}

	newD, err := client.CreateDashboard(dToCreate)
	if err != nil {
		return fmt.Errorf("Failed to Create Dashboard: %s:", err.Error())
	}

	spec.Microservice.DatadogDashboardID = newD.GetId()
	viper.Set("microservice", spec.Microservice)
	viper.WriteConfig()
	fmt.Println("Wrote Dashboard ID to config")

	return nil
}

func dashboard(spec specForDatadog) *datadog.Dashboard {
	d := datadog.Dashboard{}
	d.SetTitle(strings.Title(spec.Name + " Service Timeboard"))
	d.SetDescription(strings.Title(spec.Name))
	d.TemplateVariables = templateVariables(spec)
	d.Graphs = graphs(spec)
	return &d
}

func templateVariables(spec specForDatadog) []datadog.TemplateVariable {
	app := datadog.TemplateVariable{
		Name:    sp("Application"),
		Prefix:  sp("namespace"),
		Default: tsp("[[.Name]]-prod", spec),
	}
	kapp := datadog.TemplateVariable{
		Name:    sp("KApplication"),
		Prefix:  sp("kube_namespace"),
		Default: tsp("[[.Name]]-prod", spec),
	}

	return []datadog.TemplateVariable{app, kapp}
}

func graphs(spec specForDatadog) []datadog.Graph {
	nfthLatencyVal := datadog.Graph{
		Title: sp("95th Percentile Latency"),
		Definition: &datadog.GraphDefinition{
			Requests: []datadog.GraphDefinitionRequest{
				{
					Query: tsp("avg:[[.DatadogName]].gRPC.Latency.95percentile{$Application}", spec),
				},
			},
			Viz:        sp("query_value"),
			Precision:  sp("0"),
			CustomUnit: sp("ms"),
		},
	}
	rps := datadog.Graph{
		Title: sp("Requests Per Second"),
		Definition: &datadog.GraphDefinition{
			Requests: []datadog.GraphDefinitionRequest{
				{
					Query: tsp("sum:[[.DatadogName]].gRPC{$Application}.as_rate()", spec),
				},
			},
			Viz:        sp("query_value"),
			Precision:  sp("2"),
			CustomUnit: sp("req/s"),
		},
	}
	avail := datadog.Graph{
		Title: sp("Availability"),
		Definition: &datadog.GraphDefinition{
			Requests: []datadog.GraphDefinitionRequest{
				{
					Query: tsp("sum:[[.DatadogName]].gRPC.Latency.count{status:200,$Application}.as_count() / ( sum:[[.DatadogName]].gRPC.Latency.count{status:200,$Application}.as_count() + sum:[[.DatadogName]].gRPC.Latency.count{status:500,$Application}.as_count() ) * 100", spec),
				},
			},
			Viz:        sp("query_value"),
			Precision:  sp("3"),
			CustomUnit: sp("%"),
		},
	}
	fivehundreds := datadog.Graph{
		Title: sp("500's"),
		Definition: &datadog.GraphDefinition{
			Requests: []datadog.GraphDefinitionRequest{
				{
					Query: tsp("sum:[[.DatadogName]].gRPC.Latency.count{status:500,$Application} by {path}.as_count()", spec),
					Type:  sp("bars"),
				},
			},
			Viz: sp("timeseries"),
			Events: []datadog.GraphEvent{
				{Query: tsp("tags:deploy:prod,project:[[.Name]]", spec)},
			},
			Style: &datadog.Style{
				Palette: sp("warm"),
			},
		},
	}
	fourhundreds := datadog.Graph{
		Title: sp("400's"),
		Definition: &datadog.GraphDefinition{
			Requests: []datadog.GraphDefinitionRequest{
				{
					Query: tsp("sum:[[.DatadogName]].gRPC.Latency.count{status:400,$Application} by {path}.as_count()", spec),
					Type:  sp("bars"),
				},
			},
			Viz: sp("timeseries"),
		},
	}
	overallTraffic := datadog.Graph{
		Title: sp("Overall Traffic"),
		Definition: &datadog.GraphDefinition{
			Requests: []datadog.GraphDefinitionRequest{
				{
					Query: tsp("sum:[[.DatadogName]].gRPC.Latency.count{$Application}.as_count()", spec),
					Type:  sp("bars"),
				},
			},
			Events: []datadog.GraphEvent{
				{Query: tsp("tags:deploy:prod,project:[[.Name]]", spec)},
			},
			Viz: sp("timeseries"),
		},
	}
	nfthLatency := datadog.Graph{
		Title: sp("95th Latency"),
		Definition: &datadog.GraphDefinition{
			Requests: []datadog.GraphDefinitionRequest{
				{
					Query: tsp("default(avg:[[.DatadogName]].gRPC.Latency.95percentile{$Application}, 0)", spec),
					Type:  sp("line"),
				},
			},
			Events: []datadog.GraphEvent{
				{Query: tsp("tags:deploy:prod,project:[[.Name]]", spec)},
			},
			Viz:        sp("timeseries"),
			Precision:  sp("0"),
			CustomUnit: sp("ms"),
		},
	}

	highest500s := datadog.Graph{
		Title: sp("Highest 500's"),
		Definition: &datadog.GraphDefinition{
			Requests: []datadog.GraphDefinitionRequest{
				{
					Query: tsp("top(sum:[[.DatadogName]].gRPC.Latency.count{status:500,$Application} by {path}.as_count(), 10, 'sum', 'desc')", spec),
				},
			},
			Viz: sp("toplist"),
			Style: &datadog.Style{
				Palette: sp("warm"),
			},
		},
	}
	highestActivity := datadog.Graph{
		Title: sp("Highest Activity"),
		Definition: &datadog.GraphDefinition{
			Requests: []datadog.GraphDefinitionRequest{
				{
					Query: tsp("top(sum:[[.DatadogName]].gRPC.Latency.count{$Application} by {path}.as_count(), 20, 'sum', 'desc')", spec),
				},
			},
			Viz: sp("toplist"),
		},
	}
	highestLatency := datadog.Graph{
		Title: sp("Highest Latency"),
		Definition: &datadog.GraphDefinition{
			Requests: []datadog.GraphDefinitionRequest{
				{
					Query: tsp("top(avg:[[.DatadogName]].gRPC.Latency.95percentile{$Application} by {path}, 20, 'mean', 'desc')", spec),
				},
			},
			Viz: sp("toplist"),
			Style: &datadog.Style{
				Palette: sp("warm"),
			},
		},
	}
	hpa := datadog.Graph{
		Title: sp("HPA, Desired/Max"),
		Definition: &datadog.GraphDefinition{
			Requests: []datadog.GraphDefinitionRequest{
				{
					Query: tsp("avg:kubernetes_state.hpa.max_replicas{$Application}", spec),
				},
				{
					Query: tsp("avg:kubernetes_state.hpa.min_replicas{$Application}", spec),
				},
				{
					Query: tsp("avg:kubernetes_state.hpa.current_replicas{$Application}", spec),
					Type: sp("bars"),
				},
			},
			Viz: sp("timeseries"),
			Style: &datadog.Style{
				Palette: sp("warm"),
			},
		},
	}
	unavailablePods := datadog.Graph{
		Title: sp("Pods Unavailable"),
		Definition: &datadog.GraphDefinition{
			Requests: []datadog.GraphDefinitionRequest{
				{
					Query: tsp("sum:kubernetes_state.deployment.replicas_unavailable{$Application}", spec),
					Type: sp("area"),
					Style: &datadog.GraphDefinitionRequestStyle{
						Palette: sp("orange"),
					},
				},
			},
			Viz: sp("timeseries"),
		},
	}

	return []datadog.Graph{
		nfthLatencyVal, rps, avail,
		fivehundreds, fourhundreds, overallTraffic, nfthLatency,
		highest500s, highestActivity, highestLatency, hpa, unavailablePods}
}

func init() {
	client = datadog.NewClient("691e8c11282932c13f9d94555e61a6d2", "cee7a473ff5f3b9e63cd22b14c553748ff86ff3c")
}

// sp is a helper for string pointers
func sp(s string) *string {
	return &s
}

// tsp is a helper for spec templated string pointers
func tsp(t string, spec specForDatadog) *string {
	// Template JSON
	tmpl, err := template.New("t").Delims("[[", "]]").Parse(t)
	if err != nil {
		panic(err) // TODO: Handle more gracefully?
	}
	buff := bytes.NewBufferString("")
	tmpl.Execute(buff, spec)
	so := buff.String()
	return &so
}
