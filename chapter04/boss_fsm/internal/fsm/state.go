package fsm

import "github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/shared"

type State interface {
	EntryState(ctx shared.Context, event Event)
	ProcessEntryState(ctx shared.Context)
	ExitState(ctx shared.Context, event Event)
	ProcessExitState(ctx shared.Context)
	OnEvent(ctx shared.Context, event Event) Event
	GetName() string
	SetContext(ctx shared.Context)
	GetParent() FiniteStateMachine
	SetParent(parent FiniteStateMachine)
}
