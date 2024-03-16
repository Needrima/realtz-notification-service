package helpers

import (
	"realtz-notification-service/internal/core/domain/dto"
	"realtz-notification-service/internal/core/domain/entity"
)

func CreateNotificationFromNotificationDto(SendNotificationDto dto.SendNotificationDto) entity.SendNotification {
	return entity.SendNotification(SendNotificationDto)
}
