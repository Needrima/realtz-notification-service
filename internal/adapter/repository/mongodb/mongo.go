package repository

import (
	"context"
	"fmt"
	"realtz-notification-service/internal/core/domain/entity"
	errorHelper "realtz-notification-service/internal/core/helpers/error-helper"
	logHelper "realtz-notification-service/internal/core/helpers/log-helper"

	"go.mongodb.org/mongo-driver/bson"
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

func (m mongoRepo) GetNotifications(ctx context.Context, currentUser entity.User, skip, limit int) (interface{}, int64, error) {
	matchStage := bson.M{"$match": bson.M{"user_reference": currentUser.Reference, "type": "in_app"}}
	sortStage := bson.M{"$sort": bson.M{"created_on": -1}}
	skipStage := bson.M{"$skip": skip}
	limitStage := bson.M{"$limit": limit}

	notifications := []entity.Notification{}
	cursor, err := m.collection.Aggregate(ctx, []bson.M{matchStage, sortStage, skipStage, limitStage})
	if err != nil {
		logHelper.LogEvent(logHelper.ErrorLog, fmt.Sprintf("could not retrieve notifications from database, error: %s", err.Error()))
		return nil, 0, errorHelper.NewServiceError("something went wrong", 500)
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &notifications); err != nil {
		logHelper.LogEvent(logHelper.ErrorLog, fmt.Sprintf("could not decode notifications from database, error: %s", err.Error()))
		return nil, 0, errorHelper.NewServiceError("something went wrong", 500)
	}

	documentsCount, _ := m.collection.CountDocuments(ctx, bson.M{"user_reference": currentUser.Reference})

	return notifications, documentsCount, nil
}
