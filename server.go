package main

import (
	"goApi/controllers"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	e.GET("/", api.Get)

	e.Logger.Fatal(e.Start(":5000"))
}
