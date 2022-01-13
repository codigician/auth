package auth_test

import (
	"context"
	"testing"

	"github.com/codigician/auth/internal/auth"
	mocks "github.com/codigician/auth/internal/mocks/token"
	"github.com/codigician/auth/internal/token"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestAuthorize_SuccessfulTokenCreation_ReturnsTokensMapAndNil(t *testing.T) {
	u := auth.User{
		ID:             "someid",
		Firstname:      "somename",
		Lastname:       "somelastname",
		Email:          "someemail@outlook.com",
		HashedPassword: "somehashedpassword",
	}

	ts := token.New(token.NewCreator(), newMockTokenRepository(t))
	s := auth.New(nil, nil, ts)
	tokens, err := s.Authorize(context.Background(), &u)

	assert.NotEmpty(t, tokens)
	assert.Nil(t, err)
}

func newMockTokenRepository(t *testing.T) *mocks.MockRepository {
	return mocks.NewMockRepository(gomock.NewController(t))
}
