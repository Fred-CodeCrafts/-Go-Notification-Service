package internal

import (
	"encoding/json"
	"notification-service/core"
)

func HandleMessage(raw string) {
	var n core.Notification
	err := json.Unmarshal([]byte(raw), &n)
	if err != nil {
		panic("❌ Failed to parse JSON: " + err.Error())
	}
	Notify(n)
}
