package persistence

import (
	"github.com/jinzhu/gorm"

	"github.com/rakutenshortintern2021-D-utopia/D-4_2/pkg/domain/model"
	"github.com/rakutenshortintern2021-D-utopia/D-4_2/pkg/domain/repository"
)

type (
	// UserPersistence is a persistence to preserve users.
	UserPersistence struct {
		Connection *gorm.DB
	}
)

// NewUserPersistence creates a new User persistence.
func NewUserPersistence() repository.UserRepository {
	return UserPersistence{
		Connection: getDBConnection(),
	}
}

// GetAll gets users from DB.
func (userPersistence UserPersistence) GetAll() ([]model.User, error) {
	users := []model.User{}

	result := userPersistence.Connection.New().
		Table("users").
		Find(&users)

	if result.RecordNotFound() {
		return users, nil
	}
	if result.Error != nil {
		return users, result.Error
	}
	return users, nil
}
