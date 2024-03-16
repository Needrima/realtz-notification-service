package main

import (
	"os"
	handler "realtz-notification-service/internal/adapter/http-handler"
	emailRepo "realtz-notification-service/internal/adapter/repository/email"
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
	redisRepo := redisRepo.ConnectToRedis()

	// connect to gmail smtp
	gmailSmtpClient := emailRepo.ConnectToGmail()

	// connect to mailgun
	mailgunClient := emailRepo.ConnectToMailGun()

	// connect to twilio
	twilioSmsClient := smsRepo.ConnectToTwilio()

	// start api on service level
	service := services.NewService(gmailSmtpClient, mailgunClient, twilioSmsClient, redisRepo)

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
		redisRepo.SubsribeToEvent(redisHelper.USERCREATED, redisHelper.SendNotificationHandler)
	}()

	go func() {
		redisRepo.SubsribeToEvent(redisHelper.USERLOGGEDIN, redisHelper.SendNotificationHandler)
	}()

	go func() {
		redisRepo.SubsribeToEvent(redisHelper.SENDOTP, redisHelper.SendNotificationHandler)
	}()

	go func() {
		redisRepo.SubsribeToEvent(redisHelper.UPDATEPHONENUMBER, redisHelper.SendNotificationHandler)
	}()

	go func() {
		redisRepo.SubsribeToEvent(redisHelper.EMAILVERIFED, redisHelper.SendNotificationHandler)
	}()

	go func() {
		redisRepo.SubsribeToEvent(redisHelper.PHONENUMBERVERIFIED, redisHelper.SendNotificationHandler)
	}()

	go func() {
		redisRepo.SubsribeToEvent(redisHelper.PRODUCTADDED, redisHelper.SendNotificationHandler)
	}()

	select {}
}
