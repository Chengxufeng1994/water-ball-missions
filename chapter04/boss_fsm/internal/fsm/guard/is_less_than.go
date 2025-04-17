package guard

import (
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/shared"
)

type IsLessThan struct {
	Key   string
	Value int
}

func NewIsLessThan(key string, value int) *IsLessThan {
	return &IsLessThan{
		Key:   key,
		Value: value,
	}
}

func (g *IsLessThan) Check(ctx shared.Context) bool {
	return ctx.GetValue(g.Key).(int) < g.Value
}
