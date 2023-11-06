package main

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2/bson"
)

type DBAccount struct {
	ID 							primitive.ObjectID 	`bson:"_id,omitempty"`
	UserID 						primitive.ObjectID 	`bson:"user_id"`
	StripeCustomerID 			string 				`bson:"stripe_id"`
	StripeSubscriptionID 		string 				`bson:"subscription_id"`
	SubscriptionStatus 			string 				`bson:"subscription_status"`
	Plan 						string 				`bson:"plan"`
}

type AccountDBCreateParams struct {
	UserID 						string 
	StripeCustomerID 			string 
	StripeSubscriptionID 		string 
	SubscriptionStatus 			string 
	Plan 						string 
}

func (a *AccountDBCreateParams) Validate() map[string]string {
	errors := map[string]string{}
	if len(a.UserID) <= 0 {
		errors["user_id"] = "user ID is required to create account"
	}
	// add more validations ehre
	return errors
}

type AccountDBUpdateParams struct {
	StripeCustomerID 			string 
	StripeSubscriptionID 		string 
	SubscriptionStatus 			string 
	Plan 						string 
}

func (a *AccountDBUpdateParams) ToBSON() bson.M {
	m := bson.M{}
	if len(a.StripeCustomerID) > 0 {
		m["stripe_id"] = a.StripeCustomerID
	}
	if len(a.StripeSubscriptionID) > 0 {
		m["subscription_id"] = a.StripeSubscriptionID
	}
	if len(a.SubscriptionStatus) > 0 {
		m["subscription_status"] = a.SubscriptionStatus
	}
	if len(a.Plan) > 0 {
		m["plan"] = a.Plan
	}
	return m
}