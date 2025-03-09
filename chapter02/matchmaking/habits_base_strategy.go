package main

import (
	"sort"
)

type HabitsBasedStrategy struct{}

var _ MatchmakingStrategy = (*HabitsBasedStrategy)(nil)

func NewHabitsBasedStrategy() *HabitsBasedStrategy {
	return &HabitsBasedStrategy{}
}

func (h *HabitsBasedStrategy) match(individual Individual, allIndividuals []Individual) []Individual {
	sort.Slice(allIndividuals, func(i, j int) bool {
		if allIndividuals[i].CalculateHabitsScore(individual) == allIndividuals[j].CalculateHabitsScore(individual) {
			return allIndividuals[i].ID > allIndividuals[j].ID
		}
		return allIndividuals[i].CalculateHabitsScore(individual) > allIndividuals[j].CalculateHabitsScore(individual)
	})

	return allIndividuals
}
