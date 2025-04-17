package event

import (
	"time"

	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/fsm"
)

type ElapsedTimePayload struct {
	Duration time.Duration `json:"Duration"`
}

func NewTimeElapsedEvent(payload ElapsedTimePayload) fsm.Event {
	return fsm.NewEvent(TimeElapsedEvent, payload)
}
