package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/rakutenshortintern2021-D-utopia/D-4_2/pkg/usecase"
)

func GetRoom() echo.HandlerFunc {
	return func(context echo.Context) error {
		idParameter := context.Param("id")
		id, error := strconv.ParseUint(idParameter, 10, 64)
		if error != nil {
			return context.JSON(http.StatusBadRequest, nil)
		}

		room := usecase.GetRoom(id)
		return context.JSON(http.StatusOK, room)
	}
}

func DeleteRoom() echo.HandlerFunc {
	return func(context echo.Context) error{
		idParameter := context.Param("id")
		id, err := strconv.ParseUint(idParameter, 10, 64)
		if err != nil {
			return context.JSON(http.StatusBadRequest, nil)
		}

		err = usecase.DeleteRoom(id)
		return context.JSON(http.StatusOK, err)
	}

}