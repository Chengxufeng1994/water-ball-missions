package guard

import (
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/fsm"
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/shared"
)

type FalsyGuard struct{}

var _ fsm.Guard = (*FalsyGuard)(nil)

func NewFalsyGuard() *FalsyGuard {
	return &FalsyGuard{}
}

func (f *FalsyGuard) Check(ctx shared.Context, event fsm.Event) bool {
	return false
}
