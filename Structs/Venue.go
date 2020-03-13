package Structs

type Venue struct {
	Location       Location `json:"location"`
	Title          string   `json:"title"`
	Address        string   `json:"address"`
	FourSquareId   string   `json:"foursquare_id,omitempty"`
	FourSquareType string   `json:"foursquare_type,omitempty"`
}
