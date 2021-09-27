package model

type (
	Tag struct {
		ID   uint64 `json:"id" gorm:"column:id;type:bigserial;not null;primary_key"`
		Name string `json:"name" gorm:"column:name;type:text;not null"`
	}
)

// TableName returns table name of the model.
func (tag Tag) TableName() string {
	return "tags"
}
