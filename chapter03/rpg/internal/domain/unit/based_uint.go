package unit

import (
	"fmt"
	"os"
	"time"

	"slices"

	"github.com/Chengxufeng1994/water-ball-missions/chapter03/rpg/internal/domain"
	"github.com/Chengxufeng1994/water-ball-missions/chapter03/rpg/internal/domain/action"
	"github.com/Chengxufeng1994/water-ball-missions/chapter03/rpg/internal/domain/state"
)

type BasedUnit struct {
	ID               int
	Troop            *domain.Troop
	Name             string
	HP               int
	MP               int
	STR              int
	actionable       bool
	State            domain.State
	DecisionStrategy domain.DecisionStrategy
	Actions          []domain.Action
	Summoner         domain.Summoner
	Cursers          []domain.Cusrer
}

var _ interface {
	domain.Unit
	domain.Summoner
	domain.Cusrer
} = (*BasedUnit)(nil)

func NewBasedUnit(id int, name string, hp, mp, str int, actions []domain.Action, decisionStrategy domain.DecisionStrategy) *BasedUnit {
	unit := &BasedUnit{
		ID:               id,
		Name:             name,
		HP:               hp,
		MP:               mp,
		STR:              str,
		actionable:       true,
		State:            state.NewNormalState(),
		DecisionStrategy: decisionStrategy,
		Actions:          actions,
		Summoner:         nil,
		Cursers:          make([]domain.Cusrer, 0),
	}

	unit.Actions = append([]domain.Action{
		action.NewBasicAttack(),
	}, unit.Actions...)

	return unit
}

func (unit *BasedUnit) TakeTurn(rpg *domain.RPG) {
	fmt.Printf("輪到 %v\n", unit.Detail())
	// Prepare for the turn
	unit.State.PreTurn()

	if unit.Actionable() {
		// List available actions
		actionList := "選擇行動："
		for i, action := range unit.Actions {
			actionList += fmt.Sprintf("(%d) %s ", i, action)
		}
		fmt.Fprintf(os.Stdout, "%s\n", actionList)

		// Choose an action
		chosenAction := unit.DecisionStrategy.ChooseAction(unit.Actions, unit.MP)

		// Choose a targets (if needed)
		candidates := rpg.GetOppositeTroop().Units
		var chosenTargets []domain.Unit
		switch chosenAction.RequiredTargetType() {
		case action.TARGET_TYPE_SELF:
			candidates = []domain.Unit{unit}
		case action.TARGET_TYPE_ALL:
			candidates = rpg.GetAllAliveUnits()
		case action.TARGET_TYPE_ALL_ALLY:
			candidates = rpg.GetCurrentUnitTroopUnits()
		case action.TARGET_TYPE_ALL_ENEMY:
			candidates = rpg.GetCurrentUnitOppositeTroopUnits()
		}

		if chosenAction.RequiredOfTargets() <= 0 {
			chosenTargets = candidates
		} else if len(candidates) <= chosenAction.RequiredOfTargets() {
			chosenTargets = candidates
		} else {
			targetList := fmt.Sprintf("選擇 %d 目標：", chosenAction.RequiredOfTargets())
			for i, unit := range candidates {
				if unit.IsAlive() {
					targetList += fmt.Sprintf("(%d) %s ", i, unit)
				}
			}
			fmt.Fprintf(os.Stdout, "%s\n", targetList)
			chosenTargets = unit.DecisionStrategy.ChooseTargets(candidates, chosenAction.RequiredOfTargets())
		}

		fmt.Println(chosenAction.Description(unit, chosenTargets))

		// Execute the action
		chosenAction.Execute(rpg, unit, chosenTargets)
	}

	unit.State.PostTurn()
}

func (unit *BasedUnit) GetID() int {
	return unit.ID
}

func (unit *BasedUnit) IsHero() bool {
	return false
}

func (unit *BasedUnit) IsAlive() bool {
	return unit.HP > 0
}

func (unit *BasedUnit) IsDead() bool {
	return unit.HP <= 0
}

func (unit *BasedUnit) LoseMagicPoint(amount int) {
	unit.MP -= amount

	if unit.MP < 0 {
		unit.MP = 0
	}
}

func (unit *BasedUnit) GetStrength() int {
	return unit.STR + unit.State.BonusStrength()
}

