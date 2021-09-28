package body

type (
	PutRoomRequest struct {
		Name         string   `json:"name"`
		AgeLower     *uint32  `json:"ageLower"`
		AgeUpper     *uint32  `json:"ageUpper"`
		Gender       *string  `json:"gender"`
		MemberLimit  *uint32  `json:"memberLimit"`
		Introduction *string  `json:"introduction"`
		Member       uint32   `json:"member"`
		TagIDs       []uint64 `json:"tagIDs"`
	}

	PutUserRequest struct {
		Name     string   `json:"name"`
		Email    string   `json:"email"`
		Password string   `json:"password"`
		RoomID   *uint64  `json:"roomID"`
		TagIDs   []uint64 `json:"tagIDs"`
	}
)
