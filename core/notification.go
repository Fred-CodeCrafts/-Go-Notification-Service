package core

type Priority int

const (
	HIGH Priority = iota
	MEDIUM
	LOW
)

type Notification struct {
	To       string   `json:"to"`
	Message  string   `json:"message"`
	Channel  string   `json:"channel"`  // e.g. "sms", "email"
	Priority Priority `json:"priority"` // 0 = HIGH
}
