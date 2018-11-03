package main

import (
	"github.com/labstack/echo"
	"github.com/utranslator-server/controllers"
	"github.com/utranslator-server/database"
	"github.com/utranslator-server/utils"
)

func main() {
	_, err := db.Connect("", "", "utranslator", "127.0.0.1")

	if err != nil {
		panic(err)
	}

	credential.Init()

	e := echo.New()

	// Testing purpose
	e.GET("/", credential.HandleMain())
	e.GET("/login", credential.Login())
	e.GET("/callback", credential.Authorize())

	//
	e.GET("/member/:id", controllers.GetMember())
	e.POST("/member/google_token", controllers.CreateGoogleToken())

	e.Logger.Fatal(e.Start(":7777"))
}
