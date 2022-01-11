package main

import (
	"fmt"
	"log"

	"github.com/codigician/auth/internal/auth"
	authmongo "github.com/codigician/auth/internal/auth/mongo"
	"github.com/codigician/auth/internal/handler"
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
	authRepository := authmongo.New(m.Collection("auth", "users"))
	authService := auth.New(authRepository, nil)
	authHandler := handler.NewAuth(authService)

	authHandler.RegisterRoutes(e)
	log.Fatal(e.Start(":8888"))

	at, rt, err := authService.Authorize(auth.NewUser(&auth.RegistrationInfo{
		Firstname: "laco",
		Lastname:  "bilgo",
		Email:     "lolo@outlook.com",
		Password:  "12345",
	}))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("access token: %s\nrefresh token: %s", at, rt)
}
