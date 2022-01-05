package mail

import (
	"crypto/tls"
	"os"

	gomail "gopkg.in/mail.v2"
)

type Type string

const (
	Verification = iota + 1
	PasswordReset

	_host = "smtp.gmail.com"
	_port = 587
	_from = "gokcelbilgin@gmail.com"
)

type Mailer struct {
	dialer *gomail.Dialer
}

func New() *Mailer {
	d := gomail.NewDialer(_host, _port, _from, os.Getenv("APP_PSW"))
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	return &Mailer{
		dialer: d,
	}
}

func (m *Mailer) Mail(to, subject, body string) error {
	msg := gomail.NewMessage()
	msg.SetHeader("From", _from)
	msg.SetHeader("To", to)
	msg.SetHeader("Subject", subject)
	msg.SetBody("text/plain", body)

	return m.dialer.DialAndSend(msg)
}
