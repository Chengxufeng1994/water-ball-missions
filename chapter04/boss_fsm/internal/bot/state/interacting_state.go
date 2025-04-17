package state

import (
	"fmt"
	"strings"

	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/bot"
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/fsm"
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/shared"
)

type Interacting struct {
	*BaseState
	Index    int
	Messages map[string][]string
}

var _ shared.State = (*Interacting)(nil)

func NewInteractingState(entryAction, exitAction fsm.Action) *Interacting {
	return &Interacting{
		BaseState: NewBaseState(entryAction, exitAction),
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
func (is *Interacting) EntryState() {
	is.Index = 0
}

func (is *Interacting) GenerateMessage(channelType, authorId string) string {
	switch channelType {
	case "chat":
		message := is.Messages["chat"][is.Index]
		is.Index = (is.Index + 1) % len(is.Messages["chat"])
		return message
	case "forum":
		userList := is.Ctx.GetValue(bot.UsersList).([]string)
		tags := make([]string, len(userList))
		for i := range tags {
			tags[i] = "@" + userList[i]
		}
		return fmt.Sprintf("%s %s", is.Messages["forum"][0], strings.Join(tags, ", "))
	}

	return ""
}
