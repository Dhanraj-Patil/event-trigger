package scheduler

import (
	"fmt"
	"log"

	"github.com/hibiken/asynq"
)

var taskScheduler *asynq.Scheduler

func StartScheduler() {
	redisConn := asynq.RedisClientOpt{Addr: "localhost:6379"}
	taskScheduler = asynq.NewScheduler(&redisConn, nil)

	fmt.Println("Scheduler Running...")
	if err := taskScheduler.Run(); err != nil {
		log.Fatal(err)
	}
}
