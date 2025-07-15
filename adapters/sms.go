package adapters

import "fmt"

type SMSAdapter struct{}

func (s SMSAdapter) Send(to, message string) error {
	fmt.Printf("📲 SMS sent to %s: %s\n", to, message)
	return nil
}

func (s SMSAdapter) ChannelName() string {
	return "sms"
}
