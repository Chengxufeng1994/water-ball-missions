package guard

import (
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/fsm"
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/shared"
)

type FunctionGuard func(ctx shared.Context) bool

var _ fsm.Guard = (*FunctionGuard)(nil)

func (fg FunctionGuard) Check(ctx shared.Context, event fsm.Event) bool {
	return fg(ctx)
}
