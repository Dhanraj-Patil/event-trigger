package models

import (
	"github.com/hibiken/asynq"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
)

type EventLog struct {
	mgm.DefaultModel `bson:",inline"`
	Trigger          asynq.TaskInfo `json:"event" bson:"event"`
	Response         bson.M         `json:"response" bson:"response"`
	Request          Trigger        `json:"request" bson:"request"`
}
