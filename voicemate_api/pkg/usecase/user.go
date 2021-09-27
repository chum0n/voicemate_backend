package usecase

import (
	"github.com/rakutenshortintern2021-D-utopia/D-4_2/internal/infrastructure/persistence"
	"github.com/rakutenshortintern2021-D-utopia/D-4_2/pkg/domain/model"
)

// GetUserList get a user by ID.
func GetUser(id uint64) model.User {
	userPersistence := persistence.NewUserPersistence()

	user, error := userPersistence.FindUserByID(id)

	if error != nil {
		panic(error)
	}
	return user
}

func UpdateUser(id uint64, name string, email string, password string) model.User {
	userPersistence := persistence.NewUserPersistence()

	attributes := make(map[string]interface{})

	if name != "" {
		attributes["Name"] = name
	}

	if email != "" {
		attributes["Email"] = name
	}

	if password != "" {
		attributes["Password"] = password
	}

	user, error := userPersistence.UpdateUser(id, attributes)

	if error != nil {
		panic(error)
	}
	return user
}
