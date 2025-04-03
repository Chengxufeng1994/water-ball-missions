package onepunchhandler

import "github.com/Chengxufeng1994/water-ball-missions/chapter03/rpg/internal/domain"

type OnePunchHandler interface {
	Handle(attack, target domain.Unit) error
	Match(target domain.Unit) bool
}
