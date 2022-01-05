package auth

import "context"

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

	Service struct {
		repository Repository
		mailer     Mailer
	}
)

func New(repo Repository, mailer Mailer) *Service {
	return &Service{
		repository: repo,
		mailer:     mailer,
	}
}
