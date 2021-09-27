package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"

	"github.com/rakutenshortintern2021-D-utopia/D-4_2/pkg/usecase"
)

func GetUser() echo.HandlerFunc {
	return func(context echo.Context) error {
		idParameter := context.Param("id")
		id, error := strconv.ParseUint(idParameter, 10, 64)
		if error != nil {
			return context.JSON(http.StatusBadRequest, nil)
		}

		user := usecase.GetUser(id)
		return context.JSON(http.StatusOK, user)
	}
}

func UpdateUser() echo.HandlerFunc {
	return func(context echo.Context) error {
		idParameter := context.Param("id")
		id, error := strconv.ParseUint(idParameter, 10, 0)
		if error != nil {
			return context.JSON(http.StatusBadRequest, nil)
		}

		nameParameter := context.FormValue("name")
		emailParameter := context.FormValue("email")
		passwordParameter := context.FormValue("password")

		user := usecase.UpdateUser(id, nameParameter, emailParameter, passwordParameter)
		return context.JSON(http.StatusOK, user)
	}
}

func AddUser() echo.HandlerFunc {
	return func(context echo.Context) error {
		nameParameter := context.FormValue("name")
		emailParameter := context.FormValue("email")
		passwordParameter := context.FormValue("password")

		user := usecase.AddUser(nameParameter, emailParameter, passwordParameter)
		return context.JSON(http.StatusOK, user)
	}
}
