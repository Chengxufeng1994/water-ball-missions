package domain

type PopulationFactory interface {
	InitializePopulation(numberOfIndividuals int) Population
}
