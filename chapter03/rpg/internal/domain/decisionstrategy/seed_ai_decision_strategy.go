package decisionstrategy

import (
	"fmt"

	"github.com/Chengxufeng1994/water-ball-missions/chapter03/rpg/internal/domain"
)

type SeedAIDecisionStrategy struct {
	seed int
}

var _ domain.DecisionStrategy = (*SeedAIDecisionStrategy)(nil)

func NewSeedAIDecisionStrategy() *SeedAIDecisionStrategy {
	return &SeedAIDecisionStrategy{
		seed: 0,
	}
}

func (decisionStrategy *SeedAIDecisionStrategy) ChooseAction(actions []domain.Action, requiredMagicPoint int) domain.Action {
	for {
		actionIndex := decisionStrategy.seed % len(actions)
		current := actions[actionIndex]

		if current.MagicPointCost() > requiredMagicPoint {
			fmt.Println("你缺乏 MP，不能進行此行動。")
			decisionStrategy.seed++
			continue
		}

		return current
	}
}

func (decisionStrategy *SeedAIDecisionStrategy) ChooseTargets(targets []domain.Unit, requiredTargets int) []domain.Unit {
	seed := decisionStrategy.seed

	var chosenTargets []domain.Unit
	for i := 0; i < requiredTargets; i++ {
		targetIndex := seed % len(targets)
		chosenTargets = append(chosenTargets, targets[targetIndex])
		seed++
	}

	decisionStrategy.seed++

	return chosenTargets
}
