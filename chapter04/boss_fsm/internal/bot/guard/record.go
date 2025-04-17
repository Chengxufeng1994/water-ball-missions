package guard

import (
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
	if !(event.GetEventType() == fsm.StopBroadcastingEvent) {
		return false
	}

	payload, ok := event.GetEventPayload().(fsm.StopBroadcastingPayload)
	if !ok {
		return false
	}

	recordId := ctx.GetValue(fields.RecorderId)
	if recordId == nil {
		return false
	}

	return payload.SpeakerID == recordId
}
