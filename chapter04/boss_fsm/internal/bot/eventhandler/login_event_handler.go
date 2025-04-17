package eventhandler

import (
	botevent "github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/bot/event"
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/bot/fields"
	fsmpkg "github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/fsm"
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/shared"
)

type LoginEventHandler struct {
	next fsmpkg.EventHandler
}

var _ fsmpkg.EventHandler = (*LoginEventHandler)(nil)

func NewLoginEventHandler(next fsmpkg.EventHandler) *LoginEventHandler {
	return &LoginEventHandler{
		next: next,
	}
}

func (h *LoginEventHandler) HandleEvent(ctx shared.Context, fsm fsmpkg.FiniteStateMachine, event fsmpkg.Event) {
	if event.GetEventType() == botevent.LoginEvent {
		payload := event.GetEventPayload().(botevent.LoginPayload)
		ctx.SetValue(fields.UserCount, ctx.GetValue(fields.UserCount).(int)+1)
		if payload.IsAdmin {
			ctx.SetValue(fields.UserList, append(ctx.GetValue(fields.UserList).([]*shared.Member), shared.NewAdminMember(payload.UserID)))
		} else {
			ctx.SetValue(fields.UserList, append(ctx.GetValue(fields.UserList).([]*shared.Member), shared.NewMember(payload.UserID)))
		}
	} else if h.next != nil {
		h.next.HandleEvent(ctx, fsm, event)
	}
}
