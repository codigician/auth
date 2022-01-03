package auth_test

import (
	"errors"
	"testing"

	"github.com/codigician/auth/internal/auth"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

type mockRepository struct {
	saveReturn error
}

func (m *mockRepository) Save(u *auth.User) error {
	return m.saveReturn
}

func TestRegister_SaveFailed_ReturnsNillAndError(t *testing.T) {
	r := auth.NewRegistrator(&mockRepository{saveReturn: errors.New("some error")})

	u, err := r.Register(auth.RegistrationInfo{})

	assert.Nil(t, u)
	assert.NotNil(t, err)
}

func TestRegister_SaveSuccessful_ReturnsUserAndNil(t *testing.T) {
	r := auth.NewRegistrator(&mockRepository{saveReturn: nil})

	u, err := r.Register(auth.RegistrationInfo{})

	assert.NotNil(t, u)
	assert.Nil(t, err)
}

func TestRegister_Password_SuccessfullyHashesPassword(t *testing.T) {
	r := auth.NewRegistrator(&mockRepository{saveReturn: nil})

	ri := auth.RegistrationInfo{
		Firstname: "lacin",
		Lastname:  "bilgin",
		Email:     "nobody@outlook.com",
		Password:  "12345",
	}

	u, _ := r.Register(ri)

	psw, _ := auth.HashPassword(ri.Password)
	err := bcrypt.CompareHashAndPassword([]byte(psw), []byte(ri.Password))

	assert.Nil(t, err)
	assert.NotNil(t, u)
}
