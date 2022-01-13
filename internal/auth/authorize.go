package auth

import (
	"context"

	"github.com/codigician/auth/internal/token"
)

func (s *Service) Authorize(ctx context.Context, user *User) (*token.Pair, error) {
	return s.tokenService.CreateTokens(ctx, user.ID)
}
