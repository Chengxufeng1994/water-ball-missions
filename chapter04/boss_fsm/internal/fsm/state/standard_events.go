package state

import (
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/fsm"
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/fsm/action"
)

type StandardState struct {
	*BaseState
}

func newStandardState(name string) fsm.State {
	return &StandardState{
		BaseState: NewBaseState(name, action.NewNoAction(), action.NewNoAction(), nil),
	}
}

func NewInitState() fsm.State {
	return newStandardState("init")
}
