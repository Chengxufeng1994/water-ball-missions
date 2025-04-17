package normal

import (
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/fsm"
)

type NormalState struct {
	*fsm.SubFiniteStateMachine
}

var _ interface {
	fsm.FiniteStateMachine
	fsm.State
} = (*NormalState)(nil)

func NewNormalState(subFiniteStateMachine *fsm.SubFiniteStateMachine) *NormalState {
	return &NormalState{
		SubFiniteStateMachine: subFiniteStateMachine,
	}
}
