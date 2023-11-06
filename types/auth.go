package types

type User struct {
	ID 					string           	`bson:"_id,omitempty" json:"id,omitempty"`
	FirstName 			string 				`bson:"firstName" json:"firstName"`
	LastName 			string 				`bson:"lastName" json:"lastName"`
	Email 				string 				`bson:"email" json:"email"`
	PhoneNumber 		string 				`bson:"phone_number" json:"phoneNumber"`
	EncryptedPassword 	string 				`bson:"EncryptedPassword" json:"-"` 
	EmailVerified 		bool 				`bson:"email_verified" json:"emailVerified"`
	PhoneVerified		bool          		`bson:"phone_verified" json:"phoneVerified"`
}

func (u *User) ExcludePass() *User {
	u.EncryptedPassword = ""
	return u
}

//types that are only for this service
type UpdateUserParams struct {
	FilterBy 			string  `json:"update_by"`
	FilterItem 			string  `json:"filter_item"`
	FirstName 	      	string 	`json:"firstName"`
	LastName 		  	string 	`json:"lastName"`
	Email 				string 	`json:"email"`
	PhoneNumber 		string 	`json:"phoneNumber"`
	Password 			string 	`json:"password"`
}

type CreateUserParams struct {
	FirstName 	string 	`json:"firstName"`
	LastName 	string 	`json:"lastName"`
	Email 		string 	`json:"email"`
	Password 	string 	`json:"password"`
}

type ResetPasswordParams struct {
	Email 	string `json:"email"`
	NewPass string `json:"newPassword"`
	Code 	string `json:"code"`
}

