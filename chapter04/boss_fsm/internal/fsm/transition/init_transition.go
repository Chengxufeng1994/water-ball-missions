package transition

import (
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/fsm"
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/fsm/action"
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/fsm/state"
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/fsm/trigger"
)

type InitTransition struct {
	*BaseTransition
}

var _ fsm.Transition = (*InitTransition)(nil)

func NewInitTransition(fsmState fsm.State, guard fsm.Guard) *InitTransition {
	return &InitTransition{
		BaseTransition: NewBaseTransition(state.NewInitState(), fsmState, trigger.NewEventTrigger(fsm.InitEvent), guard, action.NewNoAction()),
	}
}
