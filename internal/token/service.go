package token

type (
	Issuer interface {
		GenerateAccessToken(id string) (string, error)
		GenerateRefreshToken(id string) *RefreshToken
	}

	Service struct {
		issuer Issuer
	}
)

func New(issuer Issuer) *Service {
	return &Service{issuer}
}
