package token_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/codigician/auth/internal/token"
)

var c = token.NewTokenCreator()
var id = "01234"

func TestAccessToken_SuccessfulCreation_ReturnsTokenString(t *testing.T) {
	accessTokenString, err := c.GenerateAccessToken(id)

	assert.NotEmpty(t, accessTokenString)
	assert.Nil(t, err)
}

func TestRefreshToken_SuccessfulCreation_ReturnsToken(t *testing.T) {
	refreshToken := c.GenerateRefreshToken(id)

	assert.IsType(t, &token.RefreshToken{}, refreshToken)
	assert.Equal(t, id, refreshToken.ID)
}
