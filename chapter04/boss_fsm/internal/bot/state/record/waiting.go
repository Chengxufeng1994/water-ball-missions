package record

import (
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/fsm"
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/fsm/state"
)

type Waiting struct {
	*state.BaseState
}

var _ fsm.State = (*Waiting)(nil)

func NewWaitingState(entryAction, exitAction fsm.Action) *Waiting {
	return &Waiting{
		BaseState: state.NewBaseState("waiting", entryAction, exitAction),
	}
}
