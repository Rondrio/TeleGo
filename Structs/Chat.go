package Structs

type Chat struct {
	Id        int    `json:"id"`
	Type      string `json:"type"`
	Title     string `json:"title,,omitempty"`
	Username  string `json:"username,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`

	Photo            ChatPhoto       `json:"chat_photo,omitempty"`
	Description      string          `json:"description,omitempty"`
	InviteLink       string          `json:"invite_link,omitempty"`
	PinnedMessage    *Message        `json:"pinned_message,omitempty"`
	Permissions      ChatPermissions `json:"permissions,omitempty"`
	SlowModeDelay    int             `json:"slow_mode_delay,omitempty"`
	StickerSetName   string          `json:"sticker_set_name,omitempty"`
	CanSetStickerSet bool            `json:"can_set_sticker_set,omitempty"`
}
