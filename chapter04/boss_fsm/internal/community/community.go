package community

import (
	"slices"

	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/bot"
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/bot/eventhandler"
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/community/channel"
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/community/content"
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/community/tag"
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/fsm"
)

type Community struct {
	Name           string
	Members        []*Member
	ChatRoom       *channel.Chatroom  // Channel
	Forum          *channel.Forum     // Channel
	Broadcast      *channel.Broadcast // Channel
	Bot            *bot.CommunityBot
	EventListeners []fsm.EventListener
}

func NewCommunity(name string) *Community {
	return &Community{
		Name:      name,
		Members:   make([]*Member, 0),
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
		var member *Member
		if payload.IsAdmin {
			member = NewAdminMember(payload.UserID)
		} else {
			member = NewMember(payload.UserID)
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
	case GoBroadcastingEvent:
		payload, ok := event.Payload().(GoBroadcastingEventPayload)
		if !ok {
			return ErrInvalidPayload
		}
		c.GoBroadcasting(payload.SpeakerID)
		return nil
	case StopBroadcastingEvent:
		payload, ok := event.Payload().(StopBroadcastingEventPayload)
		if !ok {
			return ErrInvalidPayload
		}
		c.StopBroadcasting(payload.SpeakerID)
		return nil
	case SpeakEvent:
		payload, ok := event.Payload().(SpeakEventPayload)
		if !ok {
			return ErrInvalidPayload
		}
		speakMessage := content.NewSpeak(payload.SpeakerID, payload.Content)
		c.Speak(speakMessage)
		return nil
	case EndEvent:
		return nil
	default:
		return ErrCommunityEvent
	}
}

func (c *Community) Started(payload StartedEventPayload) {
	eventHandlers := eventhandler.NewLoginEventHandler(
		eventhandler.NewLogoutEventHandler(
			eventhandler.NewNewMessageEventHandler(
				eventhandler.NewNewPostEventHandler(
					eventhandler.NewGoBroadcastingEventHandler(nil),
				),
			),
		),
	)

	bot := bot.NewBot(payload.Quota, eventHandlers)
	c.EventListeners = append(c.EventListeners, bot)
	c.Members = append(c.Members, NewAdminMember("bot"))
}

func (c *Community) Login(member *Member) {
	c.Members = append(c.Members, member)
	c.NotifyEvent(fsm.NewLoginEvent(fsm.LoginPayload{
		UserID:  member.UserID,
		IsAdmin: member.IsAdmin,
	}))
}

func (c *Community) Logout(memberId string) {
	for i, m := range c.Members {
		if m.UserID == memberId {
			c.Members = slices.Delete(c.Members, i, i+1)
			c.NotifyEvent(fsm.NewLogoutEvent(fsm.LogoutPayload{
				UserID: memberId,
			}))
		}
	}
}

func (c *Community) NewMessage(message *content.Message) {
	c.ChatRoom.SendContent(message)
	c.NotifyEvent(fsm.NewNewMessageEvent(fsm.NewMessagePayload{
		AuthorID: message.GetAuthorID(),
		Content:  message.GetContent(),
		Tags:     c._toTags(message.GetTags()),
	}))
}

func (c *Community) NewPost(post *content.Post) {
	c.Forum.SendContent(post)
	c.NotifyEvent(fsm.NewNewPostEvent(
		fsm.NewPostPayload{
			PostID:   post.GetID(),
			AuthorID: post.GetAuthorID(),
			Title:    post.GetTitle(),
			Content:  post.GetContent(),
			Tags:     c._toTags(post.GetTags()),
		},
	))
}

func (c *Community) GoBroadcasting(speakerId string) {
	if c.Broadcast.IsBroadcasting() {
		return
	}
	c.Broadcast.StartBroadcasting(speakerId)
	c.NotifyEvent(fsm.NewGoBroadcastingEvent(fsm.GoBroadcastingPayload{
		SpeakerID: speakerId,
	}))
}

func (c *Community) StopBroadcasting(speakerId string) {
	if !c.Broadcast.IsBroadcasting() {
		return
	}
	c.Broadcast.StopBroadcasting(speakerId)
	c.NotifyEvent(fsm.NewStopBroadcastingEvent(fsm.StopBroadcastingPayload{
		SpeakerID: speakerId,
	}))
}

func (c *Community) Speak(speakMessage *content.Speak) {
	c.Broadcast.SendContent(speakMessage)
	c.NotifyEvent(fsm.NewSpeakEvent(fsm.SpeakPayload{
		SpeakerID: speakMessage.GetAuthorID(),
		Content:   speakMessage.GetContent(),
	}))
}

func (c *Community) NotifyEvent(event *fsm.BaseEvent) {
	for _, listener := range c.EventListeners {
		listener.OnEvent(event)
	}
}

func (c *Community) _toTags(domainTags []tag.Tag) []string {
	tags := make([]string, len(domainTags))
	for i, tag := range domainTags {
		tags[i] = tag.UserID
	}
	return tags
}
