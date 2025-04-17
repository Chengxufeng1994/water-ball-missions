package trigger

import "github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/fsm"

type EventTrigger struct {
	EventType fsm.EventType
}

var _ fsm.Trigger = (*EventTrigger)(nil)

func NewEventTrigger(eventType fsm.EventType) *EventTrigger {
	return &EventTrigger{
		EventType: eventType,
	}
}

func (e *EventTrigger) Match(event fsm.Event) bool {
	return event.GetEventType() == e.EventType
}
