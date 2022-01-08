package auth

import (
	"fmt"

	"github.com/codigician/auth/internal/token"
)

func (s *Service) Authorize(user *User) (string, error) {
	j, err := token.NewJWT()
	tokenService := token.New(j)

	if err != nil {
		fmt.Println("jwt could not be generated:", err)
		return "", err
	}
	tokenString, err := tokenService.Creator.Create(&token.Claims{
		ID:    user.ID,
		Email: user.Email,
	})
	if err != nil {
		fmt.Println("token could not be generated:", err)
		return "", err
	}
	return tokenString, nil
}
