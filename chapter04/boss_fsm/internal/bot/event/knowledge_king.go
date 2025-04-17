package event

import "github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/fsm"

func NewKnowledgeKingEndEvent() fsm.Event {
	return fsm.NewEvent(KnowledgeKingEndEvent, nil)
}
