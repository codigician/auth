package handler

import (
	"context"
	"fmt"
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
	e.GET("/users", a.Authenticate)
	e.GET("/users/forgot-password/:email", a.ForgotPassword)
}

func (a *Auth) Authenticate(c echo.Context) error {
	var req AuthenticateRequest
	if err := c.Bind(&req); err != nil {
		return echo.ErrBadRequest
	}

	err := a.service.Authenticate(c.Request().Context(), req.To())
	if err != nil {
		fmt.Println(err)
		return echo.ErrUnauthorized
	}

	return c.JSON(http.StatusOK, err)
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

	return c.JSON(http.StatusCreated, FromUser(u))
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

	AuthenticateRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
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

func (ar AuthenticateRequest) To() *auth.UserCredentials {
	return &auth.UserCredentials{
		Email:    ar.Email,
		Password: ar.Password,
	}
}

func FromUser(u *auth.User) *RegisterResponse {
	return &RegisterResponse{
		Email:     u.Email,
		Firstname: u.Firstname,
		Lastname:  u.Lastname,
	}
}
