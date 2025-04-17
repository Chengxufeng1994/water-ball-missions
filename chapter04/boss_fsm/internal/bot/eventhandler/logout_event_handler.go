package eventhandler

import (
	"slices"

	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/bot/event"
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/bot/fields"
	fsmpkg "github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/fsm"
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/shared"
)

type LogoutEventHandler struct {
	next fsmpkg.EventHandler
}

var _ fsmpkg.EventHandler = (*LogoutEventHandler)(nil)

func NewLogoutEventHandler(next fsmpkg.EventHandler) *LogoutEventHandler {
	return &LogoutEventHandler{
		next: next,
	}
}

func (h *LogoutEventHandler) HandleEvent(ctx shared.Context, finiteStateMachine fsmpkg.FiniteStateMachine, evt fsmpkg.Event) {
	if evt.GetEventType() == event.LogoutEvent {
		ctx.SetValue(fields.UserCount, ctx.GetValue(fields.UserCount).(int)-1)
		userList := ctx.GetValue(fields.UserList).([]*shared.Member)
		for i, user := range userList {
			if user.UserID == evt.GetEventPayload().(event.LogoutPayload).UserID {
				userList = slices.Delete(userList, i, i+1)
				break
			}
		}
		ctx.SetValue(fields.UserList, userList)
	} else if h.next != nil {
		h.next.HandleEvent(ctx, finiteStateMachine, evt)
	}
}
