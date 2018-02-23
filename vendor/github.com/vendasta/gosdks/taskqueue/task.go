package taskqueue

import (
	"encoding/base64"
	"errors"
	"google.golang.org/api/cloudtasks/v2beta2"
)

// Task represents a cloudtasks Task with a simpler api.
type Task struct {
	// ackID is set by the taskqueue service and represents an ID that must be provided to the taskqueue API for the task to be successfully acknowledged
	// In the case of cloud tasks, this role is filled by cloudtasks.Task.ScheduleTime
	ackID   string
	taskID  string
	Payload string
	tag     string
	attempt int64
}

func taskFromAPI(task *cloudtasks.Task) (Task, error) {
	t := Task{}
	t.ackID = task.ScheduleTime
	t.taskID = task.Name
	if task.Status != nil {
		t.attempt = task.Status.AttemptDispatchCount
	}
	if task.PullMessage != nil {
		p, err := base64.StdEncoding.DecodeString(task.PullMessage.Payload)
		if err != nil {
			return Task{}, errors.New("error decoding task payload")
		}
		t.Payload = string(p[:])
		t.tag = task.PullMessage.Tag
	}
	return t, nil
}
