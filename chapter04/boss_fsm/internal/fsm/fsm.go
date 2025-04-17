package fsm

import (
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/shared"
)

type FiniteStateMachine interface {
	Initialize(ctx shared.Context)
	GetCurrentState() State
	HandleEvent(ctx shared.Context, event Event)
}

type BaseFiniteStateMachine struct {
	CurrentState       State
	InitialTransitions []Transition
	Transitions        []Transition
}

var _ FiniteStateMachine = (*BaseFiniteStateMachine)(nil)

func NewBaseFiniteStateMachine(initialState State, initTransitions []Transition, transitions ...Transition) *BaseFiniteStateMachine {

	fsm := &BaseFiniteStateMachine{
		CurrentState: initialState,
		Transitions:  append(initTransitions, transitions...),
	}

	for _, transition := range fsm.Transitions {
		transition.GetFromState().SetParent(fsm)
		transition.GetToState().SetParent(fsm)
	}

	return fsm
}

func (fsm *BaseFiniteStateMachine) Initialize(ctx shared.Context) {
	fsm.HandleEvent(ctx, NewInitEvent())
}

func (fsm *BaseFiniteStateMachine) HandleEvent(ctx shared.Context, event Event) {
	fsm.CurrentState.OnEvent(ctx, event)

	for _, trans := range fsm.Transitions {
		if trans.CanTrigger(event) {
			if trans.GetGuard() != nil && !trans.GetGuard().Check(ctx, event) {
				continue
			}
			currentState := trans.StartTransition(ctx, event)
			currentState.SetContext(ctx)
			fsm.CurrentState = currentState
			break
		}
	}
}

func (fsm *BaseFiniteStateMachine) GetCurrentState() State {
	return fsm.CurrentState
}
