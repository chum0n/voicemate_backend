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
		Preload("Tags").
		Find(&room)

	if result.RecordNotFound() {
		return room, nil
	}
	if result.Error != nil {
		return room, result.Error
	}
	return room, nil
}

// GetAll gets rooms from DB.
func (roomPersistence RoomPersistence) GetRooms(name string, age_lower uint32, age_upper uint32, gender string, member_limit uint32) ([]model.Room, error) {
	rooms := []model.Room{}

	query := roomPersistence.Connection.New().
		Where("age_lower >= ? AND age_upper <= ? AND member_limit <= ?", age_lower, age_upper, member_limit)
	if name != "" {
		query = query.Where("name LIKE ?", "%"+name+"%")
	}
	if gender != "" {
		query = query.Where("gender = ?", gender)
	}
	result := query.Preload("Tags").Find(&rooms)

	if result.RecordNotFound() {
		return rooms, nil
	}
	if result.Error != nil {
		return rooms, result.Error
	}

	return rooms, nil
}
