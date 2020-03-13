package Structs

type Contact struct {
	PhoneNumber string `json:"phone_number"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name,omitempty"`
	UserId      int    `json:"last_name,omitempty"`
	Vcard       string `json:"vcard,omitempty"`
}
