package auth

import (
	"github.com/codigician/auth/internal/token"
)

func (s *Service) Authorize(user *User) (atString, rtString string, err error) {
	j := token.NewJWT()
	tokenService := token.New(j)
	accessTokenString, refreshTokenString, err := tokenService.Creator.CreateTokenPair(&token.Claims{
		ID: user.ID,
	})
	return accessTokenString, refreshTokenString, err
}
