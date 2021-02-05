package main

import (
	"fmt"
	"log"

	app "github.com/menkap/projects/api"

	helper "github.com/menkap/projects/helper"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	helper.InitViper()

	//Bind API
	Init(e)
	serverPort := helper.GetConfig("serverPort")
	fmt.Println("server started on localhost:", serverPort)
	err := e.Start(":" + serverPort)
	if err != nil {
		log.Fatal(err)
	}
}

func Init(e *echo.Echo) {
	app.Init(e)
}
