package taskqueue

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/vendasta/gosdks/config"
	"github.com/vendasta/gosdks/util"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/cloudtasks/v2beta2"
	"google.golang.org/api/googleapi"
)

const defaultProjectID = "repcore-prod"
const defaultLocationID = "us-central1"
const parentTemplate = "projects/{{PROJECT_ID}}/locations/{{LOCATION_ID}}"

var NoTaskToPull = errors.New("no task to pull")

// Interface for the taskqueue client
type Interface interface {
	LeaseTasks(ctx context.Context, queueID string, leaseDuration time.Duration, maxToLease int64, filter string) ([]Task, error)
	LeaseTask(ctx context.Context, queueID string, leaseDuration time.Duration, maxToLease int64, filter string) (Task, error)
	RenewLease(ctx context.Context, t Task, leaseDuration time.Duration) error
	AckTask(ctx context.Context, taskID string, ackID string) error
	NackTask(ctx context.Context, taskID string, ackID string) error
	ScheduleTask(ctx context.Context, queueID string, payload PayloadDefinition, opts ...TaskSchedulingOption) error
}

// client communicates with the Cloud Tasks api
type client struct {
	service     *cloudtasks.Service
	environment config.Env
	projectID   string
	locationID  string
	parentPath  string
}

// NewClient creates a new client for cloud tasks
func NewClient(ctx context.Context, env config.Env, opts ...ClientOption) (Interface, error) {
	cfg := &clientConfig{
		projectID:  defaultProjectID,
		locationID: defaultLocationID,
	}
	for _, o := range opts {
		o(cfg)
	}

	c, err := google.DefaultClient(ctx, cloudtasks.CloudPlatformScope)
	if err != nil {
		return nil, fmt.Errorf("error instantiating google client: %v", err)
	}
	svc, err := cloudtasks.New(c)
	if err != nil {
		return nil, fmt.Errorf("error instantiating service: %v", err)
	}

	pp := strings.Replace(parentTemplate, "{{PROJECT_ID}}", cfg.projectID, 1)
	pp = strings.Replace(pp, "{{LOCATION_ID}}", cfg.locationID, 1)

	return &client{
		service:     svc,
		projectID:   cfg.projectID,
		locationID:  cfg.locationID,
		environment: env,
		parentPath:  pp,
	}, nil
}

func (c *client) queueID(queueID string) string {
	return fmt.Sprintf("%s-%s", queueID, c.environment.Name())
}

func (c *client) queuePath(queueID string) string {
	return c.parentPath + fmt.Sprintf("/queues/%s", c.queueID(queueID))
}

func (c *client) taskPath(queueID, taskID string) string {
	return c.queuePath(queueID) + fmt.Sprintf("/tasks/%s", taskID)
}

// LeaseTasks will lease `maxToLease` number of tasks from the specified queue with the specified `leaseDuration`
// Any tasks that are not acknowledged in the leaseDuration will be released back into the queue.
func (c *client) LeaseTasks(ctx context.Context, queueID string, leaseDuration time.Duration, maxToLease int64, filter string) ([]Task, error) {
	if c.isLocal() {
		return []Task{}, nil
	}

	l := fmt.Sprintf("%ds", int(leaseDuration.Seconds()))
	if maxToLease < 1 {
		maxToLease = 1
	}
	req := &cloudtasks.LeaseTasksRequest{
		Filter:        filter,
		LeaseDuration: l,
		MaxTasks:      maxToLease,
		ResponseView:  "FULL",
	}
	resp, err := c.service.Projects.Locations.Queues.Tasks.Lease(c.queuePath(queueID), req).Context(ctx).Do()
	if err != nil {
		return nil, err
	}
	tasks := make([]Task, len(resp.Tasks))
	for i, t := range resp.Tasks {
		tasks[i], err = taskFromAPI(t)
		if err != nil {
			return nil, err
		}
	}
	return tasks, err
}

// LeaseTask will lease a single task as per LeaseTasks
func (c *client) LeaseTask(ctx context.Context, queueID string, leaseDuration time.Duration, maxToLease int64, filter string) (Task, error) {
	tasks, err := c.LeaseTasks(ctx, queueID, leaseDuration, 1, filter)
	if err != nil {
		return Task{}, err
	}
	if len(tasks) == 0 {
		return Task{}, NoTaskToPull
	}
	return tasks[0], err
}

// RenewLease will renew the lease on a single task
func (c *client) RenewLease(ctx context.Context, t Task, leaseDuration time.Duration) error {
	_, err := c.service.Projects.Locations.Queues.Tasks.RenewLease(t.taskID, &cloudtasks.RenewLeaseRequest{
		LeaseDuration: fmt.Sprintf("%ds", int(leaseDuration.Seconds())),
		ScheduleTime:  t.ackID,
	}).Do()
	return err
}

// AckTask will acknowledge a specific task as being completed. After a task has been acked, it will be removed from the task queue.
func (c *client) AckTask(ctx context.Context, taskID string, ackID string) error {
	_, err := c.service.Projects.Locations.Queues.Tasks.Acknowledge(taskID, &cloudtasks.AcknowledgeTaskRequest{
		ScheduleTime: ackID,
	}).Context(ctx).Do()
	return convertErrorToGRPC(err)
}

// NackTask will cancel the lease on a specific task, which means that task will go back into the queue to be pulled again.
func (c *client) NackTask(ctx context.Context, taskID string, ackID string) error {
	_, err := c.service.Projects.Locations.Queues.Tasks.CancelLease(taskID, &cloudtasks.CancelLeaseRequest{
		ScheduleTime: ackID,
	}).Context(ctx).Do()
	return convertErrorToGRPC(err)
}

// ScheduleTask will place a task into the queue to be processed later.
func (c *client) ScheduleTask(ctx context.Context, queueID string, payload PayloadDefinition, opts ...TaskSchedulingOption) error {
	if c.isLocal() {
		return nil
	}

	o := &taskScheduleConfig{}
	for _, f := range opts {
		f(o)
	}

	serialized, err := json.Marshal(payload)
	if err != nil {
		return errors.New("error serializing payload")
	}

	t := &cloudtasks.Task{
		PullMessage: &cloudtasks.PullMessage{
			Tag:     o.tag,
			Payload: base64.URLEncoding.EncodeToString(serialized),
		},
	}

	if o.scheduleTime != nil {
		t.ScheduleTime = o.scheduleTime.Format(time.RFC3339Nano)
	}

	if o.name != "" {
		t.Name = c.taskPath(queueID, o.name)
	}

	_, err = c.service.Projects.Locations.Queues.Tasks.Create(c.queuePath(queueID), &cloudtasks.CreateTaskRequest{
		Task: t,
	}).Context(ctx).Do()
	return convertErrorToGRPC(err)
}

func convertErrorToGRPC(err error) error {
	if err == nil {
		return nil
	}
	switch e := err.(type) {
	case *googleapi.Error:
		return util.Error(util.StatusCodeToGRPCError(e.Code), e.Error())
	default:
		return util.Error(util.Internal, "%s", err.Error())
	}
}

func (c *client) isLocal() bool {
	// Task queue execution is not yet supported locally
	if c.environment == config.Internal || c.environment == config.Local {
		return true
	}
	return false
}
