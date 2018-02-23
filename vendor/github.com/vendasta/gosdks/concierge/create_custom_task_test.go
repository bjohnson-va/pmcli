package concierge

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vendasta/gosdks/basesdk"
	"context"
	"fmt"
)

var (
	sampleTaskId = "TK-1473D493E3F5C6F4590822FB5994C4C261E3721C"
	sampleCreateCustomTaskResponse = fmt.Sprintf(`{"version": "2.0", "data": {"taskId": "%s"}, "requestId": "59b6f84100ff0d891fec90ab3d0001737e61726d2d746573740001636f6e74696e756f757300010126", "responseTime": 687, "statusCode": 200}`, sampleTaskId)
)

func mkBasicTask() *CustomTask {
	return &CustomTask{
		Name:       "Build a Website",
		PartnerID:  "ABC",
		BusinessID: "AG-1234",
	}
}

func Test_CreateCustomTaskValidatesRequiredPartnerId(t *testing.T) {
	task := mkBasicTask()
	task.PartnerID = ""
	_, err := noHTTPConciergeClient.CreateCustomTask(context.Background(), task)
	assert.EqualError(t, err, "Tasks require a PartnerID")
}

func Test_CreateCustomTaskValidatesRequiredBusinessId(t *testing.T) {
	task := mkBasicTask()
	task.BusinessID = ""
	_, err := noHTTPConciergeClient.CreateCustomTask(context.Background(), task)
	assert.EqualError(t, err, "Tasks require a BusinessID")
}

func Test_CreateCustomTaskValidatesRequiredName(t *testing.T) {
	task := mkBasicTask()
	task.Name = ""
	_, err := noHTTPConciergeClient.CreateCustomTask(context.Background(), task)
	assert.EqualError(t, err, "Tasks require a Name")
}

func Test_CreateCustomTaskReturnsErrorIfGetReturnsError(t *testing.T) {
	baseClient := &basesdk.BaseClientMock{Error: fmt.Errorf("Whoops")}
	client := conciergeClient{SDKClient: baseClient}
	task := mkBasicTask()
	_, err := client.CreateCustomTask(context.Background(), task)
	assert.EqualError(t, err, "Whoops")
}

func Test_CreateCustomTaskReturnsTaskId(t *testing.T) {
	baseClient := &basesdk.BaseClientMock{JSONBody: sampleCreateCustomTaskResponse}
	client := conciergeClient{SDKClient: baseClient}
	task := mkBasicTask()
	taskId, err := client.CreateCustomTask(context.Background(), task)
	assert.Nil(t, err)
	assert.Equal(t, sampleTaskId, taskId)
}
