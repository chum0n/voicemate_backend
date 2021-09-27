package body

type putRoomRequest struct {
	Name         string   `json:"name"`
	AgeLower     uint32   `json:"ageLower"`
	AgeUpper     uint32   `json:"ageUpper"`
	Gender       string   `json:"gender"`
	MemberLimit  uint32   `json:"memberLimit"`
	Introduction string   `json:"introduction"`
	TagIDs       []uint64 `json:"tagIDs"`
}
