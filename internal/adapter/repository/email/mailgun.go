package repository

import (
	configHelper "realtz-notification-service/internal/core/helpers/configuration-helper"

	"github.com/mailgun/mailgun-go/v4"
)

func ConnectToMailGun() *mailgun.MailgunImpl {
	mg := mailgun.NewMailgun(configHelper.ServiceConfiguration.MailgunDomain,
		configHelper.ServiceConfiguration.MailgunPrivateKey)

	return mg
}
