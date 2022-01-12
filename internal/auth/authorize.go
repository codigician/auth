package auth

import "github.com/codigician/auth/internal/token"

func (s *Service) Authorize(user *User) (*token.Pair, error) {
	return s.tokenService.CreateTokens(user.ID)
}
