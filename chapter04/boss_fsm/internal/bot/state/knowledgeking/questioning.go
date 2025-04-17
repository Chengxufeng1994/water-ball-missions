package knowledgeking

import (
	"fmt"
	"strings"
	"sync"

	botevent "github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/bot/event"
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/bot/fields"
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/fsm"
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/fsm/state"
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/boss_fsm/internal/shared"
)

type Questioning struct {
	*state.BaseState
	Questions            []*Question
	CurrentQuestionIndex int
	LeaderBoard          map[string]int
	sync.Mutex
}

var _ fsm.State = (*Questioning)(nil)

func NewQuestioning(entryAction, exitAction fsm.Action, botFsmAdapter state.BotFsmAdapter) *Questioning {
	questions := make([]*Question, 0)
	questions = append(questions, NewQuestion("請問哪個 SQL 語句用於選擇所有的行？", "A", map[string]string{
		"A": "SELECT *",
		"B": "SELECT ALL",
		"C": "SELECT ROWS",
		"D": "SELECT DATA",
	}))
	questions = append(questions, NewQuestion("請問哪個 CSS 屬性可用於設置文字的顏色？", "C", map[string]string{
		"A": "text-align",
		"B": "font-size",
		"C": "color",
		"D": "padding",
	}))
	questions = append(questions, NewQuestion("請問在計算機科學中，「XML」代表什麼？", "A", map[string]string{
		"A": "Extensible Markup Language",
		"B": "Extensible Modeling Language",
		"C": "Extended Markup Language",
		"D": "Extended Modeling Language",
	}))

	return &Questioning{
		BaseState:            state.NewBaseState("questioning", entryAction, exitAction, botFsmAdapter),
		Questions:            questions,
		LeaderBoard:          make(map[string]int),
		CurrentQuestionIndex: 0,
	}
}

func (state *Questioning) EntryState(ctx shared.Context, event fsm.Event) {
	state.BaseState.EntryAction.Execute()
	state.ProcessEntryState(ctx)
}

func (state *Questioning) ProcessEntryState(ctx shared.Context) {
	state.CurrentQuestionIndex = 0
	state.AskQuestion()
}

func (state *Questioning) OnEvent(ctx shared.Context, event fsm.Event) fsm.Event {
	if event.GetEventType() == botevent.NewMessageEvent {
		state.Lock()
		defer state.Unlock()
		payload := event.GetEventPayload().(botevent.NewMessagePayload)
		question := state.Questions[state.CurrentQuestionIndex]
		answer := payload.Content
		if question.Answer == answer {
			state.CurrentQuestionIndex++
			state.LeaderBoard[payload.AuthorID]++
			state.Adapter.SendMessage("Congrats! you got the answer!", payload.AuthorID)
			if state.CurrentQuestionIndex >= len(state.Questions) {
				return botevent.NewKnowledgeKingEndEvent()
			} else {
				state.AskQuestion()
			}
		}
	}

	return nil
}

func (state *Questioning) ExitState(ctx shared.Context, event fsm.Event) {
	state.BaseState.ExitAction.Execute()
	state.ProcessExitState(ctx)
}

func (state *Questioning) ProcessExitState(ctx shared.Context) {
	var winners []string
	highestScore := -1

	for userID, score := range state.LeaderBoard {
		if score > highestScore {
			highestScore = score
			winners = []string{userID}
		} else if score == highestScore {
			winners = append(winners, userID)
		}
	}
	if len(winners) == 1 {
		ctx.SetValue(fields.WinnerId, winners[0])
		return
	}
}

func (state *Questioning) AskQuestion() {
	problem := state.Questions[state.CurrentQuestionIndex].Problem
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("%d. %s\n", state.CurrentQuestionIndex, problem))
	n := len(state.Questions[state.CurrentQuestionIndex].Options)
	i := 0
	for k, v := range state.Questions[state.CurrentQuestionIndex].Options {
		sb.WriteString(fmt.Sprintf("%s) %s", k, v))
		if i < n-1 {
			sb.WriteString("\n")
		}
		i++
	}
	state.Adapter.SendMessage(sb.String())
}
