package token

type (
	Creator interface {
		CreateTokenPair(c *Claims) (string, string, error)
	}

	Service struct {
		Creator Creator
	}
)

func New(creator Creator) *Service {
	return &Service{Creator: creator}
}
