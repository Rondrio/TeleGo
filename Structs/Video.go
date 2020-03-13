package Structs

type Video struct {
	FileId       string    `json:"file_id"`
	FileUniqueId string    `json:"file_unique_id"`
	Width        int       `json:"width"`
	Height       int       `json:"height"`
	Duration     int       `json:"duration"`
	Thumb        PhotoSize `json:"thumb,omitempty"`
	MIMEType     string    `json:"mime_type,omitempty"`
	FileSize     int       `json:"file_size,omitempty"`
}
