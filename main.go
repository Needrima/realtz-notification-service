package main

import (
	"os"
	eventHandler "realtz-notification-service/internal/adapter/event-handler"
	handler "realtz-notification-service/internal/adapter/http-handler"
	emailRepo "realtz-notification-service/internal/adapter/repository/email"
	mongoRepo "realtz-notification-service/internal/adapter/repository/mongodb"
	redisRepo "realtz-notification-service/internal/adapter/repository/redis"
	smsRepo "realtz-notification-service/internal/adapter/repository/sms"
	"realtz-notification-service/internal/adapter/routes"
	configHelper "realtz-notification-service/internal/core/helpers/configuration-helper"
	logHelper "realtz-notification-service/internal/core/helpers/log-helper"
	redisHelper "realtz-notification-service/internal/core/helpers/redis-helper"
	validationHelper "realtz-notification-service/internal/core/helpers/validation-helper"
	services "realtz-notification-service/internal/core/service"
)

func main() {
	// iniitalize logger
	logHelper.InitializeLogger()

	// initialize struct validation for gin binding
	validationHelper.InitBindingValidation()

	// start api on database level (mongodb and redis)
	mongoRepo := mongoRepo.ConnectToMongoDB()
	redisRepo := redisRepo.ConnectToRedis()

	// connect to gmail smtp
	gmailSmtpClient := emailRepo.ConnectToGmail()

	// connect to mailgun
	mailgunClient := emailRepo.ConnectToMailGun()

	// connect to twilio
	twilioSmsClient := smsRepo.ConnectToTwilio()

	// start api on service level
	service := services.NewService(gmailSmtpClient, mailgunClient, twilioSmsClient, mongoRepo, redisRepo)

	// start api on http level
	handler := handler.NewHTTPHandler(service)
	router := routes.SetupRouter(handler)

	config := configHelper.ServiceConfiguration
	go func() {
		logHelper.LogEvent(logHelper.InfoLog, "starting server on port "+config.ServicePort)
		if err := router.Run(":" + config.ServicePort); err != nil {
			logHelper.LogEvent(logHelper.DangerLog, "could not start server "+err.Error())
			os.Exit(1)
		}
	}()

	go func() {
		redisRepo.SubsribeToEvent(redisHelper.USERCREATED, eventHandler.SendNotificationHandler)
	}()

	go func() {
		redisRepo.SubsribeToEvent(redisHelper.USERLOGGEDIN, eventHandler.SendNotificationHandler)
	}()

	go func() {
		redisRepo.SubsribeToEvent(redisHelper.SENDOTP, eventHandler.SendNotificationHandler)
	}()

	go func() {
		redisRepo.SubsribeToEvent(redisHelper.UPDATEPHONENUMBER, eventHandler.SendNotificationHandler)
	}()

	go func() {
		redisRepo.SubsribeToEvent(redisHelper.EMAILVERIFED, eventHandler.SendNotificationHandler)
	}()

	go func() {
		redisRepo.SubsribeToEvent(redisHelper.PHONENUMBERVERIFIED, eventHandler.SendNotificationHandler)
	}()

	go func() {
		redisRepo.SubsribeToEvent(redisHelper.PRODUCTADDED, eventHandler.SendNotificationHandler)
	}()

	go func() {
		redisRepo.SubsribeToEvent(redisHelper.PASSWORDRESET, eventHandler.SendNotificationHandler)
	}()

	go func() {
		redisRepo.SubsribeToEvent(redisHelper.USERRATED, eventHandler.SendNotificationHandler)
	}()

	go func() {
		redisRepo.SubsribeToEvent(redisHelper.ACCOUNTSWITCH, eventHandler.SendNotificationHandler)
	}()

	select {}
}
