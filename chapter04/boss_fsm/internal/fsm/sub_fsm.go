package fsm

import (
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/shared"
)

type SubFiniteStateMachine struct {
	name        string
	parent      FiniteStateMachine
	entryAction Action
	exitAction  Action
	*BaseFiniteStateMachine
}

var _ interface {
	FiniteStateMachine
	State
} = (*SubFiniteStateMachine)(nil)

func NewSubFiniteStateMachine(name string, initialState State, entryAction, exitAction Action, initTransitions []Transition, transitions ...Transition) *SubFiniteStateMachine {
	return &SubFiniteStateMachine{
		name:                   name,
		entryAction:            entryAction,
		exitAction:             exitAction,
		BaseFiniteStateMachine: NewBaseFiniteStateMachine(initialState, transitions, initTransitions...),
	}
}

// EntryState implements State.
func (fsm *SubFiniteStateMachine) EntryState(ctx shared.Context, event Event) {
	fsm.entryAction.Execute()
	fsm.Initialize(ctx)
}

// ProcessEntryState implements State.
func (fsm *SubFiniteStateMachine) ProcessEntryState(ctx shared.Context) {}

// HandleEvent implements FiniteStateMachine.
func (fsm *SubFiniteStateMachine) HandleEvent(ctx shared.Context, event Event) {
	fsm.BaseFiniteStateMachine.HandleEvent(ctx, event)
}

// OnEvent implements State.
func (fsm *SubFiniteStateMachine) OnEvent(ctx shared.Context, event Event) Event {
	fsm.BaseFiniteStateMachine.HandleEvent(ctx, event)
	return nil
}

// ExitState implements State.
func (fsm *SubFiniteStateMachine) ExitState(ctx shared.Context, event Event) {
	fsm.CurrentState.ExitState(ctx, event)
	fsm.exitAction.Execute()
}

// ProcessExitState implements State.
func (fsm *SubFiniteStateMachine) ProcessExitState(ctx shared.Context) {}

// GetName implements State.
func (fsm *SubFiniteStateMachine) GetName() string {
	return fsm.name
}

// GetParent implements State.
func (fsm *SubFiniteStateMachine) GetParent() FiniteStateMachine {
	return fsm.parent
}

// SetParent implements State.
func (fsm *SubFiniteStateMachine) SetParent(parent FiniteStateMachine) {
	fsm.parent = parent
}

// SetContext implements State.
func (fsm *SubFiniteStateMachine) SetContext(ctx shared.Context) {
	fsm.CurrentState.SetContext(ctx)
}
