package main

type AuthenticationService interface {
	Authenticate()
	RevokeAuthentication()
}