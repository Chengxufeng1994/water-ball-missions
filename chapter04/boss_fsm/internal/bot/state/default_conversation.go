package state

import (
	"fmt"

	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/fsm"
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/shared"
)

type DefaultConversation struct {
	*BaseState
	Index    int
	Messages map[string][]string
}

var _ shared.State = (*DefaultConversation)(nil)

func NewDefaultConversation(entryAction, exitAction fsm.Action) *DefaultConversation {
	return &DefaultConversation{
		BaseState: NewBaseState(entryAction, exitAction),
		Index:     0,
		Messages: map[string][]string{
			"chat": {
				"good to hear",
				"thank you",
				"How are you",
			},
			"forum": {
				"Nice post",
			},
		},
	}
}

func (ns *DefaultConversation) EntryState() {
	ns.Index = 0
}

func (ns *DefaultConversation) GenerateMessage(channelType, authorId string) string {
	switch channelType {
	case "chat":
		message := ns.Messages["chat"][ns.Index]
		ns.Index = (ns.Index + 1) % len(ns.Messages["chat"])
		return message
	case "forum":
		return fmt.Sprintf("%s %s", ns.Messages["forum"][0], "@"+authorId)
	}

	return ""
}

func (ns *DefaultConversation) ExitState() {
	// do nothing
}
