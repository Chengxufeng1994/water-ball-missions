package eventhandler

import (
	"slices"

	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/bot"
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/shared"
)

type LogoutEventHandler struct {
	next shared.EventHandler
}

var _ shared.EventHandler = (*LogoutEventHandler)(nil)

func NewLogoutEventHandler(next shared.EventHandler) *LogoutEventHandler {
	return &LogoutEventHandler{
		next: next,
	}
}

func (h *LogoutEventHandler) HandleEvent(event *shared.Event, ctx shared.Context) {
	if event.Type == shared.LogoutEvent {
		ctx.SetValue(bot.Users, ctx.GetValue(bot.Users).(int)-1)
		userList := ctx.GetValue(bot.UsersList).([]string)
		for i, userID := range userList {
			if userID == event.Payload.(shared.LogoutPayload).UserID {
				userList = slices.Delete(userList, i, i+1)
				break
			}
		}
		ctx.SetValue(bot.UsersList, userList)

		ctx.GetFSM().ProcessEvent(event, ctx)
	} else if h.next != nil {
		h.next.HandleEvent(event, ctx)
	}
}
