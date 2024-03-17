package services

import (
	"context"
	"net/smtp"
	"realtz-notification-service/internal/core/domain/dto"
	"realtz-notification-service/internal/core/domain/entity"
	emailHelper "realtz-notification-service/internal/core/helpers/email-helper"
	mapper "realtz-notification-service/internal/core/helpers/mapper"
	mongodbHelper "realtz-notification-service/internal/core/helpers/mongodb-helper"
	smsHelper "realtz-notification-service/internal/core/helpers/sms-helper"
	"realtz-notification-service/internal/ports"

	"github.com/mailgun/mailgun-go/v4"
	"github.com/twilio/twilio-go"
)

type Service struct {
	gmailClient     smtp.Auth            // gmail email client
	mailgunClient   *mailgun.MailgunImpl // mailgun email client
	twilioSmsClient *twilio.RestClient
	dbPort          ports.MongoDBPort
	redisPort       ports.RedisPort
}

var NotificationService Service

func NewService(gmailClient smtp.Auth, mailgunClient *mailgun.MailgunImpl, twilioSmsClient *twilio.RestClient, dbPort ports.MongoDBPort, redisPort ports.RedisPort) Service {
	NotificationService = Service{
		gmailClient:     gmailClient,
		mailgunClient:   mailgunClient,
		twilioSmsClient: twilioSmsClient,
		dbPort:          dbPort,
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

	newNotification := mapper.CreateNotificationFromNotificationDto(SendNotificationDto)

	_, err := s.dbPort.CreateNotification(context.Background(), newNotification)
	if err != nil {
		return nil, err
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

func (s Service) GetNotifications(ctx context.Context, currentUser entity.User, amount, pageNo string) (interface{}, error) {
	skip, limit, pageNoInt := mongodbHelper.GetSkipAndLimit(amount, pageNo)

	notifications, err := s.dbPort.GetNotifications(ctx, currentUser, skip, limit)
	if err != nil {
		return nil, err
	}

	getNotificationsResponse := struct {
		Notifications interface{} `json:"products"`
		PreviousPage  int         `json:"previous_page"`
		NextPage      int         `json:"next_page"`
		Success       bool        `json:"success"`
		Message       string      `json:"message"`
	}{
		Notifications: notifications,
		PreviousPage:  pageNoInt - 1,
		NextPage:      pageNoInt + 1,
		Success:       true,
		Message:       "Succesfully retrieved products",
	}

	return getNotificationsResponse, nil
}
