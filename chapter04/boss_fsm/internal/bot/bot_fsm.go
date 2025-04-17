package bot

import (
	"time"

	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/bot/action"
	botevent "github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/bot/event"
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/bot/fields"
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/bot/guard"
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/bot/state/knowledgeking"
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/bot/state/normal"
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/bot/state/record"
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/bot/trigger"
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/fsm"
	fsmaction "github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/fsm/action"
	fsmguard "github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/fsm/guard"
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/fsm/state"
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/fsm/transition"
	fsmtrigger "github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/fsm/trigger"
)

func NewBotFSM(bot *CommunityBot) *fsm.BaseFiniteStateMachine {
	// normal
	defaultConversation := normal.NewDefaultConversation(fsmaction.NewLoggerAction("entry default_conversation"), fsmaction.NewLoggerAction("exit default_conversation"), bot)
	interacting := normal.NewInteractingState(fsmaction.NewLoggerAction("entryinteracting"), fsmaction.NewLoggerAction("exit interacting"), bot)
	normalSubFSM := fsm.NewSubFiniteStateMachine("normal", defaultConversation,
		fsmaction.NewLoggerAction("entry normal"),
		fsmaction.NewLoggerAction("exit normal"),
		[]fsm.Transition{
			transition.NewInitTransition(interacting, guard.NewUserCountLimitGuard(guard.NewGreaterThanEqualNumberGuard(10))),
			transition.NewInitTransition(defaultConversation, fsmguard.NewTruthGuard()),
		},
		[]fsm.Transition{
			transition.NewBaseTransition(defaultConversation, interacting, fsmtrigger.NewEventTrigger(botevent.LoginEvent),
				guard.NewUserCountLimitGuard(guard.NewGreaterThanEqualNumberGuard(10)), fsmaction.NewNoAction()),
			transition.NewBaseTransition(interacting, defaultConversation, fsmtrigger.NewEventTrigger(botevent.LogoutEvent),
				guard.NewUserCountLimitGuard(guard.NewLessThanNumberGuard(10)), fsmaction.NewNoAction()),
		}...,
	)
	normal := normal.NewNormalState(normalSubFSM)

	// record
	recording := record.NewRecording(fsmaction.NewLoggerAction("entry recording"), fsmaction.NewLoggerAction("exit recording"), bot)
	waiting := record.NewWaitingState(fsmaction.NewLoggerAction("entry waiting"), fsmaction.NewLoggerAction("exit waiting"), bot)
	recordSubFSM := fsm.NewSubFiniteStateMachine("record", waiting,
		fsmaction.NewLoggerAction("entry record"),
		fsmaction.NewLoggerAction("exit record"),
		[]fsm.Transition{
			transition.NewInitTransition(recording, guard.NewNotNullGuard(fields.SpeakerId)),
			transition.NewInitTransition(waiting, fsmguard.NewTruthGuard()),
		}, []fsm.Transition{
			transition.NewBaseTransition(waiting, recording, fsmtrigger.NewEventTrigger(botevent.GoBroadcastingEvent),
				fsmguard.NewTruthGuard(), fsmaction.NewNoAction()),
			transition.NewBaseTransition(recording, waiting, fsmtrigger.NewEventTrigger(botevent.StopBroadcastingEvent),
				fsmguard.NewTruthGuard(), fsmaction.NewNoAction()),
		}...)
	record := record.NewRecordState(recordSubFSM)

	// knowledge king
	thanksForJoining := knowledgeking.NewThanksForJoining(fsmaction.NewLoggerAction("entry thanks for joining"), fsmaction.NewLoggerAction("exit thanks for joining"), bot)
	questioning := knowledgeking.NewQuestioning(fsmaction.NewLoggerAction("entry questioning"), fsmaction.NewLoggerAction("exit questioning"), bot)
	knowledgeKingSubFSM := fsm.NewSubFiniteStateMachine("knowledgeKing", questioning,
		fsmaction.NewLoggerAction("entry knowledgeKing"),
		fsmaction.NewLoggerAction("exit knowledgeKing"),
		[]fsm.Transition{transition.NewInitTransition(questioning, fsmguard.NewTruthGuard())},
		[]fsm.Transition{
			transition.NewBaseTransition(questioning, thanksForJoining, trigger.NewElapsedTimeTrigger(time.Hour),
				fsmguard.NewTruthGuard(), fsmaction.NewNoAction()),
			transition.NewBaseTransition(questioning, thanksForJoining, fsmtrigger.NewEventTrigger(botevent.KnowledgeKingEndEvent),
				fsmguard.NewTruthGuard(), fsmaction.NewNoAction()),
			transition.NewBaseTransition(thanksForJoining, questioning, trigger.NewNewMessageTrigger("play again"),
				fsmguard.NewTruthGuard(), action.NewSendMessageAction(bot, "KnowledgeKing is gonna start again!")),
		}...)
	knowledge := knowledgeking.NewKnowledgeKing(knowledgeKingSubFSM)

	return fsm.NewBaseFiniteStateMachine(state.NewInitState(),
		[]fsm.Transition{transition.NewInitTransition(normal, fsmguard.NewTruthGuard())},
		[]fsm.Transition{
			transition.NewBaseTransition(normal, record, trigger.NewNewMessageTrigger("record"),
				fsmguard.NewAndGuard(guard.NewTaggedGuard("bot"), guard.NewQuotaLimitGuard(guard.NewGreaterThanEqualNumberGuard(3))), fsmaction.NewNoAction()),
			transition.NewBaseTransition(normal, knowledge, trigger.NewNewMessageTrigger("king"),
				fsmguard.NewAndGuard(guard.NewAdminGuard(), guard.NewTaggedGuard("bot"), guard.NewQuotaLimitGuard(guard.NewGreaterThanEqualNumberGuard(5))), action.NewSendMessageAction(bot, "KnowledgeKing is started!")),
			transition.NewBaseTransition(knowledge, normal, trigger.NewNewMessageTrigger("king-stop"),
				fsmguard.NewAndGuard(guard.NewAdminGuard(), guard.NewTaggedGuard("bot")), fsmaction.NewNoAction()),
			transition.NewBaseTransition(record, normal, trigger.NewNewMessageTrigger("stop-recording"),
				fsmguard.NewAndGuard(guard.NewTaggedGuard("bot"), guard.NewRecordGuard()), fsmaction.NewNoAction()),
		}...)
}
