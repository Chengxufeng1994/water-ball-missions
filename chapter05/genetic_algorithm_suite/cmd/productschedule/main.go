package main

import (
	"fmt"

	"github.com/Chengxufeng1994/water-ball-missions/chapter05/genetic_alogrithm_suite/internal/domain"
	"github.com/Chengxufeng1994/water-ball-missions/chapter05/genetic_alogrithm_suite/internal/productschedule"
)

func main() {
	populationFactory := productschedule.NewProductSchedulePopulationFactory()
	population := populationFactory.InitializePopulation(20)
	productScheduleIndividualFactory := productschedule.NewProductScheduleIndividualFactory()
	productScheduleGeneFactory := productschedule.NewProductScheduleGeneFactory()
	best := domain.NewGeneticAlgorithm(productScheduleIndividualFactory, productScheduleGeneFactory, 20).Evolve(population)
	fmt.Printf("best: %v\n", best)
}
