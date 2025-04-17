package guard

import (
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/bot/fields"
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/fsm"
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/shared"
)

type UserCountLimitGuard struct {
	ng NumberGuard
}

var _ fsm.Guard = (*UserCountLimitGuard)(nil)

func NewUserCountLimitGuard(ng NumberGuard) *UserCountLimitGuard {
	return &UserCountLimitGuard{
		ng: ng,
	}
}

func (g *UserCountLimitGuard) Check(ctx shared.Context, event fsm.Event) bool {
	return g.ng.Check(ctx.GetValue(fields.UserCount).(int))
}
