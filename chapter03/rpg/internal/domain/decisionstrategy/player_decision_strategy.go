package decisionstrategy

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Chengxufeng1994/water-ball-missions/chapter03/rpg/internal/domain"
)

type PlayerDecisionStrategy struct {
	UserInput UserInput
}

var _ domain.DecisionStrategy = (*PlayerDecisionStrategy)(nil)

func NewPlayerDecisionStrategy() *PlayerDecisionStrategy {
	return &PlayerDecisionStrategy{UserInput: UserInput{}}
}

func (pds *PlayerDecisionStrategy) ChooseAction(actions []domain.Action, requiredMagicPoint int) domain.Action {
	var actionIndex int
	for {
		if err := pds.UserInput.GetInput(); err != nil {
			continue
		}

		if _, err := fmt.Sscanf(pds.UserInput.Input, "%d", &actionIndex); err != nil {
			continue
		}

		if actionIndex < 0 || actionIndex >= len(actions) {
			continue
		}

		if actions[actionIndex].MagicPointCost() > requiredMagicPoint {
			fmt.Println("你缺乏 MP，不能進行此行動。")
			continue
		}

		break
	}

	return actions[actionIndex]
}

func (pds *PlayerDecisionStrategy) ChooseTargets(units []domain.Unit, requiredTargets int) []domain.Unit {
	var indices []int

	for {
		if err := pds.UserInput.GetInput(); err != nil {
			continue
		}

		inputs := strings.Fields(pds.UserInput.Input)
		indices = make([]int, len(inputs))

		for i, input := range inputs {
			if _, err := fmt.Sscanf(input, "%d", &indices[i]); err != nil {
				continue
			}
		}

		if len(indices) == 0 || len(indices) > requiredTargets {
			continue
		}

		valid := true
		for _, idx := range indices {
			if idx < 0 || idx >= len(units) {
				valid = false
				break
			}
		}

		if valid {
			break
		}
	}

	selectedUnits := make([]domain.Unit, len(indices))
	for i, idx := range indices {
		selectedUnits[i] = units[idx]
	}

	return selectedUnits
}

type UserInput struct {
	Input string
}

func (ui *UserInput) GetInput() error {
	// reader := bufio.NewReader(os.Stdin)
	// text, _ := reader.ReadString('\n')
	// text = strings.TrimSpace(text)
	// if strings.EqualFold(text, "") {
	// 	return fmt.Errorf("輸入值不可為空")
	// }

	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		return scanner.Err()
	}

	input := scanner.Text()
	text := strings.TrimSpace(input)
	if strings.EqualFold(text, "") {
		return fmt.Errorf("輸入值不可為空")
	}

	ui.Input = text

	return nil
}
