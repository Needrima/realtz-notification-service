package helpers

import (
	"realtz-notification-service/internal/core/domain/dto"
	"realtz-notification-service/internal/core/domain/entity"
	"time"

	"github.com/google/uuid"
)

func CreateNotificationFromNotificationDto(SendNotificationDto dto.SendNotificationDto) entity.Notification {
	return entity.Notification{
		Reference:     uuid.New().String(),
		UserReference: SendNotificationDto.UserReference,
		Contact:       SendNotificationDto.Contact,
		Channel:       SendNotificationDto.Channel,
		Message:       SendNotificationDto.Message,
		Subject:       SendNotificationDto.Subject,
		Type:          SendNotificationDto.Type,
		CreatedOn: time.Now().Format(time.RFC3339),
	}
}
