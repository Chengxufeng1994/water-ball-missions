package main

import (
	"fmt"

	"github.com/Chengxufeng1994/water-ball-missions/chapter05/genetic_alogrithm_suite/internal/domain"
	"github.com/Chengxufeng1994/water-ball-missions/chapter05/genetic_alogrithm_suite/internal/productrecommendation"
)

func main() {
	populationFactory := productrecommendation.NewProductRecommendationPopulationFactory()
	population := populationFactory.InitializePopulation(20)
	individualFactory := productrecommendation.NewProductRecommendationIndividualFactory()
	geneFactory := productrecommendation.NewProductRecommendationGeneFactory()
	best := domain.NewGeneticAlgorithm(individualFactory, geneFactory, 20).Evolve(population)
	fmt.Printf("best: %v\n", best)
}
