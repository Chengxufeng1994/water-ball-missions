package bot

import (
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/bot/fields"
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/bot/guard"
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/bot/state/normal"
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/bot/state/record"
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/bot/trigger"
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/fsm"
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/fsm/action"
	fsmguard "github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/fsm/guard"
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/fsm/transition"
	fsmtrigger "github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/fsm/trigger"
)

func NewBotFSM() *fsm.BaseFiniteStateMachine {
	// normal
	defaultConversation := normal.NewDefaultConversation(action.NewPrintAction("entry default conversation"), action.NewPrintAction("exit default conversation"))
	interacting := normal.NewInteractingState(action.NewPrintAction("entry interacting"), action.NewPrintAction("exit interacting"))
	normalSubFSM := fsm.NewSubFiniteStateMachine("normal", defaultConversation,
		action.NewPrintAction("entry normal"),
		action.NewPrintAction("exit normal"),
		[]fsm.Transition{
			transition.NewInitTransition(interacting, guard.NewUserCountLimitGuard(guard.NewGreaterThanNumberGuard(10))),
			transition.NewInitTransition(defaultConversation, fsmguard.NewTruthGuard()),
		},
		[]fsm.Transition{
			transition.NewBaseTransition(defaultConversation, interacting,
				fsmtrigger.NewEventTrigger(fsm.LoginEvent),
				guard.NewUserCountLimitGuard(guard.NewGreaterThanNumberGuard(10)),
				action.NewNoAction()),
			transition.NewBaseTransition(interacting, defaultConversation,
				fsmtrigger.NewEventTrigger(fsm.LogoutEvent),
				guard.NewUserCountLimitGuard(guard.NewLessThanNumberGuard(10)),
				action.NewNoAction()),
		}...,
	)
	normal := normal.NewNormalState(normalSubFSM)

	// record
	recording := record.NewRecording(action.NewPrintAction("entry recording"), action.NewPrintAction("exit recording"))
	waiting := record.NewWaitingState(action.NewPrintAction("entry waiting"), action.NewPrintAction("exit waiting"))
	recordSubFSM := fsm.NewSubFiniteStateMachine("record", waiting,
		action.NewPrintAction("entry record"),
		action.NewPrintAction("exit record"),
		[]fsm.Transition{
			transition.NewInitTransition(recording, guard.NewNotNullGuard(fields.SpeakerId)),
			transition.NewInitTransition(waiting, fsmguard.NewTruthGuard()),
		}, []fsm.Transition{
			transition.NewBaseTransition(waiting, recording,
				fsmtrigger.NewEventTrigger(fsm.GoBroadcastingEvent),
				nil,
				action.NewNoAction()),
			transition.NewBaseTransition(recording, waiting,
				fsmtrigger.NewEventTrigger(fsm.StopBroadcastingEvent),
				nil,
				action.NewNoAction()),
		}...)
	record := record.NewRecordState(recordSubFSM)

	return fsm.NewBaseFiniteStateMachine(normal,
		[]fsm.Transition{
			transition.NewInitTransition(normal, fsmguard.NewTruthGuard()),
		},
		[]fsm.Transition{
			transition.NewBaseTransition(normal, record,
				trigger.NewNewMessageTrigger("record"),
				fsmguard.NewAndGuard(guard.NewTaggedGuard("bot"), guard.NewQuotaLimitGuard(guard.NewGreaterThanNumberGuard(3))),
				action.NewNoAction(),
			),
			transition.NewBaseTransition(record, normal,
				trigger.NewNewMessageTrigger("stop-recording"),
				fsmguard.NewAndGuard(guard.NewTaggedGuard("bot"), guard.NewRecordGuard()),
				action.NewNoAction()),
		}...)
}
