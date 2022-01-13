package token

import (
	"context"
	"crypto/ed25519"
)

type (
	Issuer interface {
		GenerateAccessToken(id string) (string, error)
		GenerateRefreshToken(id string) *RefreshToken
		PrivateKey() ed25519.PrivateKey
	}

	Repository interface {
		Save(ctx context.Context, refreshToken *RefreshToken) error
		Get(ctx context.Context, id string) (*RefreshToken, error)
	}

	Service struct {
		issuer     Issuer
		repository Repository
	}
)

func New(issuer Issuer, repo Repository) *Service {
	return &Service{
		issuer:     issuer,
		repository: repo,
	}
}
