package internal

import (
	"notification-service/adapters"
	"notification-service/core"
	"notification-service/interfaces"
	"fmt"
)

var adapterRegistry = map[string]interfaces.NotificationAdapter{
	"sms":      adapters.SMSAdapter{},
	"email":    adapters.EmailAdapter{},
	"whatsapp": adapters.WhatsAppAdapter{},
}

func Notify(n core.Notification) {
	adapter, exists := adapterRegistry[n.Channel]
	if !exists {
		fmt.Printf("⚠️ Unknown channel: %s\n", n.Channel)
		return
	}
	adapter.Send(n.To, n.Message)
}
