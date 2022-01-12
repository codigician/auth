package token_test

import (
	"testing"

	"github.com/codigician/auth/internal/token"
	"github.com/stretchr/testify/assert"
)

func TestCreateTokens_SuccessfulCreation_ReturnsTokensAndNil(t *testing.T) {
	s := token.New(token.NewCreator())
	tokens, err := s.CreateTokens(id)

	assert.NotEmpty(t, tokens)
	assert.Nil(t, err)
}
