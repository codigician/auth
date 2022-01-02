package auth

import (
	"errors"
	"regexp"
	"strings"
)

const emailRegExp = "^[\\w-\\.]+@([\\w-]+\\.)+[\\w-]{2,4}$"

func isValidPassword(password string) error {
	if !strings.ContainsAny(password, "abcdefghijklmnoprstuvyzwxqABCDEFGHIJKLMNOPRSTUVYZWXQ") {
		return errors.New("password must contain alphabetic characters")
	}
	if !strings.ContainsAny(password, "abcdefghijklmnoprstuvyzwxq") {
		return errors.New("password must contain lowercase letters")
	}
	if !strings.ContainsAny(password, "ABCDEFGHIJKLMNOPRSTUVYZWXQ") {
		return errors.New("password must contain uppercase letters")
	}
	if !strings.ContainsAny(password, "1234567890") {
		return errors.New("password must contain numbers")
	}
	if !strings.ContainsAny(password, "!@#$%&*()-_+={}[]|:;<>,.?") {
		return errors.New("password must contain special characters")
	}
	if strings.ContainsAny(password, "`'\"/\\") {
		return errors.New("password can't contain special characters `, ', \", / or \\")
	}
	return nil
}

func Register(r RegistrationInfo) (*User, error) {
	if len(r.Firstname) > 100 {
		return nil, errors.New("first name can't be longer than 100 characters")
	}
	if len(r.Lastname) > 100 {
		return nil, errors.New("last name can't be longer than 100 characters")
	}
	regExp, _ := regexp.Compile(emailRegExp)
	if !regExp.MatchString(r.Email) {
		return nil, errors.New("invalid email address")
	}
	err := isValidPassword(r.Password)
	if err != nil {
		return nil, err
	}
	return &User{}, nil
}

type RegistrationInfo struct {
	Firstname string
	Lastname  string
	Email     string
	Password  string
}

type User struct {
	id             string
	firstname      string
	lastname       string
	email          string
	hashedPassword string
}
