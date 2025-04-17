package guard

import (
	botevent "github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/bot/event"
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/bot/fields"
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/fsm"
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/shared"
)

type AdminGuard struct{}

var _ fsm.Guard = (*AdminGuard)(nil)

func NewAdminGuard() *AdminGuard {
	return &AdminGuard{}
}

func (a *AdminGuard) Check(ctx shared.Context, event fsm.Event) bool {
	members := ctx.GetValue(fields.UserList).([]*shared.Member)
	payload, ok := event.GetEventPayload().(botevent.NewMessagePayload)
	if !ok {
		return false
	}
	for _, member := range members {
		if member.UserID == payload.AuthorID {
			return member.IsAdmin
		}
	}
	return false
}
