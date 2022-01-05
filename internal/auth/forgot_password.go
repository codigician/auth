package auth

import (
	"context"
	"fmt"
	"log"
)

func (s *Service) ForgotPassword(ctx context.Context, email string) error {
	u, err := s.repository.Get(ctx, email)
	if err != nil {
		log.Printf("repository get email: %s failed, err: %v\n", email, err)
		return err
	}
	return s.SendPasswordResetEmail(u)
}

func (s *Service) SendPasswordResetEmail(u *User) error {
	link := fmt.Sprintf("http://localhost:8888/password-reset/%s", u.ID)
	body := fmt.Sprintf("Click the link below to reset your password\nPassword reset link: %s", link)
	return s.mailer.Mail(u.Email, "Password Reset", body)
}
