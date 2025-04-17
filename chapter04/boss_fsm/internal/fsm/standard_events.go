package fsm

const (
	EmptyEvent EventType = ""
	InitEvent  EventType = "init"
	ExitEvent  EventType = "exit"
)

var StdEvents = map[EventType]Event{
	InitEvent: NewStdEvent(InitEvent),
	ExitEvent: NewStdEvent(ExitEvent),
}

type StandardEvent struct {
	EventType EventType
}

func NewStdEvent(eventType EventType) Event {
	return &StandardEvent{EventType: eventType}
}

func (s *StandardEvent) GetEventType() EventType {
	return s.EventType
}

func (s *StandardEvent) GetEventPayload() EventPayload {
	return nil
}

func NewInitEvent() Event {
	return NewStdEvent(InitEvent)
}
