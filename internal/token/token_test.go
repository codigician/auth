package token_test

import (
	"crypto/ed25519"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/codigician/auth/internal/token"
)

var id = "61d6ec74eeaacbe4224a0f83"

func TestGenerateAccessToken_SuccessfulCreation_ReturnsTokenString(t *testing.T) {
	accessTokenString := creator.GenerateAccessToken(id)

	assert.NotEmpty(t, accessTokenString)
}

func TestGenerateRefreshToken_SuccessfulCreation_ReturnsToken(t *testing.T) {
	refreshToken := creator.GenerateRefreshToken(id)

	assert.IsType(t, &token.RefreshToken{}, refreshToken)
	assert.Equal(t, id, refreshToken.ID)
}

func TestPrivateKey(t *testing.T) {
	privateKey := creator.PrivateKey()

	assert.NotEmpty(t, privateKey)
	assert.IsType(t, ed25519.PrivateKey{}, privateKey)
}
