package ports

import "realtz-notification-service/internal/core/domain/dto"

// "context"
// "realtz-notification-service/internal/core/domain/dto"
// "realtz-notification-service/internal/core/domain/entity"

type HTTPPort interface {
	SendNotification(SendNotificationDto dto.SendNotificationDto) (interface{}, error)
}
