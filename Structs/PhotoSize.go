package Structs

type PhotoSize struct {
	FileId       string `json:"file_id"`
	FileUniqueId string `json:"file_unique_id"`
	Width        int    `json:"width"`
	Length       int    `json:"length"`
	FileSize     int    `json:"file_size,omitempty"`
}
