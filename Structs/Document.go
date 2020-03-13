package Structs

type Document struct {
	FileId       string    `json:"file_id"`
	FileUniqueId string    `json:"file_unique_id"`
	Thumb        PhotoSize `json:"thumb,omitempty"`
	FileName     string    `json:"file_name,omitempty"`
	MIMEType     string    `json:"mime_type,omitempty"`
	FileSize     int       `json:"file_size,omitempty"`
}
