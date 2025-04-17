package transition

import (
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/fsm"
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/shared"
)

type BaseTransition struct {
	FromState fsm.State
	ToState   fsm.State
	Trigger   fsm.Trigger
	Guard     fsm.Guard
	Action    fsm.Action
}

var _ fsm.Transition = (*BaseTransition)(nil)

func NewBaseTransition(
	fromState, toState fsm.State,
	trigger fsm.Trigger,
	guard fsm.Guard,
	action fsm.Action,
) *BaseTransition {
	return &BaseTransition{
		FromState: fromState,
		ToState:   toState,
		Trigger:   trigger,
		Guard:     guard,
		Action:    action,
	}
}

func (t *BaseTransition) StartTransition(ctx shared.Context, event fsm.Event) fsm.State {
	// fmt.Println("[BaseTransition.StartTransition]", t.FromState.GetName(), "->", t.ToState.GetName())

	t.FromState.ExitState(ctx, event)

	t.Action.Execute()

	t.ToState.EntryState(ctx, event)

	return t.ToState
}

func (t *BaseTransition) CanTrigger(event fsm.Event) bool {
	return t.Trigger.Match(event)
}

func (t *BaseTransition) GetFromState() fsm.State {
	return t.FromState
}

func (t *BaseTransition) GetToState() fsm.State {
	return t.ToState
}

func (t *BaseTransition) GetGuard() fsm.Guard {
	return t.Guard
}
