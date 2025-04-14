package domain

type GeneticAlgorithmTemplate interface {
	Select(population Population) (Individual, Individual)
	Cross(parent1, parent2 Individual) Population
	Mutate(individual Individual) Individual
	TerminationCondition() bool
	FindBestIndividual(pop Population) Individual
	// InitializePopulation() Population
	// EvaluateFitness(pop Population)
	// SelectElite(pop Population, n int) []Individual
}
