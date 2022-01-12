package main

import (
	"fmt"
	"log"

	"github.com/codigician/auth/internal/auth"
	authmongo "github.com/codigician/auth/internal/auth/mongo"
	"github.com/codigician/auth/internal/handler"
	"github.com/codigician/auth/internal/token"
	"github.com/codigician/auth/pkg/mongo"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}

	m, err := mongo.New(&mongo.Config{URI: "mongodb://localhost:27017"})
	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()
	tokenService := token.New(token.NewTokenCreator())
	authRepository := authmongo.New(m.Collection("auth", "users"))
	authService := auth.New(authRepository, nil, tokenService)
	authHandler := handler.NewAuth(authService)

	authHandler.RegisterRoutes(e)
	// log.Fatal(e.Start(":8888"))

	tokens, err := authService.Authorize(auth.NewUser(&auth.RegistrationInfo{
		Firstname: "laco",
		Lastname:  "bilgo",
		Email:     "lolo@outlook.com",
		Password:  "12345",
	}))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("access token:", tokens.AT, "\nrefresh token:", tokens.RT)
}
