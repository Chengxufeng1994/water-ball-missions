package action

import (
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/fsm"
)

type LoggerAction struct {
	Message string
}

var _ fsm.Action = (*LoggerAction)(nil)

func NewLoggerAction(message string) *LoggerAction {
	return &LoggerAction{Message: message}
}

func (action *LoggerAction) Execute() {
	// fmt.Println("[LoggerAction]", action.Message)
}
