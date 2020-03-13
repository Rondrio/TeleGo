package Structs

type VideoNote struct {
	FileId       string    `json:"file_id"`
	FileUniqueId string    `json:"file_unique_id"`
	Length       int       `json:"length"`
	Duration     int       `json:"duration"`
	Thumb        PhotoSize `json:"thumb,omitempty"`
	FileSize     int       `json:"file_size,omitempty"`
}
