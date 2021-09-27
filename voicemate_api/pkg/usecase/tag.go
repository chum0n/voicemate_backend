package usecase

import (
	"github.com/rakutenshortintern2021-D-utopia/D-4_2/internal/infrastructure/persistence"
	"github.com/rakutenshortintern2021-D-utopia/D-4_2/pkg/domain/model"
)

// GetTag gets tags.
func GetTag(id uint64) model.Tag {
	tagPersistence := persistence.NewTagPersistence()

	tag, error := tagPersistence.FindTagByID(id)

	if error != nil {
		panic(error)
	}
	return tag
}

// GetTags gets tags.
func GetTags() []model.Tag {
	tagPersistence := persistence.NewTagPersistence()

	tags, error := tagPersistence.GetAll()

	if error != nil {
		panic(error)
	}
	return tags
}

// AddTag
func AddTag(name string) model.Tag {
	tagPersistence := persistence.NewTagPersistence()

	tag, error := tagPersistence.CreateTag(name)

	if error != nil {
		panic(error)
	}
	return tag
}

// DeleteTag
func DeleteTag(id uint64) error {
	tagPersistence := persistence.NewTagPersistence()

	err := tagPersistence.DeleteTag(id)
	if err != nil {
		panic(err)
	}

	return err
}
