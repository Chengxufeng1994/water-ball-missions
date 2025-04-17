package normal

import (
	botevent "github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/bot/event"
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/bot/fields"
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/fsm"
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/fsm/state"
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/shared"
)

type Interacting struct {
	*state.BaseState
	Index    int
	Messages map[string][]string
}

var _ fsm.State = (*Interacting)(nil)

func NewInteractingState(entryAction, exitAction fsm.Action, xxx state.BotFsmAdapter) *Interacting {
	return &Interacting{
		BaseState: state.NewBaseState("interacting", entryAction, exitAction, xxx),
		Index:     0,
		Messages: map[string][]string{
			"chat": {
				"Hi hi😁",
				"I like your idea!",
			},
			"forum": {
				"How do you guys think about it?",
			},
		},
	}
}

func (state *Interacting) EntryState(ctx shared.Context, event fsm.Event) {
	state.BaseState.EntryAction.Execute()
	state.ProcessEntryState(ctx)
}

func (state *Interacting) ProcessEntryState(ctx shared.Context) {
	state.Index = 0
}

func (state *Interacting) OnEvent(ctx shared.Context, evt fsm.Event) fsm.Event {
	switch evt.GetEventType() {
	case botevent.NewMessageEvent:
		message := state.Messages["chat"][state.Index]
		authorId := evt.GetEventPayload().(botevent.NewMessagePayload).AuthorID
		state.Index = (state.Index + 1) % len(state.Messages["chat"])
		state.Adapter.SendMessage(message, authorId)
	case botevent.NewPostEvent:
		members := state.Ctx.GetValue(fields.UserList).([]*shared.Member)
		userList := make([]string, 0, len(members))
		for _, member := range members {
			userList = append(userList, member.UserID)
		}
		state.Adapter.SendMessage(state.Messages["forum"][0], userList...)
	}
	return nil
}
