package normal

import (
	"fmt"
	"strings"

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

func NewInteractingState(entryAction, exitAction fsm.Action) *Interacting {
	return &Interacting{
		BaseState: state.NewBaseState("interacting", entryAction, exitAction),
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

func (is *Interacting) EntryState(ctx shared.Context, event fsm.Event) {
	is.BaseState.EntryAction.Execute()
	is.ProcessEntryState(ctx)
}

func (is *Interacting) ProcessEntryState(ctx shared.Context) {
	is.Index = 0
}

func (is *Interacting) GenerateMessage(evt fsm.Event) string {
	switch evt.GetEventType() {
	case fsm.NewMessageEvent:
		message := is.Messages["chat"][is.Index]
		is.Index = (is.Index + 1) % len(is.Messages["chat"])
		return message
	case fsm.NewPostEvent:
		userList := is.Ctx.GetValue(fields.UsersList).([]string)
		tags := make([]string, len(userList))
		for i := range tags {
			tags[i] = "@" + userList[i]
		}
		return fmt.Sprintf("%s %s", is.Messages["forum"][0], strings.Join(tags, ", "))
	}

	return ""
}
