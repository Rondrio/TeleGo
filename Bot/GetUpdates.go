package Bot

import (
	"TeleGo/Structs"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type UpdateContext struct {
	Offset         int      `json:"offset,omitempty"`
	Limit          int      `json:"limit,omitempty"`
	Timeout        int      `json:"timeout,omitempty"`
	AllowedUpdates []string `json:"allowed_updates,omitempty"`
}

type GotUpdate struct {
	Ok     bool             `json:"ok"`
	Result []Structs.Update `json:"result"`
}

//ctx is optional
func (bot *Bot) GetUpdate(ctx *UpdateContext) (*GotUpdate, error) {
	var updates GotUpdate
	buf := new(bytes.Buffer)
	if ctx != nil {
		b, err := json.Marshal(*ctx)
		if err != nil {
			return nil, err
		}

		_, err = buf.Write(b)
		if err != nil && err != io.EOF {
			return nil, err
		}
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s%s/getUpdates", baseURL, bot.token), buf)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	_, err = buf.ReadFrom(resp.Body)
	if err != nil && err != io.EOF {
		return nil, err
	}

	err = json.Unmarshal(buf.Bytes(), &updates)
	return &updates, err
}
