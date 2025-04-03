package onepunchhandler

import (
	"fmt"

	"github.com/Chengxufeng1994/water-ball-missions/chapter03/rpg/internal/domain"
	"github.com/Chengxufeng1994/water-ball-missions/chapter03/rpg/internal/domain/state"
)

type CheerUpOnePunchHandler struct {
	next OnePunchHandler
}

func NewCheerUpOnePunchHandler(next OnePunchHandler) *CheerUpOnePunchHandler {
	return &CheerUpOnePunchHandler{
		next: next,
	}
}

func (handler *CheerUpOnePunchHandler) Handle(attacker, target domain.Unit) error {
	if handler.Match(target) {
		target.OnDamage(100)
		fmt.Printf("%v 對 %v 造成 %d 點傷害。\n", attacker, target, 100)
	} else if handler.next != nil {
		handler.next.Handle(attacker, target)
	}
	return nil
}

func (handler *CheerUpOnePunchHandler) Match(target domain.Unit) bool {
	return target.GetCurrentState().Equal(state.NewCheerUpState())
}
