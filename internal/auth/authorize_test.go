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

	mockTokenRepository := newMockTokenRepository(t)
	tokenCreator := token.NewCreator()

	ts := token.New(tokenCreator, mockTokenRepository)
	s := auth.New(nil, nil, ts)

	mockTokenRepository.EXPECT().Save(context.Background(), gomock.Any()).Return(nil).Times(1)

	tokens, err := s.Authorize(context.Background(), &u)

	assert.NotEmpty(t, tokens)
	assert.Nil(t, err)
}

func newMockTokenRepository(t *testing.T) *mocks.MockRepository {
	return mocks.NewMockRepository(gomock.NewController(t))
}
