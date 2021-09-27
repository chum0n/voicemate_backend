package model

type (
	Room struct {
		ID           uint64  `json:"id" gorm:"column:id;type:bigserial;not null;primary_key"`
		Name         string  `json:"name" gorm:"column:name;type:text;not null"`
		AgeLower     *uint32 `json:"ageLower" gorm:"column:age_lower;type:integer;"`
		AgeUpper     *uint32 `json:"ageUpper" gorm:"column:age_upper;type:integer;"`
		Gender       *string `json:"gender" gorm:"column:gender;type:text"`
		MemberLimit  *uint32 `json:"memberLimit" gorm:"column:member_limit;type:integer;"`
		Introduction *string `json:"introduction" gorm:"column:introduction;type:text;"`

		Tags []Tag `json:"tags" gorm:"many2many:user_tags;"`
	}
)

// TableName returns table name of the model.
func (room Room) TableName() string {
	return "rooms"
}
