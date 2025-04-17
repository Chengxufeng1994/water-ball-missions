package channel

import (
	"fmt"
	"strings"

	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/community/content"
)

type Forum struct {
	Posts []*content.Post
}

var _ Channel = (*Forum)(nil)

func NewForum() *Forum {
	return &Forum{
		Posts: make([]*content.Post, 0),
	}
}

func (f *Forum) SendContent(communityContent content.CommunityContent) error {
	tags := make([]string, len(communityContent.GetTags()))
	for i, tag := range communityContent.GetTags() {
		tags[i] = "@" + tag.UserID
	}
	post, ok := communityContent.(*content.Post)
	if !ok {
		err := fmt.Errorf("communityContent is not of type *content.Post")
		return err
	}
	f.Posts = append(f.Posts, post)

	fmt.Printf("%s: [%s] %s %s\n",
		post.GetAuthorID(),
		post.GetTitle(),
		post.GetContent(),
		strings.Join(tags, ", "),
	)
	return nil
}
