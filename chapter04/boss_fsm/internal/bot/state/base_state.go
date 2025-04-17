package state

import (
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/fsm"
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/shared"
)

type BaseState struct {
	EntryAction fsm.Action
	ExitAction  fsm.Action

	Ctx shared.Context
}

var _ shared.State = (*BaseState)(nil)

func NewBaseState(entryAction fsm.Action, exitAction fsm.Action) *BaseState {
	return &BaseState{
		EntryAction: entryAction,
		ExitAction:  exitAction,
	}
}

func (b *BaseState) EntryState() {
	if b.EntryAction != nil {
		b.EntryAction.Execute()
	}
}

func (b *BaseState) ExitState() {
	if b.ExitAction != nil {
		b.ExitAction.Execute()
	}
}

func (b *BaseState) GenerateMessage(channelType, authorId string) string {
	panic("unimplemented")
}

func (b *BaseState) SetContext(ctx shared.Context) {
	b.Ctx = ctx
}
