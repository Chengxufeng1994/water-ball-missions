package unit

import (
	"fmt"
	"os"

	"github.com/Chengxufeng1994/water-ball-missions/chapter03/rpg/internal/domain"
	"github.com/Chengxufeng1994/water-ball-missions/chapter03/rpg/internal/domain/action"
)

type BasedUnit struct {
	TroopID          int
	Name             string
	HP               int
	MP               int
	STR              int
	Actions          []domain.Action
	DecisionStrategy domain.DecisionStrategy
}

var _ domain.Unit = (*BasedUnit)(nil)

func NewBasedUnit(name string, hp, mp, str int, actions []domain.Action, decisionStrategy domain.DecisionStrategy) *BasedUnit {
	unit := &BasedUnit{
		Name:             name,
		HP:               hp,
		MP:               mp,
		STR:              str,
		Actions:          actions,
		DecisionStrategy: decisionStrategy,
	}
	unit.Actions = append([]domain.Action{
		action.NewBasicAttack(unit.STR),
	}, unit.Actions...)
	return unit
}

func (unit *BasedUnit) TakeTurn(rpg *domain.RPG) {
	fmt.Printf("輪到 %v\n", unit.Detail())
	// Prepare for the turn

	// List available actions
	actionList := "選擇行動："
	for i, action := range unit.Actions {
		actionList += fmt.Sprintf("(%d) %s ", i, action)
	}
	fmt.Fprintf(os.Stdout, "%s\n", actionList)

	// Choose an action
	chosenAction := unit.DecisionStrategy.ChooseAction(unit.Actions, unit.MP)

	// Choose a targets (if needed)
	skipChosenTargets := false
	oppositeTroop := rpg.GetOppositeTroop()
	candidates := oppositeTroop.GetAliveUnits()
	var chosenTargets []domain.Unit
	if chosenAction.RequiredOfTargets() == action.ALL {
		skipChosenTargets = true
		chosenTargets = candidates
	}
	if chosenAction.RequiredOfTargets() == action.SELF {
		skipChosenTargets = true
		chosenTargets = []domain.Unit{unit}
	}
	if !skipChosenTargets {
		targetList := fmt.Sprintf("選擇 %d 目標：", chosenAction.RequiredOfTargets())
		for i, unit := range oppositeTroop.Units {
			if unit.IsAlive() {
				targetList += fmt.Sprintf("(%d) %s ", i, unit)
			}
		}
		fmt.Fprintf(os.Stdout, "%s\n", targetList)
		chosenTargets = unit.DecisionStrategy.ChooseTargets(candidates, chosenAction.RequiredOfTargets())
	}
	fmt.Printf("%v 對 %v 使用了%s。\n", unit, chosenTargets, chosenAction)

	// Execute the action
	unit.LoseMagicPoint(chosenAction.MagicPointCost())
	for _, target := range chosenTargets {
		chosenAction.Execute(target)
		fmt.Printf("%v 對 %v 造成 %d 點傷害。\n", unit, target, chosenAction.Damage())
		if target.IsDead() {
			fmt.Printf("%v 死亡。\n", target)
		}
	}

	// End the turn
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

func (unit *BasedUnit) OnDamage(damage int) {
	unit.HP -= damage

	if unit.HP < 0 {
		unit.HP = 0
	}
}

func (unit *BasedUnit) Detail() string {
	formatted := "[%d]%s( HP: %d, MP: %d, STR: %d)"
	return fmt.Sprintf(formatted, unit.TroopID, unit.Name, unit.HP, unit.MP, unit.STR)
}

func (unit *BasedUnit) GetActions() []domain.Action {
	return unit.Actions
}

func (unit *BasedUnit) SetTroopID(troopID int) {
	unit.TroopID = troopID
}

func (unit BasedUnit) String() string {
	return fmt.Sprintf("[%d]%s", unit.TroopID, unit.Name)
}
