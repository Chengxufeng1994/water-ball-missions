package trigger

import (
	"strings"

	botevent "github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/bot/event"
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
	if event.GetEventType() == botevent.NewMessageEvent {
		payload := event.GetEventPayload().(botevent.NewMessagePayload)
		payload.Content = strings.TrimSpace(payload.Content)
		ret := strings.Contains(m.Message, payload.Content)
		return ret
	}
	return false
}
