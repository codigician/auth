package handler

import (
	"context"
	"net/http"

	"github.com/codigician/auth/internal/auth"
	"github.com/labstack/echo/v4"
)

type (
	AuthService interface {
		Authenticate(ctx context.Context, creds *auth.UserCredentials) error
		ForgotPassword(ctx context.Context, email string) error
		Register(ctx context.Context, info *auth.RegistrationInfo) (*auth.User, error)
	}

	Auth struct {
		service AuthService
	}
)

func NewAuth(service AuthService) *Auth {
	return &Auth{service}
}

func (a *Auth) RegisterRoutes(e *echo.Echo) {
	e.POST("/users", a.Register)
	e.GET("/users/forgot-password/:email", a.ForgotPassword)
}

func (a *Auth) Register(c echo.Context) error {
	var req RegisterRequest
	if err := c.Bind(&req); err != nil {
		return echo.ErrBadRequest
	}

	u, err := a.service.Register(c.Request().Context(), req.To())
	if err != nil {
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusCreated, FromUser(*u))
}

// TODO: ForgotPassword
func (a *Auth) ForgotPassword(c echo.Context) error {
	email := c.Param("email")
	return a.service.ForgotPassword(c.Request().Context(), email)
}

type (
	RegisterRequest struct {
		Email     string `json:"email"`
		Firstname string `json:"firstname"`
		Lastname  string `json:"lastname"`
		Password  string `json:"password"`
	}

	RegisterResponse struct {
		Email     string `json:"email"`
		Firstname string `json:"firstname"`
		Lastname  string `json:"lastname"`
	}
)

func (rr RegisterRequest) To() *auth.RegistrationInfo {
	return &auth.RegistrationInfo{
		Firstname: rr.Firstname,
		Lastname:  rr.Lastname,
		Email:     rr.Email,
		Password:  rr.Password,
	}
}

func FromUser(u auth.User) *RegisterResponse {
	return &RegisterResponse{
		Email:     u.Email,
		Firstname: u.Firstname,
		Lastname:  u.Lastname,
	}
}
