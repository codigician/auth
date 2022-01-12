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
	Creator struct {
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

func NewCreator() *Creator {
	publicKey, privateKey, _ := ed25519.GenerateKey(nil)

	return &Creator{
		Issuer:                     issuer,
		PrivateKey:                 privateKey,
		PublicKey:                  publicKey,
		AccessTokenExpireDuration:  time.Minute * 15,
		RefreshTokenString:         uuid.NewString(),
		RefreshTokenExpireDuration: time.Hour * 24 * 14,
	}
}

func (c *Creator) GenerateAccessToken(id string) (string, error) {
	tokenClaims := jwt.StandardClaims{
		Audience:  audience,
		ExpiresAt: time.Now().Add(c.AccessTokenExpireDuration).Unix(),
		Id:        id,
		IssuedAt:  time.Now().Unix(),
		Issuer:    c.Issuer,
	}
	token := jwt.NewWithClaims(&jwt.SigningMethodEd25519{}, tokenClaims)
	return token.SignedString(c.PrivateKey)
}

func (c *Creator) GenerateRefreshToken(id string) *RefreshToken {
	return &RefreshToken{
		ID:             id,
		TokenString:    c.RefreshTokenString,
		ExpirationDate: time.Now().Add(c.RefreshTokenExpireDuration).Unix(),
	}
}
