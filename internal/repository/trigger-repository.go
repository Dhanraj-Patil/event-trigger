package repository

import (
	"context"

	"github.com/Dhanraj-Patil/event-trigger/internal/models"
	"github.com/Dhanraj-Patil/event-trigger/internal/scheduler"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateTrigger(trigger *models.Trigger) error {
	if !trigger.Active {
		trigger.Active = true
	}

	info, err := scheduler.ScheduleTask(trigger)
	if err != nil {
		return err
	}
	trigger.TriggerId = info
	return mgm.Coll(trigger).Create(trigger)
}

func GetAll(userId string) ([]models.Trigger, error) {
	var triggers []models.Trigger
	err := mgm.Coll(&models.Trigger{}).SimpleFind(&triggers, bson.M{"userId": userId})
	return triggers, err
}

func EditTrigger(id string, updateData bson.M) error {
	trigger := &models.Trigger{}
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	err = mgm.Coll(trigger).FindByID(objID, trigger)
	if err != nil {
		return err
	}

	update := bson.M{"$set": updateData}

	_, err = mgm.Coll(trigger).UpdateOne(context.TODO(), bson.M{"_id": objID}, update)
	if err != nil {
		return err
	}
	return nil
}

func DeleteTrigger(id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = mgm.Coll(&models.Trigger{}).DeleteOne(context.TODO(), bson.M{"_id": objID})
	if err != nil {
		return err
	}
	return nil
}
