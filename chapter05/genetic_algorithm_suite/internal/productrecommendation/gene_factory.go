package productrecommendation

import (
	"math/rand"

	"github.com/Chengxufeng1994/water-ball-missions/chapter05/genetic_alogrithm_suite/internal/domain"
)

type ProductRecommendationGeneFactory struct {
}

var _ domain.GeneFactory = (*ProductRecommendationGeneFactory)(nil)

func NewProductRecommendationGeneFactory() *ProductRecommendationGeneFactory {
	return &ProductRecommendationGeneFactory{}
}

func (p *ProductRecommendationGeneFactory) CreateGene(num int) domain.Gene {
	num = num % len(products)
	qty := rand.Intn(3) + 1
	return NewProductRecommendationGene(products[num], qty)
}
