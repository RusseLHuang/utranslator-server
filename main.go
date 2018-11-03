package main

import (
	"fmt"
	"os"

	"github.com/labstack/echo"
	"github.com/spf13/viper"
	"github.com/utranslator-server/constant"
	"github.com/utranslator-server/controllers"
	"github.com/utranslator-server/database"
	"github.com/utranslator-server/utils"
)

func main() {

	if os.Getenv(constant.ENV) != constant.Production ||
		os.Getenv(constant.ENV) != constant.Development {
		viper.SetConfigName("config")
		viper.SetConfigType("toml")
		viper.AddConfigPath(".")
		viper.ReadInConfig()
	} else {
		viper.AutomaticEnv()
	}

	host := viper.GetString("server")
	fmt.Println("Host : ", host)

	_, err := db.Connect("", "", "zimu", host)

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
