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
	// Let the current state handle the event first
	if nextEvent := fsm.CurrentState.OnEvent(ctx, event); nextEvent != nil {
		fsm.HandleEvent(ctx, nextEvent)
		return
	}

	// Select transitions to check: all for InitEvent, otherwise only those from current state
	var candidates []Transition
	if event.GetEventType() == InitEvent {
		candidates = fsm.Transitions
	} else {
		for _, trans := range fsm.Transitions {
			if trans.GetFromState() == fsm.CurrentState {
				candidates = append(candidates, trans)
			}
		}
	}

	// Try to trigger a transition
	for _, trans := range candidates {
		if !trans.CanTrigger(event) {
			continue
		}
		if guard := trans.GetGuard(); guard != nil && !guard.Check(ctx, event) {
			continue
		}
		nextState := trans.StartTransition(ctx, event)
		nextState.SetContext(ctx)
		fsm.CurrentState = nextState
		break
	}
}

func (fsm *BaseFiniteStateMachine) GetCurrentState() State {
	return fsm.CurrentState
}
