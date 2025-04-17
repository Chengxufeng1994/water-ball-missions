package normal

import (
	botevent "github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/bot/event"
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

func NewDefaultConversation(entryAction, exitAction fsm.Action, xxx state.BotFsmAdapter) *DefaultConversation {
	dc := &DefaultConversation{
		BaseState: state.NewBaseState("default_conversation", entryAction, exitAction, xxx),
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

func (state *DefaultConversation) EntryState(ctx shared.Context, event fsm.Event) {
	state.BaseState.EntryAction.Execute()
	state.ProcessEntryState(ctx)
}

func (state *DefaultConversation) ProcessEntryState(ctx shared.Context) {
	state.Index = 0
}

func (state *DefaultConversation) OnEvent(ctx shared.Context, evt fsm.Event) fsm.Event {
	switch evt.GetEventType() {
	case botevent.NewMessageEvent:
		message := state.Messages["chat"][state.Index]
		authorId := evt.GetEventPayload().(botevent.NewMessagePayload).AuthorID
		state.Index = (state.Index + 1) % len(state.Messages["chat"])
		state.Adapter.SendMessage(message, authorId)
	case botevent.NewPostEvent:
		authorId := evt.GetEventPayload().(botevent.NewPostPayload).AuthorID
		state.Adapter.SendMessage(state.Messages["forum"][0], authorId)
	}
	return nil
}
