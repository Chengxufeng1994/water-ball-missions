package productrecommendation

import (
	"math/rand"

	"github.com/Chengxufeng1994/water-ball-missions/chapter05/genetic_alogrithm_suite/internal/domain"
)

type PopulationFactory struct{}

var _ domain.PopulationFactory = (*PopulationFactory)(nil)

func NewProductRecommendationPopulationFactory() *PopulationFactory {
	return &PopulationFactory{}
}

func (p *PopulationFactory) InitializePopulation(numberOfIndividuals int) domain.Population {
	individuals := make([]domain.Individual, numberOfIndividuals)

	for i := 0; i < numberOfIndividuals; i++ {
		var chromosome []domain.Gene
		totalPrice := 0
		totalWeight := 0

		for {
			product := products[rand.Intn(len(products))]
			number := rand.Intn(3) + 1 // 至少買 1 個，最多 3 個

			// 預估加入這個產品後的總價與總重
			nextTotalPrice := totalPrice + product.Price*number
			nextTotalWeight := totalWeight + product.Weight*number

			if nextTotalPrice > Budget || nextTotalWeight > BagCapacity {
				break
			}

			// 加入 Gene
			chromosome = append(chromosome, NewProductRecommendationGene(product, number))
			totalPrice = nextTotalPrice
			totalWeight = nextTotalWeight

			// 有機率提前結束生成（讓個體長度多樣性）
			if rand.Float64() < 0.3 {
				break
			}
		}

		individuals[i] = NewProductRecommendationIndividual(chromosome)
	}

	return domain.NewPopulation(individuals)
}
