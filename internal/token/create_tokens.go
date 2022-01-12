package token

type TokenPair struct {
	AT string
	RT string
}

func (s *Service) CreateTokens(id string) (*TokenPair, error) {
	accessTokenString, err := s.Creator.GenerateAccessToken(id)
	refreshToken := s.Creator.GenerateRefreshToken(id)
	return &TokenPair{
		AT: accessTokenString,
		RT: refreshToken.TokenString,
	}, err
}
