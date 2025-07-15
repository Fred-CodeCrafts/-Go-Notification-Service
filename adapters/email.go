package adapters

import "fmt"

type EmailAdapter struct{}

func (e EmailAdapter) Send(to, message string) error {
	fmt.Printf("📧 Email sent to %s: %s\n", to, message)
	return nil
}

func (e EmailAdapter) ChannelName() string {
	return "email"
}
