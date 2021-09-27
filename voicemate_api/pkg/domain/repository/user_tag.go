package repository

import (
	"github.com/rakutenshortintern2021-D-utopia/D-4_2/pkg/domain/model"
)

type (
	// UserTagRepository is the interface of UserTag repository.
	UserTagRepository interface {
		CreateUserTag(model.UserTag) (model.UserTag, error)
		DeleteUserTag(uint64) error
	}
)
