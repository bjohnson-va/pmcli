package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/vendasta/gosdks/config"
	"github.com/vendasta/gosdks/taskqueue"
	"github.com/vendasta/gosdks/util"
)

// gcloud alpha tasks queues create-pull-queue dwalker-test
// gcloud alpha tasks queues describe dwalker-test
// gcloud alpha tasks queues delete dwalker-test
// gcloud alpha tasks queues resume bbass-test

type FakePayload struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func main() {
	ctx := context.Background()
	c, err := taskqueue.NewClient(ctx, config.Test)
	if err != nil {
		log.Fatalf("error making client: %s", err.Error())
	}
	log.Printf("client: %v", c)

	payloads := []*FakePayload{
		{
			FirstName: "braden",
			LastName:  "bassingthwaite",
		},
		{
			FirstName: "leo",
			LastName:  "melnyk",
		},
		{
			FirstName: "jason",
			LastName:  "prokop",
		},
		{
			FirstName: "riley",
			LastName:  "wiebe",
		},
	}

	for _, p := range payloads {
		err = c.ScheduleTask(ctx, "jredl", p, taskqueue.WithTag(p.FirstName))
		if err != nil {
			if util.IsError(util.AlreadyExists, err) {
				log.Printf("task already exists: %s", err.Error())
			} else {
				log.Fatalf("error scheduling task: %s", err.Error())
			}
		}
	}

	h := func(ctx context.Context, msg taskqueue.Task) error {
		log.Printf("raw task payload: %s", msg.Payload)
		p := &FakePayload{}
		err := json.Unmarshal([]byte(msg.Payload), p)
		if err != nil {
			return err
		}
		log.Printf("deserialized first name: %s", p.FirstName)
		log.Printf("deserialized last name: %s", p.LastName)
		return nil
	}
	w := taskqueue.NewWorker("jredl", h, c, taskqueue.WithTagFilter("braden"))
	w.Work(ctx)
}
