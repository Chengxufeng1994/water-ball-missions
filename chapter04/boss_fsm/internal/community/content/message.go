package content

import "github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/community/tag"

type Message struct {
	*BasedContent
}

var _ CommunityContent = (*Message)(nil)

func NewMessage(authorID string, content string, tags []tag.Tag) *Message {
	return &Message{
		BasedContent: NewBasedContent(authorID, content, tags),
	}
}
