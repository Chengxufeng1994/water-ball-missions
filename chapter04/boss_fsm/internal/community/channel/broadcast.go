package channel

import (
	"fmt"

	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/community/content"
)

type Broadcast struct {
	InUseUser string
	Prefix    string
}

var _ Channel = (*Broadcast)(nil)

func NewBroadcast() *Broadcast {
	return &Broadcast{
		InUseUser: "",
		Prefix:    "📢",
	}
}

func (b *Broadcast) SendContent(content content.CommunityContent) error {
	fmt.Printf("%s %s: %s\n", b.Prefix, content.GetAuthorID(), content.GetContent())
	return nil
}

func (b *Broadcast) StartBroadcasting(speakId string) {
	b.InUseUser = speakId
	fmt.Printf("%s %s is broadcasting...\n", b.Prefix, b.InUseUser)
}

func (b *Broadcast) StopBroadcasting(speakId string) {
	if b.InUseUser != speakId {
		return
	}
	b.InUseUser = ""
	fmt.Printf("%s %s stop broadcasting\n", b.Prefix, speakId)
}

func (b *Broadcast) IsBroadcasting() bool {
	return b.InUseUser != ""
}
