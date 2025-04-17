package record

import (
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/bot/fields"
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/fsm"
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/shared"
)

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
	if event.GetEventType() == fsm.NewMessageEvent {
		ctx.SetValue(fields.RecorderId, event.GetEventPayload().(fsm.NewMessagePayload).AuthorID)
	}
}
