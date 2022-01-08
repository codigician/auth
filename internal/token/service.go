package token

type (
	Creator interface {
		Create(c *Claims) (string, error)
	}

	Service struct {
		Creator Creator
	}
)

func New(creator Creator) *Service {
	return &Service{Creator: creator}
}
