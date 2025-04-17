package eventhandler

import (
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/bot/fields"
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/fsm"
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/shared"
)

type LoginEventHandler struct {
	next fsm.EventHandler
}

var _ fsm.EventHandler = (*LoginEventHandler)(nil)

func NewLoginEventHandler(next fsm.EventHandler) *LoginEventHandler {
	return &LoginEventHandler{
		next: next,
	}
}

func (h *LoginEventHandler) HandleEvent(ctx shared.Context, finiteStateMachine fsm.FiniteStateMachine, event *fsm.BaseEvent) {
	if event.Type == fsm.LoginEvent {
		_ = event.Payload.(fsm.LoginPayload)
		ctx.SetValue(fields.UserCount, ctx.GetValue(fields.UserCount).(int)+1)
		ctx.SetValue(fields.UsersList, append(ctx.GetValue(fields.UsersList).([]string), event.Payload.(fsm.LoginPayload).UserID))
	} else if h.next != nil {
		h.next.HandleEvent(ctx, finiteStateMachine, event)
	}
}
