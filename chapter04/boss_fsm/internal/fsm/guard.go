package fsm

import "github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/shared"

type Guard interface {
	Check(ctx shared.Context, event Event) bool
}
