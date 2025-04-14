package productschedule

import "github.com/Chengxufeng1994/water-ball-missions/chapter05/genetic_alogrithm_suite/internal/domain"

type Product struct {
	Name     string
	Duration int // 單位：小時
	Demand   int
}

type Machine struct {
	ID int
}

type Worker struct {
	ID int
}

var (
	products = []Product{
		{"A", 2, 100},
		{"B", 4, 200},
		{"C", 6, 300},
	}
	machines = []Machine{{0}, {1}}
	workers  = []Worker{{0}, {1}, {2}, {3}}
)

type ProductScheduleIndividual struct {
	chromosome []ProductScheduleGene
	fitness    float64
}

var _ domain.Individual = (*ProductScheduleIndividual)(nil)

func NewProductScheduleIndividual(genes []domain.Gene) *ProductScheduleIndividual {
	chromosomes := make([]ProductScheduleGene, len(genes))
	for i, gene := range genes {
		if g, ok := gene.(*ProductScheduleGene); ok {
			chromosomes[i] = *g
		} else {
			panic("invalid gene type for ProductScheduleIndividual")
		}
	}

	individual := &ProductScheduleIndividual{
		chromosome: chromosomes,
		fitness:    0,
	}
	individual.CalculateFitness()

	return individual
}

func (individual *ProductScheduleIndividual) Chromosome() []domain.Gene {
	genes := make([]domain.Gene, len(individual.chromosome))
	for i, gene := range individual.chromosome {
		genes[i] = &gene
	}
	return genes
}

func (individual *ProductScheduleIndividual) CalculateFitness() {
	// 建立機器 + 工人 的生產時間軸 map
	machineWorkerTimeline := make(map[string]int)

	for _, gene := range individual.chromosome {
		resourceKey := gene.MachineID + "-" + gene.WorkerID

		// 假設每個 gene 是連續進行，加入其 duration
		machineWorkerTimeline[resourceKey] += gene.Duration * gene.Demand
	}

	// 最大值即為生產所需總時間
	var max int
	for _, time := range machineWorkerTimeline {
		if time > max {
			max = time
		}
	}

	individual.fitness = float64(max)
}

func (individual *ProductScheduleIndividual) Fitness() float64 {
	return individual.fitness
}
