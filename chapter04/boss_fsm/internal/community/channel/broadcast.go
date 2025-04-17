package channel

import (
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/community/content"
)

type Broadcast struct {
}

var _ Channel = (*Broadcast)(nil)

func NewBroadcast() *Broadcast {
	return &Broadcast{}
}

func (b *Broadcast) SendContent(content content.CommunityContent) error {
	panic("unimplemented")
}
