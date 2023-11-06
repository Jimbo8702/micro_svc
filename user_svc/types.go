package main

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2/bson"
)

type DBUser struct {
	ID 					primitive.ObjectID  `bson:"_id,omitempty"`
	FirstName 			string 				`bson:"firstName"`
	LastName 			string 				`bson:"lastName"`
	Email 				string 				`bson:"email"`
	PhoneNumber 		string 				`bson:"phone_number"`
	EncryptedPassword 	string 				`bson:"EncryptedPassword"` 
	EmailVerified 		bool 				`bson:"email_verified"`
	PhoneVerified		bool          		`bson:"phone_verified"`

	//add createdat updated at
}

type UserDBCreateParams struct {
	FirstName 	string 	
	LastName 	string 
	Email 		string 
	Password 	string 	
}

func (params *UserDBCreateParams) Validate() map[string]string {
	errors := map[string]string{}
	if len(params.FirstName) < minFirstNameLen {
		errors["firstName"] = fmt.Sprintf("firstName length should be at least %d characters", minFirstNameLen)
	}
	if len(params.LastName) < minLastNameLen {
		errors["lastName"] =fmt.Sprintf("lastName length should be at least %d characters", minLastNameLen)
	}
	if len(params.Password) < minPasswordLen {
		errors["password"] = fmt.Sprintf("password length should be at least %d characters", minPasswordLen)
	}
	if !IsEmailValid(params.Email) {
		errors["email"] = fmt.Sprintf("email %s is invalid", params.Email)
	}
	return errors
}

type UserDBUpdateParams struct {
	FirstName 	      	string 
	LastName 		  	string 
	Email 				string 	
	PhoneNumber 		string
	Password 			string 	
}

func (p *UserDBUpdateParams) ToBSON() bson.M {
	m := bson.M{}
	if len(p.FirstName) > 0 {
		m["firstName"] = p.FirstName
	}
	if len(p.LastName) > 0 {
		m["lastName"] = p.LastName
	}
	if len(p.Password) > 0 {
		m["EncryptedPassword"] = p.Password
	}
	return m
}