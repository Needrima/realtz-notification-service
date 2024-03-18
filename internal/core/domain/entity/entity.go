package entity

type Notification struct {
	Reference     string `json:"reference" bson:"reference"`
	UserReference string `json:"user_reference" bson:"user_reference"`              // reference of owner
	Contact       string `json:"contact" binding:"required,valid_contact"`          // phone number or email
	Channel       string `json:"channel" binding:"required,eq=email|eq=sms|eq=all"` // can only one of sms|email|all
	Message       string `json:"message" binding:"required"`
	Subject       string `json:"subject" binding:"required"`
	Type          string `json:"type" bson:"type" binding:"required,eq=in_app|eq=sending"` // in_app for notifications that show on frontend too and sending for those that are sent as email or sms
	CreatedOn string `json:"created_on" bson:"created_on"`
}

type User struct {
	Reference             string `json:"reference" bson:"reference"`
	UserType              string `json:"user_type" bson:"user_type"`
	Firstname             string `json:"firstname" bson:"firstname"`
	Lastname              string `json:"lastname" bson:"lastname"`
	Fullname              string `json:"fullname" bson:"fullname"`
	Email                 string `json:"email" bson:"email"`
	PhoneNumber           string `json:"phone_number" bson:"phone_number"`
	BVN                   string `json:"-" bson:"bvn"`
	Password              string `json:"-" bson:"password"`
	IsEmailVerified       bool   `json:"is_email_verified" bson:"is_email_verified"`
	IsPhoneNumberVerified bool   `json:"is_phone_number_verified" bson:"is_phone_number_verified"`
	IsBvnVerified         bool   `json:"is_bvn_verified" bson:"is_bvn_verified"`
	IsVerified            bool   `json:"is_verified" bson:"is_verified"`
	CreatedOn             string `json:"created_on" bson:"created_on"`
	LastUpdatedOn         string `json:"last_updated_on" bson:"last_updated_on"`
}
