package helpers

import (
	"github.com/pquerna/otp/totp"
	configHelper "realtz-notification-service/internal/core/helpers/configuration-helper"
	errorHelper "realtz-notification-service/internal/core/helpers/error-helper"
	logHelper "realtz-notification-service/internal/core/helpers/log-helper"
	"time"
)

func GenerateOTP(acctName string) (string, string, error) {
	// Generate a new TOTP key
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      configHelper.ServiceConfiguration.ServiceName,
		AccountName: acctName,
		Period:      uint(time.Minute) * 2,
	})

	if err != nil {
		logHelper.LogEvent(logHelper.ErrorLog, "Error generating TOTP key: "+err.Error())
		return "", "", errorHelper.NewServiceError("something went wrong", 500)
	}

	// Generate and display the current TOTP
	otp, err := totp.GenerateCode(key.Secret(), time.Now())
	if err != nil {
		logHelper.LogEvent(logHelper.ErrorLog, "Error generating otp: "+err.Error())
		return "", "", errorHelper.NewServiceError("something went wrong", 500)
	}

	return otp, key.Secret(), nil
}

func ValidateOtp(otp, secret string) bool {
	return totp.Validate(otp, secret)
}
