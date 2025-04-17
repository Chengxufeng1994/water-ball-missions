package eventhandler

import (
	"fmt"

	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/fsm"
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/shared"
)

type NewPostEventHandler struct {
	next fsm.EventHandler
}

func NewNewPostEventHandler(next fsm.EventHandler) *NewPostEventHandler {
	return &NewPostEventHandler{
		next: next,
	}
}

func (h *NewPostEventHandler) HandleEvent(ctx shared.Context, finiteStateMachine fsm.FiniteStateMachine, event *fsm.BaseEvent) {
	if event.Type == fsm.NewPostEvent {
		newPost := event.Payload.(fsm.NewPostPayload)
		message := finiteStateMachine.GetCurrentState().GenerateMessage(event)
		fmt.Printf("%s comment in post %s: %s\n", ctx.GetPrefix(), newPost.PostID, message)
		// TODO: send event to notify Community to save comment
	} else if h.next != nil {
		h.next.HandleEvent(ctx, finiteStateMachine, event)
	}
}
