package token

import (
	"crypto/ed25519"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

const (
	issuer   = "codigician"
	audience = "https://codigician.com"
)

type (
	TokenCreator struct {
		Issuer                     string
		PrivateKey                 ed25519.PrivateKey
		PublicKey                  ed25519.PublicKey
		AccessTokenExpireDuration  time.Duration
		RefreshTokenString         string
		RefreshTokenExpireDuration time.Duration
	}

	RefreshToken struct {
		ID             string
		TokenString    string
		ExpirationDate int64
	}
)

func NewTokenCreator() *TokenCreator {
	publicKey, privateKey, _ := ed25519.GenerateKey(nil)

	return &TokenCreator{
		Issuer:                     issuer,
		PrivateKey:                 privateKey,
		PublicKey:                  publicKey,
		AccessTokenExpireDuration:  time.Minute * 15,
		RefreshTokenString:         uuid.NewString(),
		RefreshTokenExpireDuration: time.Hour * 24 * 14,
	}
}

func (tc *TokenCreator) GenerateAccessToken(id string) (string, error) {
	tokenClaims := jwt.StandardClaims{
		Audience:  audience,
		ExpiresAt: time.Now().Add(tc.AccessTokenExpireDuration).Unix(),
		Id:        id,
		IssuedAt:  time.Now().Unix(),
		Issuer:    tc.Issuer,
	}
	token := jwt.NewWithClaims(&jwt.SigningMethodEd25519{}, tokenClaims)
	return token.SignedString(tc.PrivateKey)
}

func (tc *TokenCreator) GenerateRefreshToken(id string) *RefreshToken {
	return &RefreshToken{
		ID:             id,
		TokenString:    tc.RefreshTokenString,
		ExpirationDate: time.Now().Add(tc.RefreshTokenExpireDuration).Unix(),
	}
}
