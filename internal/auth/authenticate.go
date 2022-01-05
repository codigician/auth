package auth

import (
	"context"
	"errors"
	"log"
)

type UserCredentials struct {
	Email    string
	Password string
}

func (s *Service) Authenticate(ctx context.Context, c *UserCredentials) error {
	user, err := s.repository.Get(ctx, c.Email)
	if err != nil {
		log.Printf("authenticate failed for email: %v\n", c.Email)
		return errors.New("email address doesn't exist")
	}
	return user.ComparePassword(c.Password)
}
