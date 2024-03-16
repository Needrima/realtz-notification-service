package repository

import (
	configHelper "realtz-notification-service/internal/core/helpers/configuration-helper"

	"github.com/twilio/twilio-go"
)

func ConnectToTwilio() *twilio.RestClient {
	accountSid := configHelper.ServiceConfiguration.TwilioAccountSID
	authToken := configHelper.ServiceConfiguration.TwilioAuthToken
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})

	return client
}
