package main

import (
	"fmt"
)

type MatchmakingSystem struct {
	strategy MatchmakingStrategy
}

func NewMatchmakingSystem(strategy MatchmakingStrategy) *MatchmakingSystem {
	return &MatchmakingSystem{strategy: strategy}
}

func (m *MatchmakingSystem) Matching(individual Individual, allIndividuals []Individual) Individual {
	newIndividuals := make([]Individual, 0)
	for _, other := range allIndividuals {
		if other.ID != individual.ID {
			newIndividuals = append(newIndividuals, other)
		}
	}

	result := m.strategy.match(individual, newIndividuals)
	fmt.Println("match result:", result)
	return result[0]
}
