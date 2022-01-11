package token_test

import (
	"testing"

	"github.com/codigician/auth/internal/token"
	"github.com/stretchr/testify/assert"
)

func TestCreate_ValidKey_ReturnsAccessAndRefreshTokenStrings(t *testing.T) {
	j := token.NewJWT()
	s := token.New(j)
	atString, rtString, err := s.Creator.CreateTokenPair(&token.Claims{
		ID: "01234",
	})
	assert.NotEmpty(t, atString)
	assert.NotEmpty(t, rtString)
	assert.Nil(t, err)
}
