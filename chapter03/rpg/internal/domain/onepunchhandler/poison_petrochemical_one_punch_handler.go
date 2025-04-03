package onepunchhandler

import (
	"fmt"

	"github.com/Chengxufeng1994/water-ball-missions/chapter03/rpg/internal/domain"
	"github.com/Chengxufeng1994/water-ball-missions/chapter03/rpg/internal/domain/state"
)

type PoisonPetrochemicalOnePunchHandler struct {
	next OnePunchHandler
}

func NewPoisonPetrochemicalOnePunchHandler(next OnePunchHandler) *PoisonPetrochemicalOnePunchHandler {
	return &PoisonPetrochemicalOnePunchHandler{
		next: next,
	}
}

func (handler *PoisonPetrochemicalOnePunchHandler) Handle(attacker, target domain.Unit) error {
	if handler.Match(target) {
		for i := 0; i < 3; i++ {
			target.OnDamage(80)
			fmt.Printf("%v 對 %v 造成 %d 點傷害。\n", attacker, target, 80)
		}
	} else if handler.next != nil {
		handler.next.Handle(attacker, target)
	}

	return nil
}

func (handler *PoisonPetrochemicalOnePunchHandler) Match(target domain.Unit) bool {
	return target.GetCurrentState().Equal(state.NewPoisonedState()) || target.GetCurrentState().Equal(state.NewPetrochemicalState())
}
