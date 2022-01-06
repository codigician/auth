package main

import (
	"log"

	"github.com/codigician/auth/internal/auth"
	authmongo "github.com/codigician/auth/internal/auth/mongo"
	"github.com/codigician/auth/internal/handler"
	"github.com/codigician/auth/internal/mail"
	"github.com/codigician/auth/pkg/mongo"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

const _appPsw = "xoycpttqujgpgepjc"

func main() {
	godotenv.Load(".env")

	m, err := mongo.New(&mongo.Config{URI: "mongodb://localhost:27017"})
	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()
	mailer := mail.New()

	authRepository := authmongo.New(m.Collection("auth", "users"))
	authService := auth.New(authRepository, mailer)
	authHandler := handler.NewAuth(authService)

	authHandler.RegisterRoutes(e)
	log.Fatal(e.Start(":8888"))
}
