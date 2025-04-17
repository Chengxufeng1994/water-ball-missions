package knowledgeking

import (
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/bot/fields"
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/fsm"
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/shared"
)

const KnowledgeKingQuota = 5

type KnowledgeKing struct {
	*fsm.SubFiniteStateMachine
}

var _ interface {
	fsm.FiniteStateMachine
	fsm.State
} = (*KnowledgeKing)(nil)

func NewKnowledgeKing(subFiniteStateMachine *fsm.SubFiniteStateMachine) *KnowledgeKing {
	return &KnowledgeKing{
		SubFiniteStateMachine: subFiniteStateMachine,
	}
}

func (s *KnowledgeKing) EntryState(ctx shared.Context, event fsm.Event) {
	s.SubFiniteStateMachine.EntryState(ctx, event)
	ctx.SetValue(fields.Quota, ctx.GetValue(fields.Quota).(int)-KnowledgeKingQuota)
}
