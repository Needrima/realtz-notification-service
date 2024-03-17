package helpers

import (
	"realtz-notification-service/internal/core/domain/dto"
	"realtz-notification-service/internal/core/domain/eto"
	logHelper "realtz-notification-service/internal/core/helpers/log-helper"
	services "realtz-notification-service/internal/core/service"
)

func extractDataFromEvent(event eto.Event) map[string]interface{} {
	data, ok := event.Data.(map[string]interface{})
	if !ok {
		logHelper.LogEvent(logHelper.ErrorLog, "could not assert data in event to a map")
		return nil
	}

	return data
}

func SendNotificationHandler(evenJson string) {
	event, _ := eto.EventJsonToEvent(evenJson)
	data := extractDataFromEvent(event)

	sendNotificationDto := dto.SendNotificationDto{
		UserReference: data["user_reference"].(string),
		Message:       data["message"].(string),
		Subject:       data["subject"].(string),
		Channel:       data["channel"].(string),
		Contact:       data["contact"].(string),
		Type:          data["type"].(string),
	}

	services.NotificationService.SendNotification(sendNotificationDto)
}
