package usecase

import (
	"log"

	"github.com/rakutenshortintern2021-D-utopia/D-4_2/internal/infrastructure/persistence"
	"github.com/rakutenshortintern2021-D-utopia/D-4_2/pkg/domain/body"
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

// func UpdateRoom(id uint64, name string, age_lower_str string, age_upper_str string, gender string, member_limit_str string, introduction string, tag_ids string) (room model.Room) {
func UpdateRoom(id uint64, requestBody body.PutRoomRequest) (room model.Room) {
	roomPersistence := persistence.NewRoomPersistence()

	attributes := make(map[string]interface{})

	attributes["Name"] = requestBody.Name
	attributes["AgeLower"] = requestBody.AgeLower
	attributes["AgeUpper"] = requestBody.AgeUpper
	attributes["Gender"] = requestBody.Gender
	attributes["MemberLimit"] = requestBody.MemberLimit
	attributes["Introduction"] = requestBody.Introduction

	// if name != "" {
	// 	attributes["Name"] = name
	// }

	// if age_lower_str != "" {
	// 	age_lower, error := strconv.ParseUint(age_lower_str, 10, 32)
	// 	if error != nil {
	// 		return room
	// 	}
	// 	attributes["AgeLower"] = age_lower
	// }

	// if age_upper_str != "" {
	// 	age_upper, error := strconv.ParseUint(age_upper_str, 10, 32)
	// 	if error != nil {
	// 		return room
	// 	}
	// 	attributes["AgeUpper"] = age_upper
	// }

	// if gender != "" {
	// 	attributes["Gender"] = gender
	// }

	// if member_limit_str != "" {
	// 	member_limit, error := strconv.ParseUint(member_limit_str, 10, 32)
	// 	if error != nil {
	// 		return room
	// 	}
	// 	attributes["MemberLimit"] = member_limit
	// }

	// if introduction != "" {
	// 	attributes["Introduction"] = introduction
	// }

	room, err := roomPersistence.UpdateRoom(id, attributes)
	if err != nil {
		panic(err)
	}

	log.Print(requestBody.TagIDs)
	err = roomPersistence.SaveTags(id, requestBody.TagIDs)
	if err != nil {
		panic(err)
	}

	return room
}
