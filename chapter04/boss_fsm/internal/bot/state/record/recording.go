package record

import (
	"fmt"
	"strings"

	botevent "github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/bot/event"
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/bot/fields"
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/fsm"
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/fsm/state"
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/shared"
)

type Recording struct {
	*state.BaseState

	Contents []string
}

var _ fsm.State = (*Recording)(nil)

func NewRecording(entryAction, exitAction fsm.Action, xxx state.BotFsmAdapter) *Recording {
	return &Recording{
		BaseState: state.NewBaseState("recording", entryAction, exitAction, xxx),
	}
}

func (r *Recording) EntryState(ctx shared.Context, event fsm.Event) {
	r.BaseState.EntryAction.Execute()
	r.ProcessEntryState(ctx)
}

func (r *Recording) ProcessEntryState(ctx shared.Context) {
	r.Contents = make([]string, 0)
}

func (r *Recording) OnEvent(ctx shared.Context, event fsm.Event) fsm.Event {
	if event.GetEventType() == botevent.SpeakEvent {
		payload, _ := event.GetEventPayload().(botevent.SpeakPayload)
		r.Contents = append(r.Contents, payload.Content)
	}
	return nil
}

func (r *Recording) ExitState(ctx shared.Context, event fsm.Event) {
	r.BaseState.ExitAction.Execute()
	r.ProcessExitState(ctx)
}

func (r *Recording) ProcessExitState(ctx shared.Context) {
	ctx.Del(fields.SpeakerId)

	var sb strings.Builder
	for i, content := range r.Contents {
		sb.WriteString(content)
		if i == len(r.Contents)-1 {
			continue
		}
		sb.WriteString("\n")
	}

	sb.WriteString(" ")
	sb.WriteString(fmt.Sprintf("@%s", ctx.GetValue(fields.RecorderId).(string)))
	r.Adapter.Replay(sb.String())
}
