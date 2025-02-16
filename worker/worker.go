package main

import (
	"log"

	"github.com/Dhanraj-Patil/event-trigger/internal/scheduler"
	"github.com/hibiken/asynq"
)

const redisAddr = "redis:6379"

func main() {
	srv := asynq.NewServer(
		asynq.RedisClientOpt{Addr: redisAddr},
		asynq.Config{
			Concurrency: 10,
		},
	)

	mux := asynq.NewServeMux()
	mux.HandleFunc(scheduler.TypeSendSMS, scheduler.HandleSendSMSTask)

	if err := srv.Run(mux); err != nil {
		log.Fatalf("could not run server: %v", err)
	}
}
