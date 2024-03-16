package entity

type SendNotification struct {
	Contact string `json:"contact" binding:"required,valid_contact"`          // phone number or email
	Channel string `json:"channel" binding:"required,eq=email|eq=sms|eq=all"` // can only one of sms|email|all
	Message string `json:"message" binding:"required"`
	Subject string `json:"subject" binding:"required"`
}
