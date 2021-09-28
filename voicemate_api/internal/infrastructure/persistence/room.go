package persistence

import (
	"github.com/jinzhu/gorm"

	"github.com/rakutenshortintern2021-D-utopia/D-4_2/pkg/domain/model"
	"github.com/rakutenshortintern2021-D-utopia/D-4_2/pkg/domain/repository"
)

type (
	// RoomPersistence is a persistence to preserve rooms.
	RoomPersistence struct {
		Connection *gorm.DB
	}
)

// NewRoomPersistence creates a new Room persistence.
func NewRoomPersistence() repository.RoomRepository {
	return RoomPersistence{
		Connection: getDBConnection(),
	}
}

// FindRoom gets room from DB.
func (roomPersistence RoomPersistence) FindRoomByID(id uint64) (model.Room, error) {
	room := model.Room{}

	result := roomPersistence.Connection.New().
		Table("rooms").
		Where(`"id" = ?`, id).
		Find(&room).
		Related(&room.Tags, "Tags")

	if result.RecordNotFound() {
		return room, nil
	}
	if result.Error != nil {
		return room, result.Error
	}
	return room, nil
}

// GetAll gets rooms from DB.
// func (roomPersistence RoomPersistence) GetAll() ([]model.Room, error) {
// 	rooms := []model.Room{}

// 	result := roomPersistence.Connection.New().
// 		Table("rooms").
// 		Find(&rooms)

// 	if result.RecordNotFound() {
// 		return rooms, nil
// 	}
// 	if result.Error != nil {
// 		return rooms, result.Error
// 	}
// 	return rooms, nil
// }

// DeleteRoom
func (roomPersistence RoomPersistence) DeleteRoom(id uint64) error {
	tag := model.Room{
		ID: id,
	}

	result := RoomPersistence.Connection.New().Delete(&room)
	if result.RecordNotFound() {
		return nil
	}
	if result.Error != nil {
		return result.Error
	}

	return nil
}