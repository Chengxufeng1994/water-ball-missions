package knowledgeking

import (
	"fmt"

	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/bot/fields"
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/fsm"
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/fsm/state"
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/shared"
)

type ThanksForJoining struct {
	*state.BaseState
}

func NewThanksForJoining(entryAction, exitAction fsm.Action, botFsmAdapter state.BotFsmAdapter) *ThanksForJoining {
	return &ThanksForJoining{
		BaseState: state.NewBaseState("thanks_for_joining", entryAction, exitAction, botFsmAdapter),
	}
}

func (state *ThanksForJoining) EntryState(ctx shared.Context, event fsm.Event) {
	state.BaseState.EntryAction.Execute()
	state.ProcessEntryState(ctx)
}

func (state *ThanksForJoining) ProcessEntryState(ctx shared.Context) {
	winnerId, ok := ctx.GetValue(fields.WinnerId).(string)
	var message string

	switch {
	case ok && winnerId != "":
		message = fmt.Sprintf("The winner is @%s", winnerId)
	case !ok || winnerId == "":
		message = "Tie!"
	default:
		message = "⚠️ Unable to determine winner."
	}

	_, ok = ctx.GetValue(fields.SpeakerId).(string)
	if !ok {
		state.Adapter.Speak(message)
	} else {
		state.Adapter.SendMessage(message)
	}
}
