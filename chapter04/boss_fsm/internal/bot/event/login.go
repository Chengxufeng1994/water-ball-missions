package event

import "github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/fsm"

type LoginPayload struct {
	UserID  string `json:"userId"`
	IsAdmin bool   `json:"isAdmin"`
}

func NewLoginEvent(payload LoginPayload) fsm.Event {
	return fsm.NewEvent(LoginEvent, payload)
}
