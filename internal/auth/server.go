package auth

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Run() {
	e := echo.New()
	e.GET("/", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "Hello hun")
	})
	e.Start(":1323")

	// e.GET("/users", getAllUsers)

	e.GET("/users/", func(ctx echo.Context) error {
		fmt.Println("request recognized")
		mongo := Mongo{}
		users, err := mongo.GetAll()
		fmt.Println(users)
		if err != nil {
			return err
		}
		return ctx.JSON(http.StatusOK, users)
	})

}

func getUser(ctx echo.Context) error {
	// mongo := Mongo{}
	// user, err := mongo.Get("selmin@outlook.com")
	// if err != nil {
	// 	return err
	// }
	return ctx.String(http.StatusOK, "user")
}
