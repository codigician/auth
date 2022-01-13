package token_test

import (
	"context"
	"testing"

	mocks "github.com/codigician/auth/internal/mocks/token"
	"github.com/codigician/auth/internal/token"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreateTokens_SuccessfulCreation_ReturnsTokensAndNil(t *testing.T) {
	mockRepository := newMockRepository(t)
	creator := token.NewCreator()
	s := token.New(creator, mockRepository)

	mockRepository.EXPECT().Save(context.Background(), gomock.Any()).Return(nil).Times(1)

	tokens, err := s.CreateTokens(context.Background(), id)

	assert.NotEmpty(t, tokens)
	assert.Nil(t, err)
}

func newMockRepository(t *testing.T) *mocks.MockRepository {
	return mocks.NewMockRepository(gomock.NewController(t))
}
