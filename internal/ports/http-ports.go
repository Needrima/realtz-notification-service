package ports

import (
	"context"
	"realtz-notification-service/internal/core/domain/dto"
	"realtz-notification-service/internal/core/domain/entity"
)

type HTTPPort interface {
	SendNotification(SendNotificationDto dto.SendNotificationDto) (interface{}, error)
	GetNotifications(ctx context.Context, currentUser entity.User, amount, pageNo string) (interface{}, error)
}
