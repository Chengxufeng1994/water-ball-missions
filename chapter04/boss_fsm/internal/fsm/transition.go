package fsm

import "github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/shared"

type Transition struct {
	FromState shared.State
	ToState   shared.State
	EventType shared.EventType
	Trigger   Trigger
	Guard     Guard
	Action    Action
}

func NewTransition(fromState, toState shared.State, eventType shared.EventType, guard Guard, action Action) *Transition {
	return &Transition{
		FromState: fromState,
		ToState:   toState,
		EventType: eventType,
		Guard:     guard,
		Action:    action,
	}
}
