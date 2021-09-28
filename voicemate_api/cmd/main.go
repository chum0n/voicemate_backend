package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/rakutenshortintern2021-D-utopia/D-4_2/internal/interfaces/handler"
)

func main() {
	echoInstance := echo.New()
	// CORS
	echoInstance.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))
	echoInstance.Use(middleware.Logger())
	echoInstance.Use(middleware.Recover())

	// ユーザ
	echoInstance.GET("/users/:id", handler.GetUser())
	echoInstance.PUT("/users/:id", handler.UpdateUser())

	// 部屋
	echoInstance.GET("/rooms/:id", handler.GetRoom())
	echoInstance.GET("/rooms", handler.GetRooms())
	echoInstance.PUT("/rooms/:id", handler.UpdateRoom())
	echoInstance.POST("/rooms", handler.AddRoom())

	// タグ
	echoInstance.GET("/tags/:id", handler.GetTag())
	echoInstance.GET("/tags", handler.GetTags())
	echoInstance.POST("/tags", handler.AddTag())
	echoInstance.DELETE("/tags/:id", handler.DeleteTag())

	echoInstance.Start(":8000")
}
