package event

import "github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/fsm"

type NewPostPayload struct {
	PostID   string   `json:"id"`
	AuthorID string   `json:"authorId"`
	Title    string   `json:"title"`
	Content  string   `json:"content"`
	Tags     []string `json:"tags"`
}

func NewNewPostEvent(payload NewPostPayload) fsm.Event {
	return fsm.NewEvent(NewPostEvent, payload)
}
