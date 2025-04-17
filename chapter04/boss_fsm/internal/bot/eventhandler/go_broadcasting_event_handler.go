package eventhandler

import (
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/bot/fields"
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/fsm"
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/shared"
)

type GoBroadcastingEventHandler struct {
	next fsm.EventHandler
}

var _ fsm.EventHandler = (*GoBroadcastingEventHandler)(nil)

func NewGoBroadcastingEventHandler(next fsm.EventHandler) *GoBroadcastingEventHandler {
	return &GoBroadcastingEventHandler{
		next: next,
	}
}

func (g *GoBroadcastingEventHandler) HandleEvent(ctx shared.Context, finiteStateMachine fsm.FiniteStateMachine, event *fsm.BaseEvent) {
	if event.Type == fsm.GoBroadcastingEvent {
		payload, _ := event.Payload.(fsm.GoBroadcastingPayload)
		ctx.SetValue(fields.SpeakerId, payload.SpeakerID)
	} else if g.next != nil {
		g.next.HandleEvent(ctx, finiteStateMachine, event)
	}
}
