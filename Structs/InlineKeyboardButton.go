package Structs

type InlineKeyboardButton struct {
	Text                         string       `json:"text"`
	URL                          string       `json:"url,omitempty"`
	LoginURL                     LoginURL     `json:"login_url,omitempty"`
	CallbackData                 string       `json:"callback_data,omitempty"`
	SwitchInlineQuery            string       `json:"switch_inline_query,omitempty"`
	SwitchInlineQueryCurrentChat string       `json:"switch_inline_query_current_chat,omitempty"`
	CallBackGame                 CallBackGame `json:"switch_inline_query_current_chat,omitempty"`
	Pay                          bool         `json:"pay,omitempty"`
}
