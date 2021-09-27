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

// FindUserByID find a user by ID.
func (userPersistence UserPersistence) FindUserByID(id uint64) (model.User, error) {
	user := model.User{}

	result := userPersistence.Connection.New().
		Table("users").
		Where(`"id" = ?`, id).
		Find(&user).
		Related(&user.Tags, "Tags")

	if result.RecordNotFound() {
		return user, nil
	}
	if result.Error != nil {
		return user, result.Error
	}
	return user, nil
}

func (userPersistence UserPersistence) UpdateUser(id uint64, attributes map[string]interface{}) (model.User, error) {
	user := model.User{ID: id}

	result := userPersistence.Connection.New().
		Model(&user).
		Updates(attributes)

	if result.RecordNotFound() {
		return user, nil
	}
	if result.Error != nil {
		return user, result.Error
	}
	return user, nil
}
