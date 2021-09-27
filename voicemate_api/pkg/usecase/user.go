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
