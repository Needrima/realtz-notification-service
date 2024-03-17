package repository

import (
	"context"
	"fmt"
	"realtz-notification-service/internal/core/domain/entity"
	errorHelper "realtz-notification-service/internal/core/helpers/error-helper"
	logHelper "realtz-notification-service/internal/core/helpers/log-helper"

	"go.mongodb.org/mongo-driver/mongo"
)

type mongoRepo struct {
	collection *mongo.Collection
}

func NewMongoRepo(collection *mongo.Collection) mongoRepo {
	return mongoRepo{
		collection: collection,
	}
}

func (m mongoRepo) CreateNotification(ctx context.Context, notification entity.Notification) (interface{}, error) {

	_, err := m.collection.InsertOne(ctx, notification)
	if err != nil {
		logHelper.LogEvent(logHelper.ErrorLog, "could not store new user in db: "+err.Error())
		return nil, errorHelper.NewServiceError("something went wrong", 500)
	}

	logHelper.LogEvent(logHelper.SuccessLog, fmt.Sprintf("successfully created notification. Reference: %s", notification.Reference))

	return "notification created successfully", nil
}
