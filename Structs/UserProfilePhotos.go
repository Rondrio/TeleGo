package Structs

type UserProfilePhotos struct {
	TotalCount int         `json:"total_count"`
	Photos     []PhotoSize `json:"photos"`
}
