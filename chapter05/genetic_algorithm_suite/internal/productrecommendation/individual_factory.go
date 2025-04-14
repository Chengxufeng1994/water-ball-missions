package productrecommendation

import "github.com/Chengxufeng1994/water-ball-missions/chapter05/genetic_alogrithm_suite/internal/domain"

type ProductRecommendationIndividualFactory struct{}

var _ domain.IndividualFactory = (*ProductRecommendationIndividualFactory)(nil)

func NewProductRecommendationIndividualFactory() *ProductRecommendationIndividualFactory {
	return &ProductRecommendationIndividualFactory{}
}

func (p *ProductRecommendationIndividualFactory) CreateIndividual(genes []domain.Gene) domain.Individual {
	return NewProductRecommendationIndividual(genes)
}
