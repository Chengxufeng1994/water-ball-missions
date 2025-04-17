package trigger

import (
	"time"

	botevent "github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/bot/event"
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/fsm"
)

type ElapsedTime struct {
	Duration time.Duration
}

var _ fsm.Trigger = (*ElapsedTime)(nil)

func NewElapsedTimeTrigger(duration time.Duration) *ElapsedTime {
	return &ElapsedTime{
		Duration: duration,
	}
}

func (e *ElapsedTime) Match(event fsm.Event) bool {
	if event.GetEventType() == botevent.TimeElapsedEvent {
		payload := event.GetEventPayload().(botevent.ElapsedTimePayload)
		ret := payload.Duration >= e.Duration
		return ret
	}
	return false
}
