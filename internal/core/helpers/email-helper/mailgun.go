// check mailgun documentation at https://github.com/mailgun/mailgun-go

package helpers

import (
	"context"
	"fmt"
	errorHelper "realtz-notification-service/internal/core/helpers/error-helper"
	logHelper "realtz-notification-service/internal/core/helpers/log-helper"
	"time"

	"github.com/mailgun/mailgun-go/v4"
)

type MailgunEmailClient struct {
	mailgunClient *mailgun.MailgunImpl
}

func NewmailgunEmailClient(mailgunClient *mailgun.MailgunImpl) MailgunEmailClient {
	return MailgunEmailClient{
		mailgunClient: mailgunClient,
	}
}

func (m *MailgunEmailClient) SendMail(from, receiver, subject, message string) error {

	mail := m.mailgunClient.NewMessage(from, subject, message, receiver)

	_, _, err := m.mailgunClient.Send(context.Background(), mail)

	if err != nil {
		logHelper.LogEvent(logHelper.ErrorLog, fmt.Sprintf("sending email to %s unsuccessful, error: %v", receiver, err))
		return errorHelper.NewServiceError("could not send email notification to "+receiver, 500)
	}

	logHelper.LogEvent(logHelper.SuccessLog, fmt.Sprintf("mail sent to %s successfully on %s", receiver, time.Now().Format(time.RFC3339)))

	return nil
}
