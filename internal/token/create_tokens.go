package token

type Pair struct {
	AccessToken  string
	RefreshToken string
}

func (s *Service) CreateTokens(id string) (*Pair, error) {
	accessTokenString, err := s.issuer.GenerateAccessToken(id)
	refreshToken := s.issuer.GenerateRefreshToken(id)
	return &Pair{
		AccessToken:  accessTokenString,
		RefreshToken: refreshToken.TokenString,
	}, err
}
