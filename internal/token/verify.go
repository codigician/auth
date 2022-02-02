package token

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

func (s *Service) VerifyAccessToken(accessToken string) error {
	_, err := jwt.Parse(accessToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodEd25519); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return nil, nil
	})
	return err
}

func (s *Service) VerifyRefreshToken(ctx context.Context, id string) error {
	refreshToken, err := s.repository.Get(ctx, id)
	if err != nil {
		fmt.Println("no refresh token with the given id found:", err)
		return err
	}
	expirationDate := time.Unix(refreshToken.ExpirationDate, 0)
	if time.Until(expirationDate).Hours() <= 0 {
		return errors.New("refresh token expired")
	}
	return nil
}
