package Structs

type ResponseParameters struct {
	MigrateToChatId int `json:"migrate_to_chat_id,omitempty"`
	RetryAfter int `json:"retry_after,omitempty"`
}
