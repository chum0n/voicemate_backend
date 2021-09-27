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

	// ルーティング
	echoInstance.GET("/users", handler.GetUserList())

	// サーバー起動、ポート番号の指定
	echoInstance.Start(":8000")
}
