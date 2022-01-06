package auth_test

import (
	"context"
	"errors"
	"testing"

	"github.com/codigician/auth/internal/auth"
	"github.com/codigician/auth/internal/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestAuthenticate_RightCredentials_ReturnsNil(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	r := mocks.NewMockRepository(ctrl)
	m := mocks.NewMockMailer(ctrl)
	s := auth.New(r, m)

	ui := auth.UserCredentials{
		Email:    "lacin@outlook.com",
		Password: "12345",
	}
	ctx := context.Background()

	r.EXPECT().Get(ctx, ui.Email).Return(&auth.User{
		Firstname:      "lacin",
		Lastname:       "bilgin",
		Email:          "lacin@outlook.com",
		HashedPassword: auth.HashPassword(ui.Password),
	}, nil).Times(1)

	err := s.Authenticate(ctx, &ui)

	assert.Nil(t, err)
}

func TestAuthenticate_WrongCredentials_ReturnsError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	r := mocks.NewMockRepository(ctrl)
	m := mocks.NewMockMailer(ctrl)
	s := auth.New(r, m)

	ui := auth.UserCredentials{
		Email:    "lacin@outlook.com",
		Password: "12345",
	}
	ctx := context.Background()

	r.EXPECT().Get(ctx, ui.Email).Return(&auth.User{
		Firstname:      "lacin",
		Lastname:       "bilgin",
		Email:          "lacin@outlook.com",
		HashedPassword: auth.HashPassword("2346"),
	}, errors.New("wrong credentials")).Times(1)

	err := s.Authenticate(ctx, &ui)

	assert.Error(t, err, "wrong credentials")
}
