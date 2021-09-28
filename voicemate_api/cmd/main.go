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

	// 部屋
	echoInstance.GET("/rooms/:id", handler.GetRoom())

	// タグ
	echoInstance.GET("/tags/:id", handler.GetTag())
	echoInstance.GET("/tags", handler.GetTags())
	echoInstance.POST("/tags", handler.AddTag())
	echoInstance.DELETE("/tags/:id", handler.DeleteTag())
	echoInstance.DELETE("/room/:id", handler.DeleteRoom())

	echoInstance.Start(":8000")
}
