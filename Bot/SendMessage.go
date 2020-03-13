package Bot

import "TeleGo/Structs"

type MessageHandle struct {
	ChatId string `json:"chat_id"`
	Text string `json:"text"`
	ParseMode string `json:"parse_mode,omitempty"`
	DisableWebPagePreview bool `json:"disable_web_page_preview,omitempty"`
	DisableNotification bool `json:"disable_notification,omitempty"`
	ReplyToMessageId int `json:"reply_to_message_id,omitempty"`
	ReplyMarkup interface{} `json:"reply_markup,omitempty"`
}

type MessageReturn struct {
	Ok     bool   `json:"ok"`
	Result Result `json:"result"`
}

type Result struct {
	MessageID int    `json:"message_id"`
	From      Structs.User   `json:"from"`
	Chat      Structs.Chat   `json:"chat"`
	Date      int    `json:"date"`
	Text      string `json:"text"`
}

func CreateMessage(chatId string, content string) *MessageHandle {
	return &MessageHandle{
		ChatId:                chatId,
		Text:                  content,
	}
}