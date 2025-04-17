package bot

import (
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/fsm"
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/shared"
)

type BotContext struct {
	Prefix string
	Values map[string]any
	Quota  Quota
	FSM    shared.FiniteStateMachine
}

func (b *BotContext) GetPrefix() string {
	return b.Prefix
}

func (b *BotContext) GetState() shared.State {
	return b.FSM.GetCurrentState()
}

func (b *BotContext) GetFSM() shared.FiniteStateMachine {
	return b.FSM
}

func (b *BotContext) GetValue(key string) any {
	return b.Values[key]
}

func (b *BotContext) SetValue(key string, value any) {
	b.Values[key] = value
}

func initializeBotContext(quota Quota, fsm *fsm.FiniteStateMachine) *BotContext {
	values := make(map[string]any)
	// Online
	values[Users] = 1
	userList := make([]string, 0)
	userList = append(userList, "bot")
	values[UsersList] = userList
	botCtx := &BotContext{
		Prefix: "🤖",
		Values: values,
		Quota:  quota,
		FSM:    fsm,
	}

	return botCtx
}

type Bot struct {
	EventHandler shared.EventHandler
	Ctx          *BotContext
}

var _ interface {
	shared.EventListener
} = (*Bot)(nil)

func NewBot(quota Quota, eventHandler shared.EventHandler, fsm *fsm.FiniteStateMachine) *Bot {
	return &Bot{
		EventHandler: eventHandler,
		Ctx:          initializeBotContext(quota, fsm),
	}
}

func (b *Bot) OnEvent(event *shared.Event) {
	b.EventHandler.HandleEvent(event, b.Ctx)
}
