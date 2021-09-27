package handler

import (
	"net/http"

	"github.com/labstack/echo"

	"github.com/rakutenshortintern2021-D-utopia/D-4_2/pkg/usecase"
)

func GetTag(id uint64) echo.HandlerFunc {
	return func(context echo.Context) error {
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
