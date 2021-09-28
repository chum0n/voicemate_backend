package persistence

import (
	"github.com/jinzhu/gorm"

	"github.com/rakutenshortintern2021-D-utopia/D-4_2/pkg/domain/body"
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

func (roomPersistence RoomPersistence) UpdateRoom(id uint64, attributes map[string]interface{}) (model.Room, error) {
	room := model.Room{
		ID: id,
	}

	result := roomPersistence.Connection.New().
		Model(&room).
		Updates(attributes)

	if result.RecordNotFound() {
		return room, nil
	}
	if result.Error != nil {
		return room, result.Error
	}
	return room, nil
}

func (roomPersistence RoomPersistence) SaveTags(id uint64, tagIDs []uint64) error {
	room, error := roomPersistence.FindRoomByID(id)
	if error != nil {
		return error
	}

	for _, tag := range room.Tags {
		result := roomPersistence.Connection.New().
			Table("room_tags").
			Model(&room).
			Association("Tags").
			Delete(&tag)

		if result.Error != nil {
			return result.Error
		}
	}

	for _, tagID := range tagIDs {
		tag := model.Tag{ID: tagID}

		result := roomPersistence.Connection.New().
			Table("room_tags").
			Model(&room).
			Association("Tags").
			Append(&tag)

		if result.Error != nil {
			return result.Error
		}
	}

	return nil
}

func (roomPersistence RoomPersistence) CreateRoom(requestBody body.PutRoomRequest) (model.Room, error) {
	room := model.Room{
		Name:         requestBody.Name,
		AgeLower:     requestBody.AgeLower,
		AgeUpper:     requestBody.AgeUpper,
		Gender:       requestBody.Gender,
		MemberLimit:  requestBody.MemberLimit,
		Introduction: requestBody.Introduction,
		Member:       requestBody.Member,
	}

	result := roomPersistence.Connection.New().Create(&room)
	if result.RecordNotFound() {
		return room, nil
	}
	if result.Error != nil {
		return room, result.Error
	}

	return room, nil
}

// DeleteRoom
func (roomPersistence RoomPersistence) DeleteRoom(id uint64) error {
	room := model.Room{
		ID: id,
	}
	result := roomPersistence.Connection.New().Delete(&room)
	if result.RecordNotFound() {
		return nil
	}
	if result.Error != nil {
		return result.Error
	}

	return nil
}
