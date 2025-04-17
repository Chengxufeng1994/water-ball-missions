package guard

import (
	"slices"

	botevent "github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/bot/event"
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/fsm"
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/shared"
)

type TaggedGuard struct {
	value string
}

var _ fsm.Guard = (*TaggedGuard)(nil)

func NewTaggedGuard(value string) *TaggedGuard {
	return &TaggedGuard{
		value: value,
	}
}

func (g *TaggedGuard) Check(ctx shared.Context, event fsm.Event) bool {
	payload, ok := event.GetEventPayload().(botevent.NewMessagePayload)
	if !ok {
		return false
	}
	return slices.Contains(payload.Tags, g.value)
}
