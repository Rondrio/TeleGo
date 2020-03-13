package Bot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	baseURL = "https://api.telegram.org/bot"
)

type Bot struct {
	token string
}



func New(token string) *Bot {
	return &Bot{
		token: token,
	}
}

func (bot *Bot) SendMessage(handle *MessageHandle) (*MessageReturn,error) {
	var response MessageReturn

	b,err := json.Marshal(*handle)
	if err != nil{
		return nil,err
	}

	buf := new(bytes.Buffer)
	_,err = buf.Write(b)
	if err != nil && err != io.EOF{
		return nil,err
	}

	req,err := http.NewRequest("POST",fmt.Sprintf("%s%s/sendMessage",baseURL,bot.token),buf)
	if err != nil{
		return nil,err
	}

	req.Header.Add("content-type","application/json")

	resp,err := http.DefaultClient.Do(req)
	if err != nil{
		return nil,err
	}

	_,err = buf.ReadFrom(resp.Body)
	if err != nil && err != io.EOF {
		return nil,err
	}

	err = json.Unmarshal(buf.Bytes(),&response)
	return &response,err
}