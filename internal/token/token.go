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
		issuer                     string
		privateKey                 ed25519.PrivateKey
		publicKey                  ed25519.PublicKey
		accessTokenExpireDuration  time.Duration
		refreshToken               string
		refreshTokenExpireDuration time.Duration
	}

	RefreshToken struct {
		ID             string
		Token          string
		ExpirationDate int64
	}
)

func NewCreator() *Creator {
	publicKey, privateKey, _ := ed25519.GenerateKey(nil)

	return &Creator{
		issuer:                     issuer,
		privateKey:                 privateKey,
		publicKey:                  publicKey,
		accessTokenExpireDuration:  time.Minute * 15,
		refreshToken:               uuid.NewString(),
		refreshTokenExpireDuration: time.Hour * 24 * 14,
	}
}

func (c *Creator) GenerateAccessToken(id string) (string, error) {
	tokenClaims := jwt.StandardClaims{
		Audience:  audience,
		ExpiresAt: time.Now().Add(c.accessTokenExpireDuration).Unix(),
		Id:        id,
		IssuedAt:  time.Now().Unix(),
		Issuer:    c.issuer,
	}
	token := jwt.NewWithClaims(&jwt.SigningMethodEd25519{}, tokenClaims)
	return token.SignedString(c.PrivateKey)
}

func (c *Creator) GenerateRefreshToken(id string) *RefreshToken {
	return &RefreshToken{
		ID:             id,
		Token:          c.refreshToken,
		ExpirationDate: time.Now().Add(c.refreshTokenExpireDuration).Unix(),
	}
}

func (c *Creator) PrivateKey() ed25519.PrivateKey {
	return c.privateKey
}
