package event

import "github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/fsm"

type NewMessagePayload struct {
	AuthorID string   `json:"authorId"`
	Content  string   `json:"content"`
	Tags     []string `json:"tags"`
}

func NewNewMessageEvent(payload NewMessagePayload) fsm.Event {
	return fsm.NewEvent(NewMessageEvent, payload)
}
