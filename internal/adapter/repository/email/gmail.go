package repository

import (
	"net/smtp"
	configHelper "realtz-notification-service/internal/core/helpers/configuration-helper"
)

func ConnectToGmail() smtp.Auth {
	// Set up authentication
	auth := smtp.PlainAuth("",
		configHelper.ServiceConfiguration.GoogleAuthUser,
		configHelper.ServiceConfiguration.GoogleAppPassword,
		configHelper.ServiceConfiguration.GoogleSmtpHost)
	return auth
}
