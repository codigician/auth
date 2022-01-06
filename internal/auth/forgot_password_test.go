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

func TestForgotPassword_RightEmail_ReturnsNil(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	r := mocks.NewMockRepository(ctrl)
	m := mocks.NewMockMailer(ctrl)
	s := auth.New(r, m)

	ctx := context.Background()
	ri := auth.RegistrationInfo{
		Firstname: "lacin",
		Lastname:  "bilgin",
		Email:     "lacin@outlook.com",
		Password:  "1234",
	}
	link := fmt.Sprintf("http://localhost:8888/password-reset/")
	body := fmt.Sprintf("Click the link below to reset your password\nPassword reset link: %s", link)

	r.EXPECT().Get(ctx, ri.Email).Return(&auth.User{
		Firstname:      ri.Firstname,
		Lastname:       ri.Lastname,
		Email:          ri.Email,
		HashedPassword: auth.HashPassword(ri.Password),
	}, nil).Times(1)
	m.EXPECT().Mail(ri.Email, "Password Reset", body).Return(nil).Times(1)

	err := s.ForgotPassword(ctx, ri.Email)

	assert.Nil(t, err)
}

func TestForgotPassword_WrongEmail_ReturnsError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	r := mocks.NewMockRepository(ctrl)
	m := mocks.NewMockMailer(ctrl)
	s := auth.New(r, m)

	ctx := context.Background()
	wrongEmail := "lamcin@outlk.coom"

	r.EXPECT().Get(ctx, wrongEmail).Return(nil, errors.New("wrong email")).Times(1)

	err := s.ForgotPassword(ctx, wrongEmail)

	assert.Error(t, err, "wrong email")
}
