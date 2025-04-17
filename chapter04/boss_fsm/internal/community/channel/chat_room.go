package channel

import (
	"fmt"
	"strings"

	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/community/content"
)

type Chatroom struct {
	Message []*content.Message
	Prefix  string
}

var _ Channel = (*Chatroom)(nil)

func NewChatRoom() *Chatroom {
	return &Chatroom{
		Message: make([]*content.Message, 0),
		Prefix:  "💬",
	}
}

func (c *Chatroom) SendContent(communityContent content.CommunityContent) error {
	tags := make([]string, len(communityContent.GetTags()))
	for i, tag := range communityContent.GetTags() {
		tags[i] = "@" + tag.UserID
	}
	message, ok := communityContent.(*content.Message)
	if !ok {
		err := fmt.Errorf("communityContent is not of type *content.Message")
		return err
	}
	c.Message = append(c.Message, message)
	fmt.Printf("%s %s: %s %s\n",
		c.Prefix,
		communityContent.GetAuthorID(),
		communityContent.GetContent(),
		strings.Join(tags, ", "),
	)
	return nil
}
