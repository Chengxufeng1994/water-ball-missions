package eventhandler

import (
	"fmt"

	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/fsm"
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/shared"
)

type NewMessageEventHandler struct {
	next fsm.EventHandler
}

var _ fsm.EventHandler = (*NewMessageEventHandler)(nil)

func NewNewMessageEventHandler(next fsm.EventHandler) *NewMessageEventHandler {
	return &NewMessageEventHandler{
		next: next,
	}
}

func (h *NewMessageEventHandler) HandleEvent(ctx shared.Context, finiteStateMachine fsm.FiniteStateMachine, event *fsm.BaseEvent) {
	if event.Type == fsm.NewMessageEvent {
		newMessage := event.Payload.(fsm.NewMessagePayload)
		message := finiteStateMachine.GetCurrentState().GenerateMessage(event)
		fmt.Printf("%s: %s @%s\n", ctx.GetPrefix(), message, newMessage.AuthorID)
	} else if h.next != nil {
		h.next.HandleEvent(ctx, finiteStateMachine, event)
	}
}
