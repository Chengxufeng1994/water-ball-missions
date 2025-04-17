package action

import "github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/fsm"

type NoAction struct{}

var _ fsm.Action = (*NoAction)(nil)

func NewNoAction() *NoAction {
	return &NoAction{}
}

func (n *NoAction) Execute() {}
