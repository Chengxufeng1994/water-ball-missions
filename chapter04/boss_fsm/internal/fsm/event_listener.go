package fsm

import "github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/shared"

type (
	EventListener interface {
		OnEvent(event Event)
	}

	EventHandler interface {
		HandleEvent(ctx shared.Context, fsm FiniteStateMachine, event Event)
	}
)
