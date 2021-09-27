package repository

import (
	"github.com/rakutenshortintern2021-D-utopia/D-4_2/pkg/domain/model"
)

type (
	// UserRepository is the interface of User repository.
	UserRepository interface {
		GetAll() ([]model.User, error)
	}
)
