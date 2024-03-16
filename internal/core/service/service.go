package services

import (
	"net/smtp"
	"realtz-notification-service/internal/core/domain/dto"
	emailHelper "realtz-notification-service/internal/core/helpers/email-helper"
	smsHelper "realtz-notification-service/internal/core/helpers/sms-helper"
	"realtz-notification-service/internal/ports"

	"github.com/mailgun/mailgun-go/v4"
	"github.com/twilio/twilio-go"
)

type Service struct {
	gmailClient     smtp.Auth            // gmail email client
	mailgunClient   *mailgun.MailgunImpl // mailgun email client
	twilioSmsClient *twilio.RestClient
	redisPort       ports.RedisPort
}

var NotificationService Service

func NewService(gmailClient smtp.Auth, mailgunClient *mailgun.MailgunImpl, twilioSmsClient *twilio.RestClient, redisPort ports.RedisPort) Service {
	NotificationService = Service{
		gmailClient:     gmailClient,
		mailgunClient:   mailgunClient,
		twilioSmsClient: twilioSmsClient,
		redisPort:       redisPort,
	}

	return NotificationService
}

func (s Service) SendNotification(SendNotificationDto dto.SendNotificationDto) (interface{}, error) {
	switch SendNotificationDto.Channel {
	case "email":
		emailClient := emailHelper.NewGmailEmailClient(s.gmailClient)
		if err := emailClient.SendMail(SendNotificationDto.Contact, SendNotificationDto.Subject, SendNotificationDto.Message); err != nil {
			return nil, err
		}
	case "sms":
		// send sms with twilio API
		smsClient := smsHelper.NewSmsClient(s.twilioSmsClient)
		if err := smsClient.SendSMS(SendNotificationDto.Contact, SendNotificationDto.Message); err != nil {
			return nil, err
		}
	}

	sendNotificationResponse := struct {
		Message string `json:"message"`
		Success bool   `json:"success"`
	}{
		Message: "Notification sent",
		Success: true,
	}

	return sendNotificationResponse, nil
}
