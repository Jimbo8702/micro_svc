package main

import "gopkg.in/mgo.v2/bson"

type DBFilter map[string]any

type UserDBUpdateParams struct {
	FirstName 	      	string 
	LastName 		  	string 
	Email 				string 	
	PhoneNumber 		string
	EncryptedPassword 	string 	
}

func (p UserDBUpdateParams) ToBSON() bson.M {
	m := bson.M{}
	if len(p.FirstName) > 0 {
		m["firstName"] = p.FirstName
	}
	if len(p.LastName) > 0 {
		m["lastName"] = p.LastName
	}
	if len(p.EncryptedPassword) > 0 {
		m["EncryptedPassword"] = p.EncryptedPassword
	}
	return m
}