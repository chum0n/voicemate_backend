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

// GetRooms gets rooms matched with condition.
func GetRooms(name string, age_lower uint32, age_upper uint32, gender string, member_limit uint32) []model.Room {
	roomPersistence := persistence.NewRoomPersistence()

	rooms, err := roomPersistence.GetRooms(name, age_lower, age_upper, gender, member_limit)
	if err != nil {
		panic(err)
	}

	return rooms
}
