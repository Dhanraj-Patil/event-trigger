package services

import (
	"github.com/Dhanraj-Patil/event-trigger/internal/models"
	"github.com/Dhanraj-Patil/event-trigger/internal/repository"
	"github.com/Dhanraj-Patil/event-trigger/internal/scheduler"
	"github.com/hibiken/asynq"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateTrigger(trigger *models.Trigger) error {
	err := repository.CreateTrigger(trigger)
	if err != nil {
		return err
	}
	return nil
}

func TestTrigger(trigger *models.Trigger) (*asynq.TaskInfo, error) {
	info, err := scheduler.TestRun(trigger)
	if err != nil {
		return nil, err
	}
	return info, nil
}

func GetAllByUser(userId string) ([]models.Trigger, error) {
	data, err := repository.GetAll(userId)
	return data, err
}

func EditTrigger(id string, updateData bson.M) error {
	err := repository.EditTrigger(id, updateData)
	if err != nil {
		return err
	}
	return nil
}

func DeleteTrigger(id string) error {
	if err := repository.DeleteTrigger(id); err != nil {
		return err
	}
	return nil
}
