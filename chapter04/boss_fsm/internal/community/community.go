package community

import (
	"slices"

	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/bot"
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/bot/eventhandler"
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/bot/state"
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/community/channel"
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/community/content"
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/community/tag"
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/fsm"
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/fsm/guard"
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/shared"
)

type Community struct {
	Name      string
	Member    []*shared.Member
	ChatRoom  *channel.Chatroom  // Channel
	Forum     *channel.Forum     // Channel
	Broadcast *channel.Broadcast // Channel
	Bot       *bot.Bot
}

func NewCommunity(name string) *Community {
	return &Community{
		Name:      name,
		Member:    make([]*shared.Member, 0),
		ChatRoom:  channel.NewChatRoom(),
		Forum:     channel.NewForum(),
		Broadcast: channel.NewBroadcast(),
	}
}

func (c *Community) HandleCommunityEvent(event CommunityEvent) error {
	switch event.EventName() {
	case StartedEvent:
		payload, ok := event.Payload().(StartedEventPayload)
		if !ok {
			return ErrInvalidPayload
		}
		c.Started(payload)
		return nil
	case LoginEvent:
		payload, ok := event.Payload().(LoginEventPayload)
		if !ok {
			return ErrInvalidPayload
		}
		var member *shared.Member
		if payload.IsAdmin {
			member = shared.NewAdminMember(payload.UserID)
		} else {
			member = shared.NewMember(payload.UserID)
		}
		c.Login(member)
		return nil
	case LogoutEvent:
		payload, ok := event.Payload().(LogoutEventPayload)
		if !ok {
			return ErrInvalidPayload
		}
		c.Logout(payload.UserID)
		return nil
	case NewMessageEvent:
		payload, ok := event.Payload().(NewMessageEventPayload)
		if !ok {
			return ErrInvalidPayload
		}
		tags := make([]tag.Tag, len(payload.Tags))
		for i, ptag := range payload.Tags {
			tags[i] = tag.Tag{
				UserID: ptag,
			}
		}
		newMessage := content.NewMessage(payload.AuthorID, payload.Content, tags)
		c.NewMessage(newMessage)
		return nil
	case NewPostEvent:
		payload, ok := event.Payload().(NewPostEventPayload)
		if !ok {
			return ErrInvalidPayload
		}
		tags := make([]tag.Tag, len(payload.Tags))
		for i, ptag := range payload.Tags {
			tags[i] = tag.Tag{
				UserID: ptag.UserID,
			}
		}
		newPost := content.NewPost(payload.PostID, payload.AuthorID, payload.Title, payload.Content, tags)
		c.NewPost(newPost)
		return nil
	}

	return nil
}

func (c *Community) Started(payload StartedEventPayload) {
	defaultConversation := state.NewDefaultConversation(nil, nil)
	interacting := state.NewInteractingState(nil, nil)
	fsm := fsm.NewFiniteStateMachine(defaultConversation, []*fsm.Transition{
		fsm.NewTransition(defaultConversation, interacting, shared.LoginEvent, guard.NewIsGreaterThan("Users", 10), nil),
		fsm.NewTransition(interacting, defaultConversation, shared.LogoutEvent, guard.NewIsLessThan("Users", 10), nil),
	})
	eventHandlers := eventhandler.NewLoginEventHandler(
		eventhandler.NewLogoutEventHandler(
			eventhandler.NewNewMessageEventHandler(
				eventhandler.NewNewPostEventHandler(nil),
			),
		),
	)
	quota := bot.NewQuota(payload.Quota)
	bot := bot.NewBot(quota, eventHandlers, fsm)
	c.Bot = bot
}

func (c *Community) Login(member *shared.Member) {
	c.Member = append(c.Member, member)
	c.NotifyBot(shared.NewLoginEvent(shared.LoginPayload{
		UserID:  member.UserID,
		IsAdmin: member.IsAdmin,
	}))
}

func (c *Community) Logout(memberId string) {
	for i, m := range c.Member {
		if m.UserID == memberId {
			c.Member = slices.Delete(c.Member, i, i+1)
			c.NotifyBot(shared.NewLogoutEvent(shared.LogoutPayload{
				UserID: memberId,
			}))
		}
	}
}

func (c *Community) NewMessage(message *content.Message) {
	c.ChatRoom.SendContent(message)

	domainTags := message.GetTags()
	tags := make([]string, len(domainTags))
	for i, tag := range domainTags {
		tags[i] = tag.UserID
	}
	c.NotifyBot(shared.NewNewMessageEvent(shared.NewMessagePayload{
		AuthorID: message.GetAuthorID(),
		Content:  message.GetContent(),
		Tags:     tags,
	}))
}

func (c *Community) NewPost(post *content.Post) {
	c.Forum.SendContent(post)

	domainTags := post.GetTags()
	tags := make([]string, len(domainTags))
	for i, tag := range domainTags {
		tags[i] = tag.UserID
	}
	c.NotifyBot(shared.NewNewPostEvent(shared.NewPostPayload{
		PostID:   post.GetID(),
		AuthorID: post.GetAuthorID(),
		Title:    post.GetTitle(),
		Content:  post.GetContent(),
		Tags:     tags,
	}))
}

func (c *Community) NotifyBot(event *shared.Event) {
	c.Bot.OnEvent(event)
}
