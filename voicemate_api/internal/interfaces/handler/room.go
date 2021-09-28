package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/rakutenshortintern2021-D-utopia/D-4_2/pkg/domain/body"
	"github.com/rakutenshortintern2021-D-utopia/D-4_2/pkg/usecase"
)

func GetRoom() echo.HandlerFunc {
	return func(context echo.Context) error {
		idParameter := context.Param("id")
		idParameter_int, error := strconv.ParseUint(idParameter, 10, 64)
		if error != nil {
			return context.JSON(http.StatusBadRequest, nil)
		}

		room := usecase.GetRoom(idParameter_int)
		return context.JSON(http.StatusOK, room)
	}
}

func GetRooms() echo.HandlerFunc {
	return func(context echo.Context) error {
		nameParameter := context.QueryParam("name")

		ageLowerParameter := context.QueryParam("age_lower")
		ageLowerParameter_int := uint32(0)
		if ageLowerParameter != "" {
			tmp, err := strconv.ParseUint(ageLowerParameter, 10, 64)
			if err != nil {
				return context.JSON(http.StatusBadRequest, nil)
			}
			ageLowerParameter_int = uint32(tmp)
		}

		ageUpperParameter := context.QueryParam("age_upper")
		ageUpperParameter_int := uint32(1000)
		if ageUpperParameter != "" {
			tmp, err := strconv.ParseUint(ageUpperParameter, 10, 64)
			if err != nil {
				return context.JSON(http.StatusBadRequest, nil)
			}
			ageUpperParameter_int = uint32(tmp)
		}

		genderParameter := context.QueryParam("gender")

		memberLimitParameter := context.QueryParam("member_limit")
		memberLimitParameter_int := uint32(1000)
		if memberLimitParameter != "" {
			tmp, err := strconv.ParseUint(memberLimitParameter, 10, 64)
			if err != nil {
				return context.JSON(http.StatusBadRequest, nil)
			}
			memberLimitParameter_int = uint32(tmp)
		}

		rooms := usecase.GetRooms(nameParameter, ageLowerParameter_int, ageUpperParameter_int, genderParameter, memberLimitParameter_int)
		return context.JSON(http.StatusOK, rooms)
	}
}

func UpdateRoom() echo.HandlerFunc {
	return func(context echo.Context) error {
		idParameter := context.Param("id")
		id, error := strconv.ParseUint(idParameter, 10, 64)
		if error != nil {
			return context.JSON(http.StatusBadRequest, nil)
		}

		var requestBody body.PutRoomRequest
		if err := context.Bind(&requestBody); err != nil {
			return err
		}

		room := usecase.UpdateRoom(id, requestBody)
		return context.JSON(http.StatusOK, room)
	}
}

func AddRoom() echo.HandlerFunc {
	return func(context echo.Context) error {
		var requestBody body.PutRoomRequest
		if err := context.Bind(&requestBody); err != nil {
			return err
		}

		room := usecase.AddRoom(requestBody)
		return context.JSON(http.StatusOK, room)
	}
}

func DeleteRoom() echo.HandlerFunc {
	return func(context echo.Context) error {
		idParameter := context.Param("id")
		id, err := strconv.ParseUint(idParameter, 10, 64)
		if err != nil {
			return context.JSON(http.StatusBadRequest, nil)
		}

		err = usecase.DeleteRoom(id)
		return context.JSON(http.StatusOK, err)
	}

}
