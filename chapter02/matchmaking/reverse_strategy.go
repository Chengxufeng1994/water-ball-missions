package main

import "sort"

type ReverseStrategy struct {
	basedStrategy MatchmakingStrategy
}

var _ MatchmakingStrategy = (*ReverseStrategy)(nil)

func NewReverseStrategy(basedStrategy MatchmakingStrategy) *ReverseStrategy {
	return &ReverseStrategy{basedStrategy: basedStrategy}
}

func (r *ReverseStrategy) match(individual Individual, allIndividuals []Individual) []Individual {
	result := r.basedStrategy.match(individual, allIndividuals)
	sort.SliceStable(result, func(i, j int) bool { return i > j })
	return result
}
