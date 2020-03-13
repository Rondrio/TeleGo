package Structs

type PollAnswer struct {
	PollId    int   `json:"poll_id"`
	User      User  `json:"user"`
	OptionIds []int `json:"option_ids"`
}
