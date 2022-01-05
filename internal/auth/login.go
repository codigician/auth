package auth

import (
	"crypto/tls"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"golang.org/x/crypto/bcrypt"
	gomail "gopkg.in/mail.v2"
)

type UserCredentials struct {
	Email    string
	Password string
}

type Login struct {
	repository Repository
}

func NewLogin(repo Repository) *Login {
	return &Login{repository: repo}
}

func (l *Login) Authenticate(c *UserCredentials) error {
	user, err := l.repository.Get(c.Email)
	if err != nil {
		return errors.New("email address doesn't exist")
	}
	return bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(c.Password))
}

func (l *Login) ForgotPassword(email string, d Data) error {
	_, err := l.repository.Get(email)
	if err != nil {
		log.Fatal(err)
	}

	if err := d.SendPasswordResetEmail(); err != nil {
		log.Fatal(err)
	}
	return err
}

func (data Data) CheckCode(inputtedCode string) error {
	if inputtedCode != data.Code {
		return errors.New("wrong code")
	}
	fmt.Println("code accepted")
	return nil
}

type Data struct {
	Code string
}

func (data Data) SendPasswordResetEmail() error {
	m := gomail.NewMessage()
	m.SetHeader("From", "gokcelbilgin@gmail.com")
	m.SetHeader("To", "gokcelbilgin@gmail.com")
	m.SetHeader("Subject", "Gomail test")
	m.SetBody("text/plain", fmt.Sprintf("Your code: %s", data.Code))

	d := gomail.NewDialer("smtp.gmail.com", 587, "gokcelbilgin@gmail.com", os.Getenv("APP_PSW"))
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	err := d.DialAndSend(m)
	if err != nil {
		fmt.Println(err)
	}
	return err
}

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

const numbers = "1234567890"

func GenerateCode(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = numbers[seededRand.Intn(len(numbers))]
	}
	return string(b)
}