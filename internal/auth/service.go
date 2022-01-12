package auth

import (
	"context"

	"github.com/codigician/auth/internal/token"
)

type (
	Repository interface {
		Save(ctx context.Context, u *User) error
		Get(ctx context.Context, email string) (*User, error)
		Update(ctx context.Context, id string, fields map[string]interface{}) error
		Delete(ctx context.Context, id string) error
	}

	Mailer interface {
		Mail(to, subject, body string) error
	}

	TokenService interface {
		CreateTokens(id string) (*token.TokenPair, error)
	}

	Service struct {
		repository   Repository
		mailer       Mailer
		tokenService TokenService
	}
)

func New(repo Repository, mailer Mailer, tokenService TokenService) *Service {
	return &Service{
		repository:   repo,
		mailer:       mailer,
		tokenService: tokenService,
	}
}
