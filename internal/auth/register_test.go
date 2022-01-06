package auth_test

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/codigician/auth/internal/auth"
	"github.com/codigician/auth/internal/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestRegister_RegistrationInfo_ReturnsUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	r := mocks.NewMockRepository(ctrl)
	m := mocks.NewMockMailer(ctrl)
	s := auth.New(r, m)

	ri := auth.RegistrationInfo{
		Firstname: "lacin",
		Lastname:  "bilgin",
		Email:     "lacin@outlook.com",
		Password:  "12345",
	}
	ctx := context.Background()
	body := fmt.Sprintf("Hello, %s %s. This is your verification email.", ri.Firstname, ri.Lastname)

	r.EXPECT().Save(ctx, gomock.Any()).Return(nil).Times(1)
	m.EXPECT().Mail(ri.Email, "Verification", body).Times(1)
	u, _ := s.Register(ctx, &ri)

	assert.NotNil(t, u)
}

func TestRegister_SaveFails_ReturnsError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	r := mocks.NewMockRepository(ctrl)
	m := mocks.NewMockMailer(ctrl)
	s := auth.New(r, m)

	ri := auth.RegistrationInfo{
		Firstname: "lacin",
		Lastname:  "bilgin",
		Email:     "lacin@outlook.com",
		Password:  "12345",
	}
	ctx := context.Background()

	r.EXPECT().Save(ctx, gomock.Any()).Return(errors.New("save failed")).Times(1)
	u, err := s.Register(ctx, &ri)

	assert.Nil(t, u)
	assert.Error(t, err, "save failed")
}
