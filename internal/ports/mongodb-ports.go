package ports

import (
	"context"
	"realtz-notification-service/internal/core/domain/entity"
)

type MongoDBPort interface {
	CreateNotification(ctx context.Context, user entity.Notification) (interface{}, error)
}
