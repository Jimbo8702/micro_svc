package types

type Account struct {
	ID 							string `json:"id,omitempty"`
	UserID 						string `json:"userID"`
	StripeCustomerID 			string `json:"stripeID"`
	StripeSubscriptionID 		string `json:"subscriptionID"`
	SubscriptionStatus 			string `json:"subscriptionStatus"`
	Plan 						string `json:"plan"`
}

type UpdateAccountParams struct {
	FilterBy 			string  `json:"update_by"`
	FilterItem 			string  `json:"filter_item"`
	StripeCustomerID 			string `json:"stripeID"`
	StripeSubscriptionID 		string `json:"subcriptionID"`
	SubscriptionStatus 			string `json:"subscriptionStatus"`
	Plan 						string `json:"plan"`
}

type CreateAccountParams struct {
	UserID 						string `json:"userID"`
	StripeCustomerID 			string `json:"stripeID"`
	StripeSubscriptionID 		string `json:"subscriptionID"`
	SubscriptionStatus 			string `json:"subscriptionStatus"`
	Plan 						string `json:"plan"`
}
