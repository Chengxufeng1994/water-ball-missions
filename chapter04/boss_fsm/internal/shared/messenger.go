package shared

type Messenger interface {
	SendMessage(message string, authorIds ...string)
}
