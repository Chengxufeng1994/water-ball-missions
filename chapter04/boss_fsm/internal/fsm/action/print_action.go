package action

import (
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/fsm"
)

type PrintAction struct {
	Message string
}

var _ fsm.Action = (*PrintAction)(nil)

func NewPrintAction(message string) *PrintAction {
	return &PrintAction{
		Message: message,
	}
}

func (action *PrintAction) Execute() {
	// fmt.Println("[print action]", action.Message)
}
