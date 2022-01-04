package main

import (
	"fmt"
	"log"

	"github.com/codigician/auth/internal/auth"
)

// type Repo struct{}

// func (r Repo) Save(u *auth.User) error {
// 	fmt.Println("save function ran")
// 	return nil
// }

// func (r Repo) Find(e string) (*auth.User, error) {
// 	fmt.Println("find function ran")
// 	return nil, nil
// }

func main() {
	// fmt.Println("hello, world")

	// id := uuid.Generate()

	// u := auth.User{
	// 	ID:             id,
	// 	Firstname:      "lacin",
	// 	Lastname:       "bilgin",
	// 	HashedPassword: "124f",
	// }

	mongo := auth.Mongo{}
	registrator := auth.NewRegistrator(mongo)
	ri := auth.RegistrationInfo{
		Firstname: "fatma",
		Lastname:  "candir",
		Email:     "someotherbody@outlook.com",
		Password:  "4567",
	}
	u, err := registrator.Register(ri)
	if err != nil {
		log.Fatal(err)
	}

	u2, err := mongo.Find(u.Email)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(u2)
	fmt.Println(u)
}
