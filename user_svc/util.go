package main

import (
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
	newHash, err := bcrypt.GenerateFromPassword([]byte(password), bcryptCost)
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

