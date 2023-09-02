package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/hibiken/asynq"
)

// Task payload for any email related tasks.
type EmailTaskPayload struct {
	// ID for the email recipient.
	UserID int
}

const (
	QueueEmailWelcome = "email_welcome"
)

func EnqueueWelcomeEmail(client *asynq.Client, id string, userID int) {
	// Create a task with typename and payload.
	payload, err := json.Marshal(EmailTaskPayload{UserID: userID})
	if err != nil {
		log.Fatal(err)
	}
	task := asynq.NewTask("email:welcome", payload)

	opts := []asynq.Option{
		asynq.Queue(QueueEmailWelcome),
		//asynq.MaxRetry(3),
		asynq.TaskID(id),
		asynq.Deadline(),
		asynq.Unique(time.Minute * 5),
		//asynq.Retention(time.Minute * 20),
		//asynq.ProcessIn(24*time.Hour),
	}
	info, err := client.Enqueue(task, opts...)
	if err != nil {
		fmt.Printf("err %s\n", err)
	}
	fmt.Printf("info: %+v\n", info)
}

// client.go
func main() {
	client := asynq.NewClient(asynq.RedisClientOpt{Addr: "localhost:6379"})

	EnqueueWelcomeEmail(client, "10", 10)
}
