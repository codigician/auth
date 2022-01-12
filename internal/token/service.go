package token

type (
	Creator interface {
		GenerateAccessToken(id string) (string, error)
		GenerateRefreshToken(id string) *RefreshToken
	}

	Service struct {
		Creator Creator
	}
)

func New(creator Creator) *Service {
	return &Service{Creator: creator}
}
