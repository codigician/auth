package token_test

import (
	"context"
	"errors"
	"testing"

	"github.com/codigician/auth/internal/token"
	"github.com/stretchr/testify/assert"
)

var ctx = context.Background()

var refreshToken = token.RefreshToken{
	ID:             id,
	Token:          "ae5b94b4-902c-4b2d-aaa6-3e0b2c6f74e2",
	ExpirationDate: 1643293722,
}

func TestValidateAccessToken_Validated_ReturnsNil(t *testing.T) {
	accessTokenString := creator.GenerateAccessToken(id)
	err := service.ValidateAccessToken(accessTokenString)

	assert.Nil(t, err)
}

func TestValidateRefreshToken_Validated_ReturnsNil(t *testing.T) {
	mockRepository.EXPECT().Get(ctx, id).Return(&refreshToken, nil).Times(1)

	err := service.ValidateRefreshToken(ctx, id)

	assert.Nil(t, err)
}

func TestValidateRefreshToken_GetFailed_ReturnsError(t *testing.T) {
	noDocumentsErr := errors.New("mongo: no documents in result")

	mockRepository.EXPECT().Get(ctx, id).
		Return(nil, noDocumentsErr).Times(1)

	err := service.ValidateRefreshToken(ctx, id)

	assert.EqualError(t, err, noDocumentsErr.Error())
}

func TestValidateRefreshToken_ExpiredToken_ReturnsError(t *testing.T) {
	tokenExpiredErr := errors.New("refresh token expired")
	expiredToken := token.RefreshToken{
		ID:             id,
		Token:          "77e7a37d-3789-442e-845c-83eaab683444",
		ExpirationDate: 1641810890,
	}

	mockRepository.EXPECT().Get(ctx, id).
		Return(&expiredToken, nil).Times(1)

	err := service.ValidateRefreshToken(ctx, id)

	assert.EqualError(t, err, tokenExpiredErr.Error())
}
