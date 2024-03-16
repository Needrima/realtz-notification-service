package helpers

import (
	"fmt"
	configHelper "realtz-notification-service/internal/core/helpers/configuration-helper"
	errorHelper "realtz-notification-service/internal/core/helpers/error-helper"
	logHelper "realtz-notification-service/internal/core/helpers/log-helper"
	"time"

	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
)

type SmsClient struct {
	smsClient *twilio.RestClient
}

func NewSmsClient(smsClient *twilio.RestClient) SmsClient {
	return SmsClient{
		smsClient: smsClient,
	}
}

func (e *SmsClient) SendSMS(receiver, message string) error {
	params := &twilioApi.CreateMessageParams{
		From: &configHelper.ServiceConfiguration.TwilioAuthPhoneNumber,
		To:   &receiver,
		Body: &message,
	}

	_, err := e.smsClient.Api.CreateMessage(params)
	if err != nil {
		logHelper.LogEvent(logHelper.ErrorLog, fmt.Sprintf("sending email to %s unsuccessful, error: %v", receiver, err))
		return errorHelper.NewServiceError("could not send sms notification to "+receiver, 500)
	}

	logHelper.LogEvent(logHelper.SuccessLog, fmt.Sprintf("sms sent to %s successfully on %s", receiver, time.Now().Format(time.RFC3339)))

	return nil
}
