package handler

import (
	"net/http"

	"github.com/labstack/echo"

	"github.com/rakutenshortintern2021-D-utopia/D-4_2/pkg/usecase"
)

func GetUserList() echo.HandlerFunc {
	return func(context echo.Context) error {
		users := usecase.GetUserList()
		return context.JSON(http.StatusOK, users)
	}
}
