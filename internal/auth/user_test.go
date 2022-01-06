package auth_test

import (
	"testing"

	"github.com/codigician/auth/internal/auth"
	"github.com/stretchr/testify/assert"
)

func TestNewUser_RegistrationInfo_ReturnsUser(t *testing.T) {
	ri := mockRegistrationInfo()
	u := auth.NewUser(ri)
	assert.Equal(t, ri.Firstname, u.Firstname)
	assert.Equal(t, ri.Lastname, u.Lastname)
	assert.Equal(t, ri.Email, u.Email)
	assert.NotEqual(t, ri.Password, u.HashedPassword)
}

func TestComparePassword_RightRawPassword_ReturnsNil(t *testing.T) {
	ri := mockRegistrationInfo()
	u := auth.NewUser(ri)
	err := u.ComparePassword(ri.Password)
	assert.Nil(t, err)
}

func TestComparePassword_WrongRawPassword_ReturnsError(t *testing.T) {
	ri := mockRegistrationInfo()
	u := auth.NewUser(ri)
	err := u.ComparePassword("somepassword")
	assert.Error(t, err)
}

func mockRegistrationInfo() *auth.RegistrationInfo {
	return &auth.RegistrationInfo{
		Firstname: "lacin",
		Lastname:  "bilgin",
		Email:     "lacin@outlook.com",
		Password:  "12345",
	}
}
