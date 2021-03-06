package main

import (
	"log"

	"github.com/codigician/auth/internal/auth"
	authmongo "github.com/codigician/auth/internal/auth/mongo"
	"github.com/codigician/auth/internal/handler"
	"github.com/codigician/auth/internal/token"
	tokenmongo "github.com/codigician/auth/internal/token/mongo"
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

	authRepository := authmongo.New(m.Collection("auth", "users"))

	tokenRepository := tokenmongo.New(m.Collection("token", "tokens"))
	tokenCreator := token.NewCreator(token.Config{})
	if err != nil {
		panic(err)
	}

	tokenService := token.New(tokenCreator, tokenRepository)
	authService := auth.New(authRepository, nil, tokenService)

	authHandler := handler.NewAuth(authService)

	e := echo.New()
	authHandler.RegisterRoutes(e)
	log.Fatal(e.Start(":8888"))
}
