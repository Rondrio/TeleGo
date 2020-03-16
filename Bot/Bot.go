package Bot

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
