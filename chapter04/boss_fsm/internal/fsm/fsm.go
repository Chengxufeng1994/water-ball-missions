package fsm

import "github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/shared"

type FiniteStateMachine struct {
	CurrentState shared.State
	Transitions  []*Transition
}

var _ shared.FiniteStateMachine = (*FiniteStateMachine)(nil)

func NewFiniteStateMachine(initialState shared.State, transitions []*Transition) *FiniteStateMachine {
	return &FiniteStateMachine{
		CurrentState: initialState,
		Transitions:  transitions,
	}
}

func (fsm *FiniteStateMachine) GetCurrentState() shared.State {
	return fsm.CurrentState
}

func (fsm *FiniteStateMachine) ProcessEvent(event *shared.Event, ctx shared.Context) {
	for _, transition := range fsm.Transitions {
		if transition.FromState == fsm.CurrentState && transition.EventType == event.Type {
			if transition.Guard != nil && !transition.Guard.Check(ctx) {
				return
			}

			transition.FromState.ExitState()
			if transition.Action != nil {
				transition.Action.Execute()
			}
			transition.ToState.EntryState()
			fsm.CurrentState = transition.ToState
			fsm.CurrentState.SetContext(ctx)
		}
	}
}
