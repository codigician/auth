package auth

import (
	"github.com/codigician/auth/internal/token"
)

func (s *Service) Authorize(user *User) (string, error) {
	j := token.NewJWT()
	tokenService := token.New(j)
	tokenString, err := tokenService.Creator.Create(&token.Claims{
		ID: user.ID,
	})
	return tokenString, err
}
