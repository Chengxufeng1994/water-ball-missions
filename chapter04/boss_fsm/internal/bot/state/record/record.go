package record

import (
	botevent "github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/bot/event"
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/bot/fields"
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/fsm"
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/shared"
)

const RecordQuota = 3

type Record struct {
	*fsm.SubFiniteStateMachine
}

func NewRecordState(subFiniteStateMachine *fsm.SubFiniteStateMachine) *Record {
	return &Record{
		SubFiniteStateMachine: subFiniteStateMachine,
	}
}

func (s *Record) EntryState(ctx shared.Context, event fsm.Event) {
	s.SubFiniteStateMachine.EntryState(ctx, event)
	if event.GetEventType() == botevent.NewMessageEvent {
		ctx.SetValue(fields.RecorderId, event.GetEventPayload().(botevent.NewMessagePayload).AuthorID)
	}
	ctx.SetValue(fields.Quota, ctx.GetValue(fields.Quota).(int)-RecordQuota)
}
