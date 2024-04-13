package helpers

import (
	"fmt"
	"net/smtp"
	configHelper "realtz-notification-service/internal/core/helpers/configuration-helper"
	errorHelper "realtz-notification-service/internal/core/helpers/error-helper"
	logHelper "realtz-notification-service/internal/core/helpers/log-helper"
	"time"
)

type GmailEmailClient struct {
	gmailClient smtp.Auth
}

func NewGmailEmailClient(gmailClient smtp.Auth) GmailEmailClient {
	return GmailEmailClient{
		gmailClient: gmailClient,
	}
}

func (e *GmailEmailClient) SendMail(receiver, subject, message string) error {
	from := configHelper.ServiceConfiguration.GoogleAuthUser
	to := []string{receiver}

	// format message in RFC 2822 standard
	msg := []byte("To: " + receiver + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" +
		message + "\r\n")

	// msg = fmt.Sprintf("From: %s\r\n", from) +
	// fmt.Sprintf("To: %s\r\n", to) +
	// fmt.Sprintf("Subject: %s\r\n", subject) +
	// "MIME-Version: 1.0\r\n" +
	// "Content-Type: text/html; charset=\"utf-8\"\r\n\r\n" +
	// body

	err := smtp.SendMail(configHelper.ServiceConfiguration.GoogleSmtpHost+":"+configHelper.ServiceConfiguration.GoogleSmtpPort, e.gmailClient, from, to, msg)
	if err != nil {
		logHelper.LogEvent(logHelper.ErrorLog, fmt.Sprintf("sending email to %s unsuccessful, error: %v", receiver, err))
		return errorHelper.NewServiceError("could not send email notification to "+receiver, 500)
	}

	logHelper.LogEvent(logHelper.SuccessLog, fmt.Sprintf("mail sent to %s successfully on %s", receiver, time.Now().Format(time.RFC3339)))

	return nil
}
