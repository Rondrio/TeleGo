package Structs

type LoginURL struct {
URL string `json:"url"`
ForwardText string `json:"forward_text,omitempty"`
BotUsername string `json:"bot_username,omitempty"`
RequestWriteAccess bool `json:"request_write_access,omitempty"`
}
