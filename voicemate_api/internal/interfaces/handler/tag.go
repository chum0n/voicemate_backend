package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"

	"github.com/rakutenshortintern2021-D-utopia/D-4_2/pkg/usecase"
)

func GetTag() echo.HandlerFunc {
	return func(context echo.Context) error {
		idParameter := context.Param("id")
		id, err := strconv.ParseUint(idParameter, 10, 64)
		if err != nil {
			return context.JSON(http.StatusBadRequest, nil)
		}

		tag := usecase.GetTag(id)
		return context.JSON(http.StatusOK, tag)
	}
}

func GetTags() echo.HandlerFunc {
	return func(context echo.Context) error {
		tags := usecase.GetTags()
		return context.JSON(http.StatusOK, tags)
	}
}

func AddTag() echo.HandlerFunc {
	return func(context echo.Context) error {
		nameParameter := context.FormValue("name")

		tag := usecase.AddTag(nameParameter)
		return context.JSON(http.StatusOK, tag)
	}
}
