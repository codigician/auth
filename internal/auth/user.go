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

func NewUser(info *RegistrationInfo) (*User, error) {
	user := &User{
		Firstname: info.Firstname,
		Lastname:  info.Lastname,
		Email:     info.Email,
	}
	hashedPassword, err := HashPassword(info.Password)
	user.HashedPassword = hashedPassword
	return user, err
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (u *User) ComparePassword(rawPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.HashedPassword), []byte(rawPassword))
}
