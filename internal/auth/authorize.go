package auth

import (
	"github.com/codigician/auth/internal/token"
)

func (s *Service) Authorize(user *User) (string, string, error) {
	j := token.NewJWT()
	tokenService := token.New(j)
	atString, rtString, err := tokenService.Creator.CreateTokenPair(&token.Claims{
		ID: user.ID,
	})
	return atString, rtString, err
}
