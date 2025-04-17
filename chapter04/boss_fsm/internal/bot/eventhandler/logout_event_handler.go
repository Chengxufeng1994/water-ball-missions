package eventhandler

import (
	"slices"

	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/bot/fields"
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/fsm"
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/shared"
)

type LogoutEventHandler struct {
	next fsm.EventHandler
}

var _ fsm.EventHandler = (*LogoutEventHandler)(nil)

func NewLogoutEventHandler(next fsm.EventHandler) *LogoutEventHandler {
	return &LogoutEventHandler{
		next: next,
	}
}

func (h *LogoutEventHandler) HandleEvent(ctx shared.Context, finiteStateMachine fsm.FiniteStateMachine, event *fsm.BaseEvent) {
	if event.Type == fsm.LogoutEvent {
		ctx.SetValue(fields.UserCount, ctx.GetValue(fields.UserCount).(int)-1)
		userList := ctx.GetValue(fields.UsersList).([]string)
		for i, userID := range userList {
			if userID == event.Payload.(fsm.LogoutPayload).UserID {
				userList = slices.Delete(userList, i, i+1)
				break
			}
		}
		ctx.SetValue(fields.UsersList, userList)
	} else if h.next != nil {
		h.next.HandleEvent(ctx, finiteStateMachine, event)
	}
}
