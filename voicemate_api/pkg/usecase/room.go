package usecase

import (
	"github.com/rakutenshortintern2021-D-utopia/D-4_2/internal/infrastructure/persistence"
	"github.com/rakutenshortintern2021-D-utopia/D-4_2/pkg/domain/model"
)

// GetRoom gets room matched with condition.
func GetRoom(id uint64) model.Room {
	roomPersistence := persistence.NewRoomPersistence()

	room, err := roomPersistence.FindRoomByID(id)
	if err != nil {
		panic(err)
	}

	return room
}

// DeleteRoom
func DeleteRoom(id uint64) error{
	roomPersistence := persistence.NewRoomPersistence()

	err := roomPersistence.DeleteRoom(id)
	if err != nil {
		panic(err)
	}

	return err
}