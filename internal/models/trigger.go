package models

import (
	"time"

	"github.com/kamva/mgm/v3"
)

type Trigger struct {
	mgm.DefaultModel `bson:",inline"`
	UserId           string    `json:"userId" bson:"userId" binding:"required"`
	Schedule         time.Time `json:"schedule" bson:"schedule" binding:"required"`
	Active           bool      `json:"active" bson:"active"`
	Message          string    `json:"message" bson:"message" binding:"required"`
	PhoneNo          string    `json:"phoneNo" bson:"phoneNo" binding:"required,len=10,numeric"`
	TriggerId        string    `json:"triggerId" bson:"triggerId"`
}

var AllowedFields = map[string]bool{
	"repeat":   true,
	"interval": true,
	"schedule": true,
	"message":  true,
	"phoneNo":  true,
}
