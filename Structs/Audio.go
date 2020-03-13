package Structs

type Audio struct {
	FileId       string    `json:"file_id"`
	FileUniqueId string    `json:"file_unique_id"`
	Duration     int       `json:"duration"`
	Performer    string    `json:"performer,omitempty"`
	Title        string    `json:"title,omitempty"`
	MIMEType     string    `json:"mime_type,omitempty"`
	FileSize     int       `json:"file_size,omitempty"`
	Thumb        PhotoSize `json:"thumb,omitempty"`
}
