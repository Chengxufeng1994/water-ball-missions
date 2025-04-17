package state

import (
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/fsm"
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/shared"
)

type (
	BotFsmAdapter interface {
		shared.Messenger
		shared.RecorderReplay
		shared.Speaker
	}

	BaseState struct {
		Name        string
		State       fsm.State
		Parent      fsm.FiniteStateMachine
		Adapter     BotFsmAdapter
		Ctx         shared.Context
		EntryAction fsm.Action
		ExitAction  fsm.Action
	}
)

var _ fsm.State = (*BaseState)(nil)

func NewBaseState(name string, entryAction fsm.Action, exitAction fsm.Action, fsmAdapter BotFsmAdapter) *BaseState {
	return &BaseState{
		Name:        name,
		Ctx:         nil,
		Adapter:     fsmAdapter,
		EntryAction: entryAction,
		ExitAction:  exitAction,
	}
}

func (b *BaseState) EntryState(ctx shared.Context, event fsm.Event) {
	b.EntryAction.Execute()
	b.ProcessEntryState(ctx)
}
func (b *BaseState) ProcessEntryState(ctx shared.Context) {}
func (b *BaseState) ExitState(ctx shared.Context, event fsm.Event) {
	b.ExitAction.Execute()
	b.ProcessExitState(ctx)
}
func (b *BaseState) ProcessExitState(ctx shared.Context)                   {}
func (b *BaseState) OnEvent(ctx shared.Context, event fsm.Event) fsm.Event { return nil }
func (b *BaseState) GenerateMessage(event fsm.Event) string                { return "" }
func (b *BaseState) GetName() string                                       { return b.Name }
func (b *BaseState) GetParent() fsm.FiniteStateMachine                     { return b.Parent }
func (b *BaseState) SetParent(parent fsm.FiniteStateMachine)               { b.Parent = parent }
func (b *BaseState) SetContext(ctx shared.Context)                         { b.Ctx = ctx }
