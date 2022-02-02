package token

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

const (
	issuer   = "codigician"
	audience = "https://codigician.com"
)

type (
	Config struct {
		PrivateKeyFilePath string
		PublicKeyFilePath  string
	}

	Creator struct {
		issuer                     string
		privateKey                 *rsa.PrivateKey
		publicKey                  *rsa.PublicKey
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

func NewCreator(c Config) *Creator {
	rawPubKeyBytes, rawPriKeyBytes, err := readKeyFiles(c)
	if err != nil {
		panic(err)
	}

	pub, pri := decodeRawKeys(rawPubKeyBytes, rawPriKeyBytes)
	publicKey, privateKey, err := parseKeys(pub, pri)
	if err != nil {
		panic(err)
	}

	return &Creator{
		issuer:                     issuer,
		privateKey:                 privateKey,
		publicKey:                  publicKey,
		accessTokenExpireDuration:  time.Minute * 15,
		refreshToken:               uuid.NewString(),
		refreshTokenExpireDuration: time.Hour * 24 * 14,
	}
}

func parseKeys(pub, pri *pem.Block) (*rsa.PublicKey, *rsa.PrivateKey, error) {
	if pub == nil || pri == nil {
		return nil, nil, fmt.Errorf("public or private key is nil")
	}

	publicKey, err := x509.ParsePKCS1PublicKey(pub.Bytes)
	if err != nil {
		return nil, nil, err
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(pri.Bytes)
	if err != nil {
		return nil, nil, err
	}

	return publicKey, privateKey, nil
}

func decodeRawKeys(rawPubKey, rawPriKey []byte) (pub *pem.Block, pri *pem.Block) {
	pub, _ = pem.Decode(rawPubKey)
	pri, _ = pem.Decode(rawPriKey)

	return pub, pri
}

func readKeyFiles(conf Config) (pubKey []byte, priKey []byte, err error) {
	if pubKey, err = os.ReadFile(conf.PublicKeyFilePath); err != nil {
		return nil, nil, err
	}

	if priKey, err = os.ReadFile(conf.PrivateKeyFilePath); err != nil {
		return nil, nil, err
	}

	return pubKey, priKey, nil
}

func (c *Creator) GenerateAccessToken(id string) string {
	tokenClaims := jwt.StandardClaims{
		Audience:  audience,
		ExpiresAt: time.Now().Add(c.accessTokenExpireDuration).Unix(),
		Id:        id,
		IssuedAt:  time.Now().Unix(),
		Issuer:    c.issuer,
	}
	token := jwt.NewWithClaims(&jwt.SigningMethodEd25519{}, tokenClaims)
	tokenString, _ := token.SignedString(c.privateKey)
	fmt.Println("private key of creator:", c.privateKey)
	return tokenString
}

func (c *Creator) GenerateRefreshToken(id string) *RefreshToken {
	return &RefreshToken{
		ID:             id,
		Token:          c.refreshToken,
		ExpirationDate: time.Now().Add(c.refreshTokenExpireDuration).Unix(),
	}
}
