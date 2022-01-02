package auth_test

import (
	"strings"
	"testing"

	"github.com/codigician/auth/internal/auth"
	"github.com/stretchr/testify/assert"
)

func TestRegister_FirstnameOutOfRange_ReturnErr(t *testing.T) {
	_, err := auth.Register(auth.RegistrationInfo{
		Firstname: strings.Repeat("Lacin", 25),
	})

	assert.NotNil(t, err)
}

func TestRegister_LastnameOutOfRange_ReturnErr(t *testing.T) {
	_, err := auth.Register(auth.RegistrationInfo{
		Lastname: strings.Repeat("Bilgin", 20),
	})

	assert.NotNil(t, err)
}

func TestRegister_EmailInvalid_ReturnErr(t *testing.T) {
	_, err := auth.Register(auth.RegistrationInfo{
		Email: "house",
	})

	assert.NotNil(t, err)
}
func TestRegister_PasswordInvalid_ReturnErr(t *testing.T) {
	testCases := []struct {
		desc        string
		password    string
		expectedErr string
	}{
		{
			desc:        "given password with no alphabetic chars then should return error",
			password:    "1234",
			expectedErr: "password must contain alphabetic characters",
		},
		{
			desc:        "given password with no lowercase letters then should return error",
			password:    "ABCD",
			expectedErr: "password must contain lowercase letters",
		},
		{
			desc:        "given password with no uppercase letters then should return error",
			password:    "abcd",
			expectedErr: "password must contain uppercase letters",
		},
		{
			desc:        "given password with no numbers then should return error",
			password:    "abCD",
			expectedErr: "password must contain numbers",
		},
		{
			desc:        "given password with no special characters then should return error",
			password:    "abCD12",
			expectedErr: "password must contain special characters",
		},
		{
			desc:        "given password with non permitted special characters then should return error",
			password:    "abCD12!@'",
			expectedErr: "password can't contain special characters `, ', \", / or \\",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			_, err := auth.Register(auth.RegistrationInfo{
				Email:    "nobody@outlook.com",
				Password: tC.password,
			})

			assert.NotNil(t, err)
			assert.Equal(t, tC.expectedErr, err.Error())
		})
	}
}
