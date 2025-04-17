package event

import "github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/fsm"

type GoBroadcastingPayload struct {
	SpeakerID string `json:"speakerId"`
}

type StopBroadcastingPayload struct {
	SpeakerID string `json:"speakerId"`
}

type SpeakPayload struct {
	SpeakerID string `json:"speakerId"`
	Content   string `json:"content"`
}

func NewGoBroadcastingEvent(payload GoBroadcastingPayload) fsm.Event {
	return fsm.NewEvent(GoBroadcastingEvent, payload)
}

func NewStopBroadcastingEvent(payload StopBroadcastingPayload) fsm.Event {
	return fsm.NewEvent(StopBroadcastingEvent, payload)
}

func NewSpeakEvent(payload SpeakPayload) fsm.Event {
	return fsm.NewEvent(SpeakEvent, payload)
}
