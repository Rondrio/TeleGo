package Bot

import (
	"TeleGo/Structs"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type MessageHandle struct {
	ChatId                string      `json:"chat_id"`
	Text                  string      `json:"text"`
	ParseMode             string      `json:"parse_mode,omitempty"`
	DisableWebPagePreview bool        `json:"disable_web_page_preview,omitempty"`
	DisableNotification   bool        `json:"disable_notification,omitempty"`
	ReplyToMessageId      int         `json:"reply_to_message_id,omitempty"`
	ReplyMarkup           interface{} `json:"reply_markup,omitempty"`
}

type MessageReturn struct {
	Ok     bool          `json:"ok"`
	Result MessageResult `json:"result"`
}

type MessageResult struct {
	MessageID int          `json:"message_id"`
	From      Structs.User `json:"from"`
	Chat      Structs.Chat `json:"chat"`
	Date      int          `json:"date"`
	Text      string       `json:"text"`
}

func CreateMessage(chatId string, content string) *MessageHandle {
	return &MessageHandle{
		ChatId: chatId,
		Text:   content,
	}
}

func (bot *Bot) SendMessage(handle *MessageHandle) (*MessageReturn, error) {
	var response MessageReturn

	b, err := json.Marshal(*handle)
	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	_, err = buf.Write(b)
	if err != nil && err != io.EOF {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s%s/sendMessage", baseURL, bot.token), buf)
	if err != nil {
		return nil, err
	}

	req.Header.Add("content-type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	_, err = buf.ReadFrom(resp.Body)
	if err != nil && err != io.EOF {
		return nil, err
	}

	err = json.Unmarshal(buf.Bytes(), &response)
	return &response, err
}
