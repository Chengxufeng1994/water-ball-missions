package guard

import "github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/shared"

type IsGreaterThan struct {
	Key   string
	Value int
}

func NewIsGreaterThan(key string, value int) *IsGreaterThan {
	return &IsGreaterThan{
		Key:   key,
		Value: value,
	}
}

func (g *IsGreaterThan) Check(ctx shared.Context) bool {
	return ctx.GetValue(g.Key).(int) >= g.Value
}
