package main

import (
	"Jimbo8702/randomThoughts/diggity-dawg/types"
	"fmt"
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

const (
	bcryptCost = 12
	minFirstNameLen = 2
	minLastNameLen  = 2
	minPasswordLen 	= 7
)

func NewPassword(password string) (string, error) {
	newHash, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return "", err
	}
	return string(newHash), nil
}

//see if the passwords match
func PasswordMatches(enpw, pw string) bool {
	return bcrypt.CompareHashAndPassword([]byte(enpw), []byte(pw)) == nil
}

// is valid email
func IsEmailValid(e string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(e)
}

func ValidateCreateUserParams(params types.CreateUserParams)  map[string]string {
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