package auth_test

import (
	"testing"

	"github.com/codigician/auth/internal/auth"
	"github.com/stretchr/testify/assert"
)

func TestAuthorize_SuccessfulAuthorization_ReturnsTokenStringAndNil(t *testing.T) {
	s := auth.New(nil, nil)
	atString, rtString, err := s.Authorize(&auth.User{
		ID:             "someid",
		Firstname:      "somename",
		Lastname:       "somelastname",
		Email:          "someemail@outlook.com",
		HashedPassword: "somehashedpassword",
	})

	assert.NotEmpty(t, atString)
	assert.NotEmpty(t, rtString)
	assert.Nil(t, err)
}
