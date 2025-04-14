package productschedule

import "github.com/Chengxufeng1994/water-ball-missions/chapter05/genetic_alogrithm_suite/internal/domain"

type ProductScheduleIndividualFactory struct{}

var _ domain.IndividualFactory = (*ProductScheduleIndividualFactory)(nil)

func NewProductScheduleIndividualFactory() *ProductScheduleIndividualFactory {
	return &ProductScheduleIndividualFactory{}
}

func (p *ProductScheduleIndividualFactory) CreateIndividual(genes []domain.Gene) domain.Individual {
	return NewProductScheduleIndividual(genes)
}
