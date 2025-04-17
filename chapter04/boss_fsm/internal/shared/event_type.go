package shared

type EventType string

const (
	LoginEvent            EventType = "login"
	LogoutEvent           EventType = "logout"
	TimeElapsedEvent      EventType = "time_elapsed"
	NewMessageEvent       EventType = "new_message"
	NewPostEvent          EventType = "new_post"
	GoBroadcastingEvent   EventType = "go_broadcasting"
	SpeakEvent            EventType = "speak"
	StopBroadcastingEvent EventType = "stop_broadcasting"
	EndEvent              EventType = "end"
)
