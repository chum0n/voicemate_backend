package repository

import (
	"github.com/rakutenshortintern2021-D-utopia/D-4_2/pkg/domain/body"
	"github.com/rakutenshortintern2021-D-utopia/D-4_2/pkg/domain/model"
)

type (
	// RoomRepository is the interface of Room repository.
	RoomRepository interface {
		FindRoomByID(id uint64) (model.Room, error)
		GetRooms(name string, age_lower uint32, age_upper uint32, gender string, member_limit uint32) ([]model.Room, error)
		UpdateRoom(uint64, map[string]interface{}) (model.Room, error)
		SaveTags(uint64, []uint64) error
		CreateRoom(body.PutRoomRequest) (model.Room, error)
		DeleteRoom(id uint64) error
	}
)
