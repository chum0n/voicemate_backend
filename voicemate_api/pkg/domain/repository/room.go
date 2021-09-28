package repository

import (
	"github.com/rakutenshortintern2021-D-utopia/D-4_2/pkg/domain/model"
)

type (
	// RoomRepository is the interface of Room repository.
	RoomRepository interface {
		FindRoomByID(id uint64) (model.Room, error)
		// GetAll() ([]model.Room, error)
		DeleteRoom(id uint64) error
	}
)
