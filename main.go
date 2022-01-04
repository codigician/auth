package main

import (
	"fmt"
	"log"

	"github.com/codigician/auth/internal/auth"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load(".env")
	// mongo := auth.Mongo{}
	// registrator := auth.NewRegistrator(mongo)
	// ri := auth.RegistrationInfo{
	// 	Firstname: "fatma",
	// 	Lastname:  "candir",
	// 	Email:     "someotherbody@outlook.com",
	// 	Password:  "4567",
	// }
	// u, err := registrator.Register(ri)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// u2, err := mongo.Find(u.Email)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(u2)
	// fmt.Println(u)

	uc := auth.UserCredentials{
		Email:    "someotherbody@outlook.com",
		Password: "4567",
	}

	mongo := auth.Mongo{}
	login := auth.NewLogin(mongo)
	fmt.Println("moving on")
	// err := login.Authenticate(&uc)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(err)
	if err := login.NewPassword(uc.Email); err != nil {
		log.Fatal(err)
	}

}
