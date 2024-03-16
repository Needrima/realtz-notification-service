package helpers

import (
	"regexp"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func InitBindingValidation() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("valid_phone_number", ValidNigerianPhoneNumber)
		v.RegisterValidation("valid_contact", ValidContact)
	}
}

// EqualFieldValidation is a custom validation function
func ValidNigerianPhoneNumber(fl validator.FieldLevel) bool {
	fieldValue := fl.Field().String()

	phonePattern := `^0\d{10}$`

	return regexp.MustCompile(phonePattern).MatchString(fieldValue)
}

// EqualFieldValidation is a custom validation function
func ValidContact(fl validator.FieldLevel) bool {
	fieldValue := fl.Field().String()

	phonePattern := `^0\d{10}$`
	emailPattern := `(?:[a-z0-9!#$%&'*+\/=?^_{|}~-]+(?:\.[a-z0-9!#$%&'*+\/=?^_{|}~-]+)*|"(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21\x23-\x5b\x5d-\x7f]|
		\\[\x01-\x09\x0b\x0c\x0e-\x7f])*")@(?:(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?|
		\[(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?|[a-z0-9-]*[a-z0-9]:
		(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21-\x5a\x53-\x7f]|\\[\x01-\x09\x0b\x0c\x0e-\x7f])+)\])`

	isPhoneMatch := regexp.MustCompile(phonePattern).MatchString(fieldValue)
	isEmailMatch := regexp.MustCompile(emailPattern).MatchString(fieldValue)

	return isPhoneMatch || isEmailMatch
}
