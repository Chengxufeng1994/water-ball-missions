package fsm

import "github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/shared"

type Transition interface {
	GetFromState() State
	GetToState() State
	GetGuard() Guard
	StartTransition(ctx shared.Context, event Event) State
	CanTrigger(event Event) bool
}
