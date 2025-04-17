package bot

import (
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/bot/fields"
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/shared"
)

type BotContext struct {
	Prefix string
	Values map[string]any
}

var _ shared.Context = (*BotContext)(nil)

func NewBotContext(quota int) *BotContext {
	values := make(map[string]any)
	values[fields.UserCount] = 1
	userList := make([]*shared.Member, 0)
	userList = append(userList, shared.NewAdminMember("bot"))
	values[fields.UserList] = userList
	values[fields.Quota] = quota

	return &BotContext{
		Values: values,
	}
}

func (b *BotContext) GetValue(key string) any {
	val, ok := b.Values[key]
	if !ok {
		return nil
	}
	return val
}

func (b *BotContext) SetValue(key string, value any) {
	b.Values[key] = value
}

func (b *BotContext) Del(key string) {
	delete(b.Values, key)
}
