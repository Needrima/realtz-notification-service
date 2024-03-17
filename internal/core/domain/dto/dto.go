package dto

type SendNotificationDto struct {
	UserReference string `json:"user_reference" bson:"user_reference"`                             // reference of owner
	Contact       string `json:"contact" bson:"contact" binding:"required"`                        // phone number or email
	Channel       string `json:"channel" bson:"channel" binding:"required,eq=email|eq=sms|eq=all"` // can only one of sms|email|all
	Message       string `json:"message" bson:"message" binding:"required"`
	Subject       string `json:"subject" bson:"subject" binding:"required"`
	Type          string `json:"type" bson:"type" binding:"required,eq=in_app|eq=sending"` // in_app for notifications that show on frontend too and sending for those that are sent as email or sms
}
