package bot

import (
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/fsm"
)

type Bot interface {
	GetFSM() fsm.FiniteStateMachine
}

type CommunityBot struct {
	Ctx          *BotContext
	Fsm          fsm.FiniteStateMachine
	EventHandler fsm.EventHandler
}

var _ interface {
	Bot
	fsm.EventListener
} = (*CommunityBot)(nil)

func NewBot(quota int, eventHandler fsm.EventHandler) *CommunityBot {
	ctx := NewBotContext(quota)
	fsm := NewBotFSM()

	bot := &CommunityBot{
		Ctx:          ctx,
		Fsm:          fsm,
		EventHandler: eventHandler,
	}

	bot.Fsm.Initialize(bot.Ctx)

	return bot
}

func (b *CommunityBot) OnEvent(event *fsm.BaseEvent) {
	b.EventHandler.HandleEvent(b.Ctx, b.Fsm, event)
	b.Fsm.HandleEvent(b.Ctx, event)
}

func (b *CommunityBot) GetFSM() fsm.FiniteStateMachine {
	return b.Fsm
}
