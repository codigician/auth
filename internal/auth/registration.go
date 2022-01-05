package auth

import (
	"crypto/tls"
	"fmt"
	"os"

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
	ID             uuid.UUID `bson: "_id"`
	Firstname      string    `bson: "firstname"`
	Lastname       string    `bson: "lastname"`
	Email          string    `bson: "email"`
	HashedPassword string    `bson: "hashed_password`
}

type Repository interface {
	Save(u *User) error
	Find(uniqueVal string) (*User, error)
	Update(id uuid.UUID, key string, val string) error
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
	if err := r.SendVerificationEmail(&u); err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *Registrator) SendVerificationEmail(u *User) error {
	m := gomail.NewMessage()
	m.SetHeader("From", "gokcelbilgin@gmail.com")
	m.SetHeader("To", "gokcelbilgin@gmail.com")
	m.SetHeader("Subject", "Gomail test")
	m.SetBody("text/plain", fmt.Sprintf("Hello, %s %s. This is your verification email.", u.Firstname, u.Lastname))

	d := gomail.NewDialer("smtp.gmail.com", 587, "gokcelbilgin@gmail.com", os.Getenv("APP_PSW"))
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
