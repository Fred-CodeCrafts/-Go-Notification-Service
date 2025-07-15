package interfaces

type NotificationAdapter interface {
	Send(to string, message string) error
	ChannelName() string
}
