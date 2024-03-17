package entity

type Notification struct {
	Reference string `json:"reference" bson:"reference"`
	Contact   string `json:"contact" binding:"required,valid_contact"`          // phone number or email
	Channel   string `json:"channel" binding:"required,eq=email|eq=sms|eq=all"` // can only one of sms|email|all
	Message   string `json:"message" binding:"required"`
	Subject   string `json:"subject" binding:"required"`
	Type      string `json:"type" bson:"type" binding:"required,eq=in_app|eq=sending"` // in_app for notifications that show on frontend too and sending for those that are sent as email or sms
}
