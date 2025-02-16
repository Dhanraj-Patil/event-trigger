package scheduler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	integration "github.com/Dhanraj-Patil/event-trigger/internal/integrations"
	"github.com/Dhanraj-Patil/event-trigger/internal/utils"
	"github.com/hibiken/asynq"
)

const (
	TypeSendSMS = "sms:send"
)

type SendSMSPayload struct {
	PhoneNo string
	Message string
}

func NewSendSMSTask(phoneNo string, message string) (*asynq.Task, error) {
	payload, err := json.Marshal((SendSMSPayload{PhoneNo: phoneNo, Message: message}))
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TypeSendSMS, payload), nil
}

func HandleSendSMSTask(ctx context.Context, t *asynq.Task) error {
	var p SendSMSPayload
	taskId, ok := asynq.GetTaskID(ctx)

	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}

	log.Printf("Sending SMS to Number: +91%s, Message: %s", p.PhoneNo, p.Message)
	res, err := integration.TwilioSMSAPI(p.PhoneNo, p.Message)
	if err != nil && ok {
		utils.LogById(taskId, err.Error())
	}
	utils.LogById(taskId, res)
	return nil
}
