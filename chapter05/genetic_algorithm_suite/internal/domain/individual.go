package domain

type Individual interface {
	Chromosome() []Gene
	Fitness() float64
	CalculateFitness()
}

// type Individual struct {
// 	Chromosomes []Gene
// 	Fitness     int
// }

// func NewIndividual(chromosomes []Gene, fitness int) Individual {
// 	return Individual{
// 		Chromosomes: chromosomes,
// 		Fitness:     fitness,
// 	}
// }
