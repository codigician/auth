package auth

import (
	"crypto/tls"
	"fmt"

	"github.com/docker/distribution/uuid"
	"golang.org/x/crypto/bcrypt"
	gomail "gopkg.in/mail.v2"
)

type RegistrationInfo struct {
	Firstname string
	Lastname  string
	Email     string
	Password  string
}

type User struct {
	ID             uuid.UUID
	Firstname      string
	Lastname       string
	Email          string
	HashedPassword string
}

type Repository interface {
	Save(u *User) error
}

type Registrator struct {
	repository Repository
}

func NewRegistrator(repo Repository) *Registrator {
	return &Registrator{
		repository: repo,
	}
}

func (r *Registrator) Register(ri RegistrationInfo) (*User, error) {
	id := uuid.Generate()
	hashedPsw, err := HashPassword(ri.Password)
	if err != nil {
		return nil, err
	}

	u := User{
		ID:             id,
		Firstname:      ri.Firstname,
		Lastname:       ri.Lastname,
		Email:          ri.Email,
		HashedPassword: hashedPsw,
	}
	if err := r.repository.Save(&u); err != nil {
		return nil, err
	}

	if err := r.Verify(); err != nil {
		return nil, err
	}

	return &u, nil
}

func (r *Registrator) Verify() error {
	m := gomail.NewMessage()
	m.SetHeader("From", "gokcelb@hotmail.com")
	m.SetHeader("To", "gokcelbilgin@gmail.com")
	m.SetHeader("Subject", "Gomail test")
	m.SetBody("text/plain", "This is your verification email.")

	d := gomail.NewDialer("smtp.gmail.com", 587, "gokcelb@hotmail.com", "Lacmucum.96")
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	err := d.DialAndSend(m)
	if err != nil {
		fmt.Println(err)
	}
	return err
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
