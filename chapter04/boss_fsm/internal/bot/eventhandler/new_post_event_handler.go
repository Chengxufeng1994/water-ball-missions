package eventhandler

import (
	"fmt"

	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/shared"
)

type NewPostEventHandler struct {
	next shared.EventHandler
}

func NewNewPostEventHandler(next shared.EventHandler) *NewPostEventHandler {
	return &NewPostEventHandler{
		next: next,
	}
}

func (h *NewPostEventHandler) HandleEvent(event *shared.Event, ctx shared.Context) {
	if event.Type == shared.NewPostEvent {
		newPost := event.Payload.(shared.NewPostPayload)
		message := ctx.GetState().GenerateMessage("forum", newPost.AuthorID)
		fmt.Printf("%s comment in post %s: %s\n", ctx.GetPrefix(), newPost.PostID, message)
		// TODO: send event to notify Community to save comment
		return
	} else if h.next != nil {
		h.next.HandleEvent(event, ctx)
	}
}
