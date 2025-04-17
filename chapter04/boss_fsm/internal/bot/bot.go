package bot

import (
	"fmt"
	"strings"

	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/bot/eventhandler"
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/bot/recordreplayformatstrategy"
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/fsm"
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/fsm/state"
)

type Bot interface {
	GetFSM() fsm.FiniteStateMachine
}

type CommunityBot struct {
	Prefix                     string
	Ctx                        *BotContext
	Fsm                        fsm.FiniteStateMachine
	EventHandler               fsm.EventHandler
	RecordReplayFormatStrategy recordreplayformatstrategy.RecordReplayFormatStrategy
}

var _ interface {
	Bot
	fsm.EventListener
	state.BotFsmAdapter
} = (*CommunityBot)(nil)

func NewBot(quota int) *CommunityBot {
	eventHandler := eventhandler.NewLoginEventHandler(
		eventhandler.NewLogoutEventHandler(
			eventhandler.NewGoBroadcastingEventHandler(nil),
		),
	)
	ctx := NewBotContext(quota)
	bot := &CommunityBot{
		Prefix:                     "🤖",
		Ctx:                        ctx,
		EventHandler:               eventHandler,
		RecordReplayFormatStrategy: recordreplayformatstrategy.NewDefaultRecordReplayFormat(),
	}

	fsm := NewBotFSM(bot)
	bot.Fsm = fsm
	bot.Fsm.Initialize(bot.Ctx)

	return bot
}

func (b *CommunityBot) OnEvent(event fsm.Event) {
	b.EventHandler.HandleEvent(b.Ctx, b.Fsm, event)
	b.Fsm.HandleEvent(b.Ctx, event)
}

func (b *CommunityBot) GetFSM() fsm.FiniteStateMachine {
	return b.Fsm
}

func (b *CommunityBot) SendMessage(message string, authorIds ...string) {
	for i, authorId := range authorIds {
		authorIds[i] = "@" + authorId
	}

	fmt.Printf("%s: %s %s\n", b.Prefix, message, strings.Join(authorIds, ", "))
}

func (b *CommunityBot) Replay(message string) {
	format := b.RecordReplayFormatStrategy.Format()
	fmt.Printf("%s: %s\n", b.Prefix, fmt.Sprintf(format, message))
}

func (b *CommunityBot) Speak(message string) {
	fmt.Printf("%s speaking: %s\n", b.Prefix, message)
}
