package main

import "sort"

type DistanceBasedStrategy struct{}

var _ MatchmakingStrategy = (*DistanceBasedStrategy)(nil)

func NewDistanceBasedStrategy() *DistanceBasedStrategy {
	return &DistanceBasedStrategy{}
}

func (strategy *DistanceBasedStrategy) match(individual Individual, allIndividuals []Individual) []Individual {
	sort.Slice(allIndividuals, func(i, j int) bool {
		if allIndividuals[i].CalculateDistance(individual) == allIndividuals[j].CalculateDistance(individual) {
			return allIndividuals[i].ID > allIndividuals[j].ID
		}
		return allIndividuals[i].CalculateDistance(individual) < allIndividuals[j].CalculateDistance(individual)
	})

	return allIndividuals
}
