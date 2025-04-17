package guard

import (
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/bot/fields"
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/fsm"
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/shared"
)

type QuotaGuard struct {
	numberGuard NumberGuard
}

var _ fsm.Guard = (*QuotaGuard)(nil)

func NewQuotaLimitGuard(numberGuard NumberGuard) QuotaGuard {
	return QuotaGuard{
		numberGuard: numberGuard,
	}
}

func (g QuotaGuard) Check(ctx shared.Context, event fsm.Event) bool {
	return g.numberGuard.Check(ctx.GetValue(fields.Quota).(int))
}
