package repository

import (
	"github.com/rakutenshortintern2021-D-utopia/D-4_2/pkg/domain/model"
)

type (
	// UserRepository is the interface of User repository.
	UserRepository interface {
		FindUserByID(uint64) (model.User, error)
		CreateUser(model.User) (model.User, error)
		UpdateUser(uint64, map[string]interface{}) (model.User, error)
		SaveTags(uint64, []uint64) error
	}
)
