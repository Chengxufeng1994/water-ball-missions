package main

type MatchmakingStrategy interface {
	match(individual Individual, allIndividuals []Individual) []Individual
}
