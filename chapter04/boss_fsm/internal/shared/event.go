package shared

type Event struct {
	Type    EventType
	Payload EventPayload
}

func NewEvent(eventType EventType, payload EventPayload) *Event {
	return &Event{
		Type:    eventType,
		Payload: payload,
	}
}

func NewLoginEvent(payload LoginPayload) *Event {
	return NewEvent(LoginEvent, payload)
}

func NewLogoutEvent(payload LogoutPayload) *Event {
	return NewEvent(LogoutEvent, payload)
}

func NewNewMessageEvent(payload NewMessagePayload) *Event {
	return NewEvent(NewMessageEvent, payload)
}

func NewNewPostEvent(payload NewPostPayload) *Event {
	return NewEvent(NewPostEvent, payload)
}
