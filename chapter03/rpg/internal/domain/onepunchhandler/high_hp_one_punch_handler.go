package onepunchhandler

import (
	"fmt"

	"github.com/Chengxufeng1994/water-ball-missions/chapter03/rpg/internal/domain"
)

type HighHpOnePunchHandler struct {
	next OnePunchHandler
}

func NewHighHpOnePunchHandler(next OnePunchHandler) *HighHpOnePunchHandler {
	return &HighHpOnePunchHandler{
		next: next,
	}
}

func (handler *HighHpOnePunchHandler) Handle(attacker, target domain.Unit) error {
	if handler.Match(target) {
		target.OnDamage(300)
		fmt.Printf("%v 對 %v 造成 %d 點傷害。\n", attacker, target, 300)
	} else if handler.next != nil {
		handler.next.Handle(attacker, target)
	}
	return nil
}

func (handler *HighHpOnePunchHandler) Match(target domain.Unit) bool {
	return target.GetHP() >= 500
}
