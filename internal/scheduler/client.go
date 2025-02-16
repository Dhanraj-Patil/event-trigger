package scheduler

import (
	"os"
	"time"

	"github.com/Dhanraj-Patil/event-trigger/internal/models"
	"github.com/Dhanraj-Patil/event-trigger/internal/utils"
	"github.com/hibiken/asynq"
	"github.com/joho/godotenv"
)

var client *asynq.Client

type Trigger struct {
	UserId   string    `json:"userId"`
	Repeat   bool      `json:"repeat"`
	Interval string    `json:"interval"`
	Schedule time.Time `json:"schedule"`
	Message  string    `json:"message"`
	PhoneNo  string    `json:"phoneNo"`
}

func InitAsynqClient() {
	godotenv.Load()
	client = asynq.NewClient(asynq.RedisClientOpt{Addr: os.Getenv("REDIS_ADDR")})
	// defer client.Close()
}

func TestRun(data *models.Trigger) (*asynq.TaskInfo, error) {
	task, err := NewSendSMSTask(data.PhoneNo, data.Message)
	if err != nil {
		return nil, err
	}
	info, err := client.Enqueue(task, asynq.Retention(2*time.Hour))
	if err != nil {
		return nil, err
	}
	var log models.EventLog
	log.Request = *data
	log.Trigger = *info
	utils.LogEvent(log)
	return info, nil
}

func ScheduleTask(data *models.Trigger) (*asynq.TaskInfo, error) {
	task, err := NewSendSMSTask(data.PhoneNo, data.Message)
	if err != nil {
		return nil, err
	}
	info, err := client.Enqueue(task, asynq.ProcessAt(data.Schedule), asynq.Retention(2*time.Hour))
	if err != nil {
		return nil, err
	}
	var log models.EventLog
	log.Request = *data
	log.Trigger = *info
	utils.LogEvent(log)
	return info, nil
}

// func SchedulePeriodicTask(data *models.Trigger) (string, error) {
// 	task, err := NewSendSMSTask(data.PhoneNo, data.Message)
// 	if err != nil {
// 		return "", err
// 	}
// 	entryId, err := taskScheduler.Register("@every "+data.Interval, task)
// 	if err != nil {
// 		return "", err
// 	}
// 	return entryId, nil
// }
