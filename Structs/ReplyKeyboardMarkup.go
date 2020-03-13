package Structs

type ReplyKeyboardMarkup struct {
	Keyboard        []KeyboardButton `json:"keyboard"`
	ResizeKeyboard  bool             `json:"resize_keyboard,omitempty"`
	OneTimeKeyboard bool             `json:"one_time_keyboard,omitempty"`
	Selective       bool             `json:"selective,omitempty"`
}
