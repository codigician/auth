package main

import (
	"github.com/codigician/auth/internal/auth"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load(".env")

	mongo := auth.Mongo{}
	user, _ := mongo.Find("lor@outlook.com")
	// mongo.Delete(user.ID)
	mongo.Update(user.ID, "lastname", "mukkemmel")
	mongo.List()
	// user, _ := mongo.Find(string(u.Email))
	// mongo.DeleteAll()
	// mongo.List()
}
