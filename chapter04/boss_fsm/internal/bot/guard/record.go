package guard

import (
	botevent "github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/bot/event"
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/bot/fields"
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/fsm"
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/shared"
)

type RecordGuard struct {
}

var _ fsm.Guard = (*RecordGuard)(nil)

func NewRecordGuard() *RecordGuard {
	return &RecordGuard{}
}

func (r *RecordGuard) Check(ctx shared.Context, event fsm.Event) bool {
	if !(event.GetEventType() == botevent.NewMessageEvent) {
		return false
	}

	payload, ok := event.GetEventPayload().(botevent.NewMessagePayload)
	if !ok {
		return false
	}

	speakerId := ctx.GetValue(fields.RecorderId)
	if speakerId == nil {
		return false
	}

	return payload.AuthorID == speakerId
}
