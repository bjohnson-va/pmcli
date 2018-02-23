package datadogapi

import (
	"fmt"
	"bytes"
	"encoding/json"
	"net/http"
	"io/ioutil"
	"strconv"
)

const apiKey string = "691e8c11282932c13f9d94555e61a6d2"

// AlertType one of error, warning, info, or success
type AlertType int64

// AlertType enums
const (
	Error   AlertType = iota
	Warning
	Info
	Success
)

func (a AlertType) String() string {
	switch a {
	case Error:
		return "error"
	case Warning:
		return "warning"
	case Info:
		return "info"
	case Success:
		return "success"
	}
	panic("Unable to determine AlertType string.")
}

// EventData the json structure for sending an event to datadog
type EventData struct {
	Title     string   	`json:"title"`
	Text      string   	`json:"text"`
	Priority  string   	`json:"priority"`
	Tags      []string 	`json:"tags"`
	AlertType string     	`json:"alert_type"`
}

// PushEventToDatadog Sends an event to datadogapi
func PushEventToDatadog(version string, projectName string, environment string, alertType AlertType) {
	url := "https://app.datadoghq.com/api/v1/events/?api_key=" + apiKey
	tags := []string{
		"deploy:" + environment,
		"project:" + projectName,
		"version:" + version,
	}
	data := &EventData{
		Title:     fmt.Sprintf("%s has been deployed to %s", version, projectName),
		Text:      "",
		Priority:  "normal",
		Tags:      tags,
		AlertType: alertType.String(),
	}
	buffer := new(bytes.Buffer)
	err := json.NewEncoder(buffer).Encode(data)
	if err != nil {
		fmt.Println(err)
		return
	}

	request, err := http.NewRequest("POST", url, buffer)
	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	response, err := client.Do(request)
	if response.StatusCode != 202 {
		fmt.Println("failed to send deploy event to datadog")
		bodyBytes, _ := ioutil.ReadAll(response.Body)
		fmt.Printf("response: %s\n", string(bodyBytes))
		fmt.Printf("response_code: %s\n", strconv.Itoa(response.StatusCode))
	}
	if err != nil {
		fmt.Println(err)
		return
	}
}