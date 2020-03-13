package Structs

type MessageEntity struct {
	Type   string `json:"type"`
	OffSet int    `json:"offset"`
	Length int    `json:"length"`

	URL      string `json:"url,omitempty"`
	User     User   `json:"user"`
	Language string `json:"language"`
}
