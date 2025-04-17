package community

import (
	"time"
)

type CommunityEventName string

type CommunityEventPayload interface{}

type (
	CommunityEvent interface {
		EventName() CommunityEventName
		Payload() CommunityEventPayload
		OccurredAt() time.Time
	}

	communityEvent struct {
		eventName  CommunityEventName
		payload    CommunityEventPayload
		occurredAt time.Time
	}
)

var _ CommunityEvent = (*communityEvent)(nil)

func NewCommunityEvent(name CommunityEventName, payload CommunityEventPayload) CommunityEvent {
	return &communityEvent{
		eventName:  name,
		payload:    payload,
		occurredAt: time.Now(),
	}
}

func (c *communityEvent) EventName() CommunityEventName {
	return c.eventName
}

func (c *communityEvent) Payload() CommunityEventPayload {
	return c.payload
}

func (c *communityEvent) OccurredAt() time.Time {
	return c.occurredAt
}

const (
	StartedEvent          CommunityEventName = "started"
	LoginEvent            CommunityEventName = "login"
	LogoutEvent           CommunityEventName = "logout"
	NewMessageEvent       CommunityEventName = "new message"
	NewPostEvent          CommunityEventName = "new post"
	GoBroadcastingEvent   CommunityEventName = "go broadcasting"
	StopBroadcastingEvent CommunityEventName = "stop broadcasting"
	SpeakEvent            CommunityEventName = "speak"
	ElapsedTimeEvent      CommunityEventName = "elapsed"
	EndEvent              CommunityEventName = "end"
)

type StartedEventPayload struct {
	Time  time.Time `json:"time"`
	Quota int       `json:"quota"`
}

func NewStartedEvent(payload StartedEventPayload) CommunityEvent {
	return NewCommunityEvent(StartedEvent, payload)
}

type LoginEventPayload struct {
	UserID  string `json:"userId"`
	IsAdmin bool   `json:"isAdmin"`
}

func NewLoginEvent(payload LoginEventPayload) CommunityEvent {
	return NewCommunityEvent(LoginEvent, payload)
}

type LogoutEventPayload struct {
	UserID string `json:"userId"`
}

func NewLogoutEvent(payload LogoutEventPayload) CommunityEvent {
	return NewCommunityEvent(LogoutEvent, payload)
}

type NewMessageEventPayload struct {
	AuthorID string   `json:"authorId"`
	Content  string   `json:"content"`
	Tags     []string `json:"tags"`
}

func NewNewMessageEvent(payload NewMessageEventPayload) CommunityEvent {
	return NewCommunityEvent(NewMessageEvent, payload)
}

type NewPostEventPayload struct {
	PostID   string   `json:"id"`
	Title    string   `json:"title"`
	Content  string   `json:"content"`
	AuthorID string   `json:"authorId"`
	Tags     []string `json:"tags"`
}

func NewNewPostEvent(payload NewPostEventPayload) CommunityEvent {
	return NewCommunityEvent(NewPostEvent, payload)
}

type GoBroadcastingEventPayload struct {
	SpeakerID string `json:"speakerId"`
}

func NewGoBroadcastingEvent(payload GoBroadcastingEventPayload) CommunityEvent {
	return NewCommunityEvent(GoBroadcastingEvent, payload)
}

type StopBroadcastingEventPayload struct {
	SpeakerID string `json:"speakerId"`
}

func NewStopBroadcastingEvent(payload StopBroadcastingEventPayload) CommunityEvent {
	return NewCommunityEvent(StopBroadcastingEvent, payload)
}

type SpeakEventPayload struct {
	SpeakerID string `json:"speakerId"`
	Content   string `json:"content"`
}

func NewSpeakEvent(payload SpeakEventPayload) CommunityEvent {
	return NewCommunityEvent(SpeakEvent, payload)
}

type ElapsedTimeEventPayload struct {
	Number string `json:"number"`
	Unit   string `json:"unit"`
}

func NewElapsedTimeEvent(payload ElapsedTimeEventPayload) CommunityEvent {
	return NewCommunityEvent(ElapsedTimeEvent, payload)
}

func NewEndEvent() CommunityEvent {
	return NewCommunityEvent(EndEvent, nil)
}
