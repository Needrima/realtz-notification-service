package helpers

import (
	"context"
	"realtz-notification-service/internal/core/domain/eto"
	errorHelper "realtz-notification-service/internal/core/helpers/error-helper"
	logHelper "realtz-notification-service/internal/core/helpers/log-helper"

	"github.com/redis/go-redis/v9"
)

type EventPublisher struct {
	redisClient *redis.Client
}

func NewPublisher(redisClient *redis.Client) EventPublisher {
	return EventPublisher{
		redisClient: redisClient,
	}
}

func (e *EventPublisher) PublishEvent(ctx context.Context, channelName string, data interface{}) error {
	event := eto.NewEvent(data)
	eventJson := event.ToJSON()
	if err := e.redisClient.Publish(ctx, channelName, eventJson).Err(); err != nil {
		logHelper.LogEvent(logHelper.ErrorLog, "could not publish event: "+err.Error())
		return errorHelper.NewServiceError("something went wrong", 500)
	}

	return nil
}

type EventSubscriber struct {
	redisClient *redis.Client
}

func NewSubscriber(redisClient *redis.Client) EventSubscriber {
	return EventSubscriber{
		redisClient: redisClient,
	}
}

func (e *EventSubscriber) SubsribeToEvent(channelName string, handler func(eventJson string)) {
	pubSub := e.redisClient.PSubscribe(context.Background(), channelName)
	defer pubSub.Close()

	ch := pubSub.Channel()
	for {
		select {
		case msg := <-ch:
			logHelper.LogEvent(logHelper.InfoLog, "received data from channle: "+channelName)
			handler(msg.Payload) // Pass the appropriate UserRepository instance here
		}
	}
}
