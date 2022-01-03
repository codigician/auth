package auth_test

import (
	"errors"
	"testing"

	"github.com/codigician/auth/internal/auth"
	"github.com/stretchr/testify/assert"
)

type mockRepository struct{}

func (m *mockRepository) Save(u *auth.User) error {
	var userInfoValues = []string{u.ID, u.Firstname, u.Lastname, u.Email, u.HashedPassword}
	targetVal := "force-error"
	errMsg := "some error"

	if contains(userInfoValues, targetVal) {
		return errors.New(errMsg)
	}

	return nil
}

func contains(arr []string, val string) bool {
	for _, str := range arr {
		if str == val {
			return true
		}
	}

	return false
}

func TestRegister_InvalidUserInfo_ReturnsNillAndError(t *testing.T) {
	r := auth.NewRegistrator(&mockRepository{})
	registrationInfo := auth.RegistrationInfo{
		Firstname: "Lacin",
		Lastname:  "Bilgin",
		Email:     "nobody@outlook.com",
		Password:  "force-error",
	}

	u, err := r.Register(registrationInfo)

	assert.Nil(t, u)
	assert.NotNil(t, err)
}

func TestRegister_ValidUserInfo_ReturnsUserAndNill(t *testing.T) {
	r := auth.NewRegistrator(&mockRepository{})
	registrationInfo := auth.RegistrationInfo{
		Firstname: "Yuksel",
		Lastname:  "Bilgin",
		Email:     "nobody@outlook.com",
		Password:  "123@clA",
	}
	user := &auth.User{
		ID:             "0",
		Firstname:      registrationInfo.Firstname,
		Lastname:       registrationInfo.Lastname,
		Email:          registrationInfo.Email,
		HashedPassword: registrationInfo.Password,
	}

	u, err := r.Register(registrationInfo)

	assert.Equal(t, user, u)
	assert.Nil(t, err)
}
