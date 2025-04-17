package trigger

import (
	"strings"

	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/fsm"
)

type NewMessageTrigger struct {
	Message string
}

var _ fsm.Trigger = (*NewMessageTrigger)(nil)

func NewNewMessageTrigger(message string) *NewMessageTrigger {
	return &NewMessageTrigger{
		Message: message,
	}
}

func (m *NewMessageTrigger) Match(event fsm.Event) bool {
	if event.GetEventType() == fsm.NewMessageEvent {
		payload := event.GetEventPayload().(fsm.NewMessagePayload)
		payload.Content = strings.TrimSpace(payload.Content)
		return strings.Contains(m.Message, payload.Content)
	}
	return false
}
