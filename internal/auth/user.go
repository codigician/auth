package auth

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID             string
	Firstname      string
	Lastname       string
	Email          string
	HashedPassword string
}

const _cost = 14

func NewUser(info *RegistrationInfo) *User {
	user := &User{
		Firstname: info.Firstname,
		Lastname:  info.Lastname,
		Email:     info.Email,
	}
	hashedPassword := HashPassword(info.Password)
	user.HashedPassword = hashedPassword
	return user
}

func HashPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), _cost)
	return string(bytes)
}

func (u *User) ComparePassword(rawPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.HashedPassword), []byte(rawPassword))
}
