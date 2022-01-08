package auth_test

import (
	"testing"

	"github.com/codigician/auth/internal/auth"
	"github.com/stretchr/testify/assert"
)

func TestAuthorize_SuccessfulAuthorization_ReturnsTokenStringAndNil(t *testing.T) {
	s := auth.New(nil, nil)
	tokenString, err := s.Authorize(&auth.User{
		ID:             "someid",
		Firstname:      "somename",
		Lastname:       "somelastname",
		Email:          "someemail@outlook.com",
		HashedPassword: "somehashedpassword",
	})

	assert.NotEmpty(t, tokenString)
	assert.Nil(t, err)
}
