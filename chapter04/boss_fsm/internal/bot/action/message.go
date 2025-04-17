package action

import (
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/fsm"
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/shared"
)

type KnowledgeStarting struct {
	bot     shared.Messenger
	message string
}

var _ fsm.Action = (*KnowledgeStarting)(nil)

func NewSendMessageAction(bot shared.Messenger, message string) *KnowledgeStarting {
	return &KnowledgeStarting{
		bot:     bot,
		message: message,
	}
}

func (action *KnowledgeStarting) Execute() {
	action.bot.SendMessage(action.message)
}
