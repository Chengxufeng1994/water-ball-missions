package guard

import (
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/fsm"
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/shared"
)

type NotNullGuard struct {
	Key string
}

var _ fsm.Guard = (*NotNullGuard)(nil)

func NewNotNullGuard(key string) *NotNullGuard {
	return &NotNullGuard{Key: key}
}

func (n *NotNullGuard) Check(ctx shared.Context, event fsm.Event) bool {
	return ctx.GetValue(n.Key) != nil
}
