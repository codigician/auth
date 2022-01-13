package token_test

import (
	"context"
	"testing"

	mocks "github.com/codigician/auth/internal/mocks/token"
	"github.com/codigician/auth/internal/token"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

var mockRepository = newMockRepository(&testing.T{})
var creator = token.NewCreator()
var service = token.New(creator, mockRepository)

func TestCreateTokens_SuccessfulCreation_ReturnsTokensAndNil(t *testing.T) {

	mockRepository.EXPECT().Save(context.Background(), gomock.Any()).Return(nil).Times(1)

	tokens, err := service.CreateTokens(context.Background(), id)

	assert.NotEmpty(t, tokens)
	assert.Nil(t, err)
}

func newMockRepository(t *testing.T) *mocks.MockRepository {
	return mocks.NewMockRepository(gomock.NewController(t))
}
