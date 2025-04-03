package decisionstrategy

import (
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
		if current.MagicPointCost() <= requiredMagicPoint {
			return current
		}
		decisionStrategy.seed++
	}
}

func (decisionStrategy *SeedAIDecisionStrategy) ChooseTargets(targets []domain.Unit, requiredTargets int) []domain.Unit {
	seed := decisionStrategy.seed

	var chosenTargets []domain.Unit
	for range requiredTargets {
		targetIndex := seed % len(targets)
		chosenTargets = append(chosenTargets, targets[targetIndex])
		seed++
	}

	decisionStrategy.seed++

	return chosenTargets
}
