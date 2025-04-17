package content

import "github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/community/tag"

type Post struct {
	*BasedContent
	ID      string
	Title   string
	Comment []*CommunityContent
}

var _ CommunityContent = (*Post)(nil)

func NewPost(id string, authorID string, title string, content string, tags []tag.Tag) *Post {
	return &Post{
		BasedContent: NewBasedContent(authorID, content, tags),
		ID:           id,
		Title:        title,
	}
}

func (p *Post) GetID() string {
	return p.ID
}

func (p *Post) GetTitle() string {
	return p.Title
}

func (p *Post) AddComment(comment *CommunityContent) {
	p.Comment = append(p.Comment, comment)
}
