package normal

import (
	"fmt"

	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/fsm"
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/fsm/state"
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/shared"
)

type DefaultConversation struct {
	*state.BaseState
	Index    int
	Messages map[string][]string
}

var _ fsm.State = (*DefaultConversation)(nil)

func NewDefaultConversation(entryAction, exitAction fsm.Action) *DefaultConversation {
	dc := &DefaultConversation{
		BaseState: state.NewBaseState("default_conversation", entryAction, exitAction),
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

	return dc
}

func (ns *DefaultConversation) EntryState(ctx shared.Context, event fsm.Event) {
	ns.BaseState.EntryAction.Execute()
	ns.ProcessEntryState(ctx)
}

func (ns *DefaultConversation) ProcessEntryState(ctx shared.Context) {
	ns.Index = 0
}

func (ns *DefaultConversation) GenerateMessage(evt fsm.Event) string {
	switch evt.GetEventType() {
	case fsm.NewMessageEvent:
		message := ns.Messages["chat"][ns.Index]
		ns.Index = (ns.Index + 1) % len(ns.Messages["chat"])
		return message
	case fsm.NewPostEvent:
		authorId := evt.GetEventPayload().(fsm.NewPostPayload).AuthorID
		return fmt.Sprintf("%s %s", ns.Messages["forum"][0], "@"+authorId)
	}

	return ""
}
