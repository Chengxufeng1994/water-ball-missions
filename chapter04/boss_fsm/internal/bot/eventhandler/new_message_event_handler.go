package eventhandler

import (
	"fmt"

	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/shared"
)

type NewMessageEventHandler struct {
	next shared.EventHandler
}

var _ shared.EventHandler = (*NewMessageEventHandler)(nil)

func NewNewMessageEventHandler(next shared.EventHandler) *NewMessageEventHandler {
	return &NewMessageEventHandler{
		next: next,
	}
}

func (h *NewMessageEventHandler) HandleEvent(event *shared.Event, ctx shared.Context) {
	if event.Type == shared.NewMessageEvent {
		newMessage := event.Payload.(shared.NewMessagePayload)
		message := ctx.GetState().GenerateMessage("chat", newMessage.AuthorID)
		for _, tag := range newMessage.Tags {
			if tag == "bot" {
				// TODO: check message is command
			}
		}
		fmt.Printf("%s: %s @%s\n", ctx.GetPrefix(), message, newMessage.AuthorID)
	} else if h.next != nil {
		h.next.HandleEvent(event, ctx)
	}
}
