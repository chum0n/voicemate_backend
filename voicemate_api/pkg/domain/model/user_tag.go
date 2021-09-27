package model

type (
	UserTag struct {
		UserID uint64 `json:"userID" gorm:"column:user_id;type:bigint;not null"`
		TagID  uint64 `json:"tagID" gorm:"column:tag_id;type:bigint;not null"`
	}
)

// TableName returns table name of the model.
func (userTag UserTag) TableName() string {
	return "user_tags"
}
