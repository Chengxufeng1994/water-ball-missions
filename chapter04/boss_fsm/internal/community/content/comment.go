package content

import "github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/community/tag"

type Comment struct {
	*BasedContent
}

func NewComment(authorID string, content string, tags []tag.Tag) *Comment {
	return &Comment{
		BasedContent: NewBasedContent(authorID, content, tags),
	}
}
