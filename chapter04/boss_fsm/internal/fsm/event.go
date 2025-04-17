package fsm

type (
	EventType string

	EventPayload any
	Event        interface {
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
	return &BaseEvent{Type: eventType, Payload: payload}
}

func NewEvent(eventType EventType, payload EventPayload) Event {
	return newBaseEvent(eventType, payload)
}

// GetEventType implements Event.
func (b *BaseEvent) GetEventType() EventType {
	return b.Type
}

// GetEventPayload implements Event.
func (b *BaseEvent) GetEventPayload() EventPayload {
	return b.Payload
}
