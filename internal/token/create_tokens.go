package token

import "context"

type Pair struct {
	AccessToken  string
	RefreshToken string
}

func (s *Service) CreateTokens(ctx context.Context, id string) (*Pair, error) {
	accessTokenString := s.issuer.GenerateAccessToken(id)
	refreshToken := s.issuer.GenerateRefreshToken(id)
	err := s.repository.Save(ctx, refreshToken)
	return &Pair{
		AccessToken:  accessTokenString,
		RefreshToken: refreshToken.Token,
	}, err
}
