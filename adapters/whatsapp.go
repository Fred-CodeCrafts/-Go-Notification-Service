package adapters

import "fmt"

type WhatsAppAdapter struct{}

func (w WhatsAppAdapter) Send(to, message string) error {
	fmt.Printf("💬 WhatsApp sent to %s: %s\n", to, message)
	return nil
}

func (w WhatsAppAdapter) ChannelName() string {
	return "whatsapp"
}
