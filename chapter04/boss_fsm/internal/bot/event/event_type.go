package event

import "github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/fsm"

const (
	LoginEvent            fsm.EventType = "login"
	LogoutEvent           fsm.EventType = "logout"
	TimeElapsedEvent      fsm.EventType = "time_elapsed"
	NewMessageEvent       fsm.EventType = "new_message"
	NewPostEvent          fsm.EventType = "new_post"
	GoBroadcastingEvent   fsm.EventType = "go_broadcasting"
	SpeakEvent            fsm.EventType = "speak"
	StopBroadcastingEvent fsm.EventType = "stop_broadcasting"
	KnowledgeKingEndEvent fsm.EventType = "knowledge_king_end"
	EndEvent              fsm.EventType = "end"
)
