package token_test

import (
	"testing"

	"github.com/codigician/auth/internal/token"
	"github.com/stretchr/testify/assert"
)

func TestCreate_ValidKey_ReturnsTokenString(t *testing.T) {
	j, _ := token.NewJWT()
	s := token.New(j)
	tokenString, err := s.Creator.Create(&token.Claims{
		ID:    "01234",
		Email: "noone@outlook.com",
	})
	assert.NotEmpty(t, tokenString)
	assert.Nil(t, err)
}
