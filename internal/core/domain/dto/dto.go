package dto

type SendNotificationDto struct {
	Contact string `json:"contact" binding:"required"`                        // phone number or email
	Channel string `json:"channel" binding:"required,eq=email|eq=sms|eq=all"` // can only one of sms|email|all
	Message string `json:"message" binding:"required"`
	Subject string `json:"subject" binding:"required"`
}
