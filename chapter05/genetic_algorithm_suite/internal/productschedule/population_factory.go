package productschedule

import (
	"fmt"
	"math/rand"

	"github.com/Chengxufeng1994/water-ball-missions/chapter05/genetic_alogrithm_suite/internal/domain"
)

type ProductSchedulePopulationFactory struct{}

var _ domain.PopulationFactory = (*ProductSchedulePopulationFactory)(nil)

func NewProductSchedulePopulationFactory() *ProductSchedulePopulationFactory {
	return &ProductSchedulePopulationFactory{}
}

func (p *ProductSchedulePopulationFactory) InitializePopulation(numberOfIndividuals int) domain.Population {
	individuals := make([]domain.Individual, numberOfIndividuals)

	for i := 0; i < numberOfIndividuals; i++ {
		var chromosome []domain.Gene
		products := []string{"A", "B", "C"}
		product := products[rand.Intn(len(products))]

		for j := 0; j < rand.Intn(10)+1; j++ { // 假設基因長度為 1~10
			duration := rand.Intn(5) + 1
			demand := rand.Intn(50) + 1
			machineID := fmt.Sprintf("M%d", rand.Intn(2)+1)
			workerID := fmt.Sprintf("W%d", rand.Intn(4)+1)
			gene := NewProductScheduleGene(product, duration, demand, machineID, workerID)
			chromosome = append(chromosome, gene)
		}

		individuals[i] = NewProductScheduleIndividual(chromosome)
	}

	return domain.NewPopulation(individuals)
}
