package auth

import (
	"context"
	"fmt"
	"log"
)

type RegistrationInfo struct {
	Firstname string
	Lastname  string
	Email     string
	Password  string
}

func (s *Service) Register(ctx context.Context, info *RegistrationInfo) (*User, error) {
	user := NewUser(info)

	if err := s.repository.Save(ctx, user); err != nil {
		log.Printf("user repository save: %v\n", err)
		return nil, err
	}

	return user, s.SendVerificationEmail(user)
}

func (s *Service) SendVerificationEmail(u *User) error {
	body := fmt.Sprintf("Hello, %s %s. This is your verification email.", u.Firstname, u.Lastname)
	return s.mailer.Mail(u.Email, "Verification", body)
}
