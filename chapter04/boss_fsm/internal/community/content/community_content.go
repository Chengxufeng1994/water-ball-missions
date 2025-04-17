package content

import "github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/community/tag"

type CommunityContent interface {
	GetAuthorID() string
	GetContent() string
	GetTags() []tag.Tag
}
