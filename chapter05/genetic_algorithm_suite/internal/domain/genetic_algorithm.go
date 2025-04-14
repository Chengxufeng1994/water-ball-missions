package domain

import (
	"math/rand"
)

type GeneticAlgorithm struct {
	IndividualFactory         IndividualFactory
	GeneFactory               GeneFactory
	MaximumNumberOfIterations int
}

var _ GeneticAlgorithmTemplate = (*GeneticAlgorithm)(nil)

func NewGeneticAlgorithm(
	individualFactory IndividualFactory,
	geneFactory GeneFactory,
	maximumNumberOfIterations int,
) *GeneticAlgorithm {
	return &GeneticAlgorithm{
		IndividualFactory:         individualFactory,
		GeneFactory:               geneFactory,
		MaximumNumberOfIterations: maximumNumberOfIterations,
	}
}

func (ga *GeneticAlgorithm) Evolve(population Population) Individual {
	currentPopulation := population

	for range ga.MaximumNumberOfIterations {

		individuals := make([]Individual, 0)
		for j := 0; j < len(currentPopulation.Individuals)/2; j++ {
			parents1, parents2 := ga.Select(currentPopulation) // 優勝劣汰：篩選這一代中最優秀的一群個體作為「父母」
			offspring := ga.Cross(parents1, parents2)          // 讓這一群父母交配生下新的一代
			mutate1 := ga.Mutate(offspring.Individuals[0])     // 多元化：新的一代中會有基因變異
			mutate2 := ga.Mutate(offspring.Individuals[1])
			individuals = append(individuals, mutate1, mutate2)
		}

		currentPopulation = NewPopulation(individuals)
		if ga.TerminationCondition() {
			break
		}
	}

	return ga.FindBestIndividual(currentPopulation) // 從最終種群中，取得適應度最好的個體
}

func (ga *GeneticAlgorithm) Select(population Population) (Individual, Individual) {
	// Tournament Selection
	numWinners := 2
	winners := make([]Individual, 0, numWinners)

	usedIndices := map[int]bool{}
	total := len(population.Individuals)

	// 兩場錦標賽
	for i := 0; i < numWinners; i++ {
		// 隨機選兩個不同個體進行比賽
		var idx1, idx2 int
		for {
			idx1 = rand.Intn(total)
			if !usedIndices[idx1] {
				break
			}
		}
		usedIndices[idx1] = true

		for {
			idx2 = rand.Intn(total)
			if idx2 != idx1 && !usedIndices[idx2] {
				break
			}
		}
		usedIndices[idx2] = true

		ind1 := population.Individuals[idx1]
		ind2 := population.Individuals[idx2]

		// 挑出適應度較高者
		if ind1.Fitness() >= ind2.Fitness() {
			winners = append(winners, ind1)
		} else {
			winners = append(winners, ind2)
		}
	}

	return winners[0], winners[1]
}

func (ga *GeneticAlgorithm) Cross(p1, p2 Individual) Population {

	minLen := min(len(p1.Chromosome()), len(p2.Chromosome()))
	if minLen == 0 {
		return NewPopulation([]Individual{
			ga.IndividualFactory.CreateIndividual(p1.Chromosome()),
			ga.IndividualFactory.CreateIndividual(p2.Chromosome()),
		})
	}

	// 隨機選交配點
	crossoverPoint := rand.Intn(minLen)

	// 執行一點交配
	child1Chromosome := append(p1.Chromosome()[:crossoverPoint], p2.Chromosome()[crossoverPoint:]...)
	child2Chromosome := append(p2.Chromosome()[:crossoverPoint], p1.Chromosome()[crossoverPoint:]...)
	child1 := ga.IndividualFactory.CreateIndividual(child1Chromosome)
	child2 := ga.IndividualFactory.CreateIndividual(child2Chromosome)

	return NewPopulation([]Individual{child1, child2})
}

func (ga *GeneticAlgorithm) Mutate(individual Individual) Individual {
	chromosome := individual.Chromosome()
	for j := 0; j < len(chromosome); j++ {
		if rand.Float64() < 0.1 {
			chromosome[j] = ga.GeneFactory.CreateGene(j)
		}
	}
	return ga.IndividualFactory.CreateIndividual(chromosome)
}

func (ga *GeneticAlgorithm) TerminationCondition() bool {
	return false
}

func (ga *GeneticAlgorithm) FindBestIndividual(population Population) Individual {
	return population.Individuals[0]
}
