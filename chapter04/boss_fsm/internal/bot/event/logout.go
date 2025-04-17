package event

import "github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/fsm"

type LogoutPayload struct {
	UserID string `json:"userId"`
}

func NewLogoutEvent(payload LogoutPayload) fsm.Event {
	return fsm.NewEvent(LogoutEvent, payload)
}
