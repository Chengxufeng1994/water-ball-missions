package eventhandler

import (
	botevent "github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/bot/event"
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/bot/fields"
	fsmpkg "github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/fsm"
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/shared"
)

type GoBroadcastingEventHandler struct {
	next fsmpkg.EventHandler
}

var _ fsmpkg.EventHandler = (*GoBroadcastingEventHandler)(nil)

func NewGoBroadcastingEventHandler(next fsmpkg.EventHandler) *GoBroadcastingEventHandler {
	return &GoBroadcastingEventHandler{
		next: next,
	}
}

func (g *GoBroadcastingEventHandler) HandleEvent(ctx shared.Context, finiteStateMachine fsmpkg.FiniteStateMachine, event fsmpkg.Event) {
	if event.GetEventType() == botevent.GoBroadcastingEvent {
		payload, _ := event.GetEventPayload().(botevent.GoBroadcastingPayload)
		ctx.SetValue(fields.SpeakerId, payload.SpeakerID)
	} else if g.next != nil {
		g.next.HandleEvent(ctx, finiteStateMachine, event)
	}
}
