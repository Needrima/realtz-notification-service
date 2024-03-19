package ports

import (
	"context"
	"realtz-notification-service/internal/core/domain/entity"
)

type MongoDBPort interface {
	CreateNotification(ctx context.Context, user entity.Notification) (interface{}, error)
	GetNotifications(ctx context.Context, currentUser entity.User, skip, limit int) (interface{}, int64, error)
}
