package token

import (
	"crypto/ed25519"
	"time"

	"github.com/golang-jwt/jwt"
)

const (
	issuer   = "codigician"
	audience = "https://codigician.com"
)

type (
	Claims struct {
		ID    string
		Email string
	}

	JWT struct {
		Issuer                     string
		PrivateKey                 ed25519.PrivateKey
		PublicKey                  ed25519.PublicKey
		AccessTokenExpireDuration  time.Duration
		RefreshTokenExpireDuration time.Duration
	}
)

func NewJWT() (*JWT, error) {
	publicKey, privateKey, err := ed25519.GenerateKey(nil)

	return &JWT{
		Issuer:                     issuer,
		PrivateKey:                 privateKey,
		PublicKey:                  publicKey,
		AccessTokenExpireDuration:  time.Minute * 15,
		RefreshTokenExpireDuration: time.Hour * 24 * 14,
	}, err
}

func (j *JWT) Create(c *Claims) (string, error) {
	claims := jwt.StandardClaims{
		Audience:  audience,
		ExpiresAt: time.Now().Add(j.AccessTokenExpireDuration).Unix(),
		Id:        c.ID,
		IssuedAt:  time.Now().Unix(),
		Issuer:    j.Issuer,
		Subject:   c.Email,
	}
	token := jwt.NewWithClaims(&jwt.SigningMethodEd25519{}, claims)
	tokenString, err := token.SignedString(j.PrivateKey)
	return tokenString, err
}

// func Validate(tokenString string) error {
// 	var privateKey []byte
// 	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 		if _, ok := token.Method.(*jwt.SigningMethodEd25519); !ok {
// 			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
// 		}
// 		return privateKey, nil
// 	})
// 	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
// 		return nil
// 	} else {
// 		return err
// 	}
// }