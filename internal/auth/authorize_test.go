package auth_test

import (
	"testing"

	"github.com/codigician/auth/internal/auth"
	"github.com/codigician/auth/internal/token"
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

	ts := token.New(token.NewCreator())
	s := auth.New(nil, nil, ts)
	tokens, err := s.Authorize(&u)

	assert.NotEmpty(t, tokens)
	assert.Nil(t, err)
}
