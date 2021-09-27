package usecase

import (
	"github.com/rakutenshortintern2021-D-utopia/D-4_2/internal/infrastructure/persistence"
	"github.com/rakutenshortintern2021-D-utopia/D-4_2/pkg/domain/model"
)

// GetUserList gets users matched with condition.
func GetUserList() []model.User {
	userPersistence := persistence.NewUserPersistence()

	users, error := userPersistence.GetAll()

	if error != nil {
		panic(error)
	}
	return users
}
