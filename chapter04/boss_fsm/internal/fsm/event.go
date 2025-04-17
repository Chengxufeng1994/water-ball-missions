package fsm

type (
	Event interface {
		GetEventType() EventType
		GetEventPayload() EventPayload
	}
	BaseEvent struct {
		Type    EventType
		Payload EventPayload
	}
)

var _ Event = (*BaseEvent)(nil)

func newBaseEvent(eventType EventType, payload EventPayload) *BaseEvent {
	return &BaseEvent{
		Type:    eventType,
		Payload: payload,
	}
}

// GetEventType implements Event.
func (b *BaseEvent) GetEventType() EventType {
	return b.Type
}

// GetEventPayload implements Event.
func (b *BaseEvent) GetEventPayload() EventPayload {
	return b.Payload
}

func NewLoginEvent(payload LoginPayload) *BaseEvent {
	return newBaseEvent(LoginEvent, payload)
}

func NewLogoutEvent(payload LogoutPayload) *BaseEvent {
	return newBaseEvent(LogoutEvent, payload)
}

func NewNewMessageEvent(payload NewMessagePayload) *BaseEvent {
	return newBaseEvent(NewMessageEvent, payload)
}

func NewNewPostEvent(payload NewPostPayload) *BaseEvent {
	return newBaseEvent(NewPostEvent, payload)
}

func NewGoBroadcastingEvent(payload GoBroadcastingPayload) *BaseEvent {
	return newBaseEvent(GoBroadcastingEvent, payload)
}

func NewStopBroadcastingEvent(payload StopBroadcastingPayload) *BaseEvent {
	return newBaseEvent(StopBroadcastingEvent, payload)
}

func NewSpeakEvent(payload SpeakPayload) *BaseEvent {
	return newBaseEvent(SpeakEvent, payload)
}
