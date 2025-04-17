package eventhandler

import (
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/bot"
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/shared"
)

type LoginEventHandler struct {
	next shared.EventHandler
}

func NewLoginEventHandler(next shared.EventHandler) *LoginEventHandler {
	return &LoginEventHandler{
		next: next,
	}
}

func (h *LoginEventHandler) HandleEvent(event *shared.Event, ctx shared.Context) {
	if event.Type == shared.LoginEvent {
		_ = event.Payload.(shared.LoginPayload)
		ctx.SetValue(bot.Users, ctx.GetValue(bot.Users).(int)+1)
		ctx.SetValue(bot.UsersList, append(
			ctx.GetValue(bot.UsersList).([]string),
			event.Payload.(shared.LoginPayload).UserID),
		)
		ctx.GetFSM().ProcessEvent(event, ctx)
	} else if h.next != nil {
		h.next.HandleEvent(event, ctx)
	}
}
