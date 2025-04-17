package community

import (
	"fmt"
	"slices"
	"strconv"
	"time"

	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/bot"
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/bot/event"
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/community/channel"
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/community/content"
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/community/tag"
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/fsm"
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/shared"
)

type Community struct {
	Name           string
	Members        []*shared.Member
	ChatRoom       *channel.Chatroom  // Channel
	Forum          *channel.Forum     // Channel
	Broadcast      *channel.Broadcast // Channel
	Bot            *bot.CommunityBot
	EventListeners []fsm.EventListener
}

func NewCommunity(name string) *Community {
	return &Community{
		Name:      name,
		Members:   make([]*shared.Member, 0),
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
		for i, t := range payload.Tags {
			tags[i] = tag.Tag{
				UserID: t,
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
	case ElapsedTimeEvent:
		payload, ok := event.Payload().(ElapsedTimeEventPayload)
		if !ok {
			return ErrInvalidPayload
		}
		c.ElapsedTime(payload)
		return nil
	case EndEvent:
		return nil
	default:
		return ErrCommunityEvent
	}
}

func (c *Community) Started(payload StartedEventPayload) {

	bot := bot.NewBot(payload.Quota)
	c.EventListeners = append(c.EventListeners, bot)
	c.Members = append(c.Members, shared.NewAdminMember("bot"))
}

func (c *Community) Login(member *shared.Member) {
	c.Members = append(c.Members, member)
	c.NotifyEvent(event.NewLoginEvent(event.LoginPayload{
		UserID:  member.UserID,
		IsAdmin: member.IsAdmin,
	}))
}

func (c *Community) Logout(memberId string) {
	for i, m := range c.Members {
		if m.UserID == memberId {
			c.Members = slices.Delete(c.Members, i, i+1)
			c.NotifyEvent(event.NewLogoutEvent(event.LogoutPayload{
				UserID: memberId,
			}))
		}
	}
}

func (c *Community) NewMessage(message *content.Message) {
	c.ChatRoom.SendContent(message)
	c.NotifyEvent(event.NewNewMessageEvent(event.NewMessagePayload{
		AuthorID: message.GetAuthorID(),
		Content:  message.GetContent(),
		Tags:     c._toTags(message.GetTags()),
	}))
}

func (c *Community) NewPost(post *content.Post) {
	c.Forum.SendContent(post)
	c.NotifyEvent(event.NewNewPostEvent(
		event.NewPostPayload{
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
	c.NotifyEvent(event.NewGoBroadcastingEvent(event.GoBroadcastingPayload{
		SpeakerID: speakerId,
	}))
}

func (c *Community) StopBroadcasting(speakerId string) {
	if !c.Broadcast.IsBroadcasting() {
		return
	}
	c.Broadcast.StopBroadcasting(speakerId)
	c.NotifyEvent(event.NewStopBroadcastingEvent(event.StopBroadcastingPayload{
		SpeakerID: speakerId,
	}))
}

func (c *Community) Speak(speakMessage *content.Speak) {
	c.Broadcast.SendContent(speakMessage)
	c.NotifyEvent(event.NewSpeakEvent(event.SpeakPayload{
		SpeakerID: speakMessage.GetAuthorID(),
		Content:   speakMessage.GetContent(),
	}))
}

func (c *Community) ElapsedTime(payload ElapsedTimeEventPayload) {
	fmt.Printf("🕑 %s %s elapsed...\n", payload.Number, payload.Unit)
	num, _ := strconv.Atoi(payload.Number)
	var duration time.Duration
	switch payload.Unit {
	case "seconds":
		duration = time.Duration(num) * time.Second
	case "minutes":
		duration = time.Duration(num) * time.Minute
	case "hours":
		duration = time.Duration(num) * time.Hour
	}

	c.NotifyEvent(event.NewTimeElapsedEvent(event.ElapsedTimePayload{Duration: duration}))
}

func (c *Community) NotifyEvent(event fsm.Event) {
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
