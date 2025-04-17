package guard

import (
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/fsm"
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/shared"
)

type AndGuard struct {
	guards []fsm.Guard
}

var _ fsm.Guard = (*AndGuard)(nil)

func NewAndGuard(guards ...fsm.Guard) *AndGuard {
	return &AndGuard{
		guards: guards,
	}
}

func (g *AndGuard) Check(ctx shared.Context, event fsm.Event) bool {
	for _, guard := range g.guards {
		if !guard.Check(ctx, event) {
			return false
		}
	}
	return true
}
