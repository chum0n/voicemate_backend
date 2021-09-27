package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/rakutenshortintern2021-D-utopia/D-4_2/internal/interfaces/handler"
)

func main() {
	echoInstance := echo.New()

	echoInstance.Use(middleware.Logger())
	echoInstance.Use(middleware.Recover())

	echoInstance.GET("/users", handler.GetUserList())

	// タグ
	echoInstance.GET("/tags/:id", handler.GetTag())
	echoInstance.GET("/tag", handler.GetTags())

	echoInstance.Start(":8000")
}
