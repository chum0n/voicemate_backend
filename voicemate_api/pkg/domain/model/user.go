package model

type (
	// User is a user on voicemate.
	User struct {
		ID       uint64  `json:"id" gorm:"column:id;type:bigserial;not null;primary_key"`
		Name     string  `json:"name" gorm:"column:name;type:text;not null"`
		Email    string  `json:"email" gorm:"column:email;type:text;not null"`
		Password string  `json:"password" gorm:"column:password;type:text;not null"`
		Age      *uint32 `json:"age" gorm:"column:age;type:integer"`
		Gender   *string `json:"gender" gorm:"column:gender;type:text"`
		RoomID   *uint64 `json:"roomID" gorm:"column:room_id;type:bigint"`
	}
)

// TableName returns table name of the model.
func (user User) TableName() string {
	return "users"
}