func (unit *BasedUnit) GetHP() int {
	return unit.HP
}

func (unit *BasedUnit) AddStrength(amount int) {
	unit.STR += amount
}

func (unit *BasedUnit) LoseStrength(amount int) {
	unit.STR -= amount
}

func (unit *BasedUnit) checkUnitDead() {
	if unit.HP <= 0 {
		fmt.Printf("%v 死亡。\n", unit)
		unit.notifySummoner()
		unit.notifyCurser()
	}
}

func (unit *BasedUnit) notifySummoner() {
	if unit.Summoner != nil {
		unit.Summoner.OnSummonedDead(30)
		unit.RemoveSummoner(unit.Summoner)
	}
}

func (unit *BasedUnit) notifyCurser() {
	m := make(map[int]struct{})
	for _, curser := range unit.Cursers {
		if _, ok := m[curser.GetID()]; ok {
			continue
		}
		curser.OnCursedDead(unit.MP)
		unit.RemoveCurser(curser)
		m[curser.GetID()] = struct{}{}
	}
}

func (unit *BasedUnit) RetrieveState(newState domain.State) {
	unit.State.ExitState()

	newState.SetUnit(unit)
	unit.State = newState
	unit.State.EntryState()
}

func (unit *BasedUnit) GetCurrentState() domain.State {
	return unit.State
}

func (unit *BasedUnit) Actionable() bool {
	return unit.actionable
}

func (unit *BasedUnit) SetActionable(enabled bool) {
	unit.actionable = enabled
}

func (unit *BasedUnit) GetActions() []domain.Action {
	return unit.Actions
}

func (unit *BasedUnit) GetTroop() *domain.Troop {
	return unit.Troop
}

func (unit *BasedUnit) SetTroop(troop *domain.Troop) {
	unit.Troop = troop
}

func (unit *BasedUnit) Summon() {
	slime := NewSlimeUnit(int(time.Now().Unix()))
	slime.AddSummoner(unit)
	unit.Troop.AddUnit(slime)
}

func (unit *BasedUnit) OnDamage(damage int) {
	unit.HP -= damage
	unit.checkUnitDead()
}

func (unit *BasedUnit) OnHeal(amount int) {
	unit.HP += amount
}

func (unit *BasedUnit) OnPoisoned(amount int) {
	unit.HP -= amount
	unit.checkUnitDead()
}

func (unit *BasedUnit) OnCurse(curser domain.Unit) {
	if cusrer, ok := curser.(domain.Cusrer); ok {
		unit.Cursers = append(unit.Cursers, cusrer)
		fmt.Println(unit.Cursers)
	} else {
		fmt.Printf("Warning: curser does not implement domain.Cusrer: %v\n", curser)
	}
}

func (unit *BasedUnit) Suicide() {
	unit.HP = 0
}

func (unit *BasedUnit) AddSummoner(summoner domain.Summoner) {
	unit.Summoner = summoner
}

func (unit *BasedUnit) RemoveSummoner(summoner domain.Summoner) {
	unit.Summoner = nil
}

func (unit *BasedUnit) OnSummonedDead(amount int) {
	if unit.IsDead() {
		return
	}
	unit.HP += amount
}

// AddCurser adds a curser to the unit.
func (unit *BasedUnit) AddCurser(curser domain.Cusrer) {
	unit.Cursers = append(unit.Cursers, curser)
}

func (unit *BasedUnit) RemoveCurser(curser domain.Cusrer) {
	for i, c := range unit.Cursers {
		if c.GetID() == curser.GetID() {
			unit.Cursers = slices.Delete(unit.Cursers, i, i+1)
		}
	}
}

func (unit *BasedUnit) OnCursedDead(amount int) {
	if unit.IsDead() {
		return
	}
	unit.HP += amount
}

func (unit *BasedUnit) Detail() string {
	formatted := "[%d]%s( HP: %d, MP: %d, STR: %d, State: %s)"
	return fmt.Sprintf(formatted, unit.Troop.ID, unit.Name, unit.HP, unit.MP, unit.STR, unit.State)
}

func (unit BasedUnit) String() string {
	return fmt.Sprintf("[%d]%s", unit.Troop.ID, unit.Name)
}
