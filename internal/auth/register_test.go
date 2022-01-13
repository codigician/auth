package auth_test

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/codigician/auth/internal/auth"
	mocks "github.com/codigician/auth/internal/mocks/auth"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestRegister_RegistrationInfo_ReturnsUser(t *testing.T) {
	r := newMockRepository(t)
	m := newMockMailer(t)
	s := auth.New(r, m, nil)

	ri := mockRegistrationInfo()
	ctx := context.Background()
	body := fmt.Sprintf("Hello, %s %s. This is your verification email.",
		ri.Firstname, ri.Lastname)

	r.EXPECT().Save(ctx, gomock.Any()).Return(nil).Times(1)
	m.EXPECT().Mail(ri.Email, "Verification", body).Times(1)
	u, _ := s.Register(ctx, ri)

	assert.NotNil(t, u)
}

func TestRegister_SaveFails_ReturnsError(t *testing.T) {
	r := newMockRepository(t)
	s := auth.New(r, nil, nil)

	ri := mockRegistrationInfo()
	ctx := context.Background()
	r.EXPECT().Save(ctx, gomock.Any()).Return(errors.New("save failed")).Times(1)

	u, err := s.Register(ctx, ri)

	assert.Nil(t, u)
	assert.Error(t, err, "save failed")
}

func newMockRepository(t *testing.T) *mocks.MockRepository {
	return mocks.NewMockRepository(gomock.NewController(t))
}

func newMockMailer(t *testing.T) *mocks.MockMailer {
	return mocks.NewMockMailer(gomock.NewController(t))
}
