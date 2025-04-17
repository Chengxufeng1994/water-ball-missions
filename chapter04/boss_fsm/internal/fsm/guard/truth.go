package guard

import (
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/fsm"
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/shared"
)

type TruthGuard struct {
}

var _ fsm.Guard = (*TruthGuard)(nil)

func NewTruthGuard() *TruthGuard {
	return &TruthGuard{}
}

func (t *TruthGuard) Check(ctx shared.Context, event fsm.Event) bool {
	return true
}
