package repository

import (
	"github.com/rakutenshortintern2021-D-utopia/D-4_2/pkg/domain/model"
)

type (
	// TagRepository is the interface of Tag repository.
	TagRepository interface {
		FindTagByID(id uint64) (model.Tag, error)
		GetAll() ([]model.Tag, error)
		CreateTag(name string) (model.Tag, error)
		DeleteTag(id uint64) error
	}
)
