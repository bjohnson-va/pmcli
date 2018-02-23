package concierge

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/vendasta/gosdks/basesdk"
	"github.com/vendasta/gosdks/util"
	"github.com/vendasta/gosdks/validation"
	"io"
)

func (c *conciergeClient) CreateCustomTask(ctx context.Context, task *CustomTask) (string, error) {
	err := validation.NewValidator().Rule(
		validation.StringNotEmpty(task.PartnerID, util.InvalidArgument, "Tasks require a PartnerID"),
		validation.StringNotEmpty(task.BusinessID, util.InvalidArgument, "Tasks require a BusinessID"),
		validation.StringNotEmpty(task.Name, util.InvalidArgument, "Tasks require a Name"),
	).Validate()
	if err != nil {
		return "", err
	}

	taskJson, err := json.Marshal(task)
	if err != nil {
		return "", err
	}
	params := map[string]interface{}{}
	err = json.Unmarshal(taskJson, &params)
	if err != nil {
		return "", err
	}

	response, err := c.SDKClient.Post(ctx, createCustomTaskPath, params, basesdk.Idempotent())
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	taskId, err := taskIdFromResponseBody(response.Body)
	if err != nil {
		return "", err
	}
	return taskId, nil
}

func taskIdFromResponseBody(body io.Reader) (string, error) {
	type CreateCustomTaskData struct {
		TaskId string `json:"taskId"`
	}

	type CreateCustomTaskResponse struct {
		Data CreateCustomTaskData `json:"data"`
	}
	res := &CreateCustomTaskResponse{}
	if err := json.NewDecoder(body).Decode(res); err != nil {
		return "", fmt.Errorf("Failed to find taskId in response: %s" + err.Error())
	}
	return res.Data.TaskId, nil
}
