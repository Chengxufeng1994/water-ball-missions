package content

import "github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/community/tag"

type BasedContent struct {
	authorID string
	content  string
	tags     []tag.Tag
}

var _ CommunityContent = (*BasedContent)(nil)

func NewBasedContent(authorID string, content string, tags []tag.Tag) *BasedContent {
	return &BasedContent{
		authorID: authorID,
		content:  content,
		tags:     tags,
	}
}

func (b *BasedContent) GetAuthorID() string {
	return b.authorID
}

func (b *BasedContent) GetContent() string {
	return b.content
}

func (b *BasedContent) GetTags() []tag.Tag {
	return b.tags
}
