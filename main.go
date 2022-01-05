package main

import (
	"github.com/codigician/auth/internal/auth"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load(".env")

	auth.Run()
}
