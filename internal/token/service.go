package token

import (
	"context"
)

type (
	Issuer interface {
		GenerateAccessToken(id string) string
		GenerateRefreshToken(id string) *RefreshToken
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
