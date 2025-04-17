package channel

import (
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/community/content"
)

type Channel interface {
	SendContent(content content.CommunityContent) error
}
