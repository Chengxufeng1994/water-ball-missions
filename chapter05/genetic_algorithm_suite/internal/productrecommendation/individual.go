package productrecommendation

import "github.com/Chengxufeng1994/water-ball-missions/chapter05/genetic_alogrithm_suite/internal/domain"

type Product struct {
	ID       int
	Name     string
	Price    int
	Weight   int
	Category string
}

var products = []Product{
	{
		ID:       1,
		Name:     "Product 1",
		Price:    100,
		Weight:   2,
		Category: "A",
	},
	{
		ID:       2,
		Name:     "Product 2",
		Price:    200,
		Weight:   3,
		Category: "A",
	},
	{
		ID:       3,
		Name:     "Product 3",
		Price:    150,
		Weight:   5,
		Category: "B",
	},
	{
		ID:       4,
		Name:     "Product 4",
		Price:    300,
		Weight:   4,
		Category: "B",
	},
	{
		ID:       5,
		Name:     "Product 5",
		Price:    180,
		Weight:   6,
		Category: "C",
	},
	{
		ID:       6,
		Name:     "Product 6",
		Price:    250,
		Weight:   7,
		Category: "C",
	},
}

type Preference struct {
	Category string
	Score    float64 // 喜好度百分比，例如 A: 0.8
}

var preferenceMap = map[string]float64{
	"A": 0.8,
	"B": 0.6,
	"C": 0.4,
}

const (
	BagCapacity = 10
	Budget      = 700
)

type ProductRecommendationIndividual struct {
	chromosome []ProductRecommendationGene
	fitness    float64
}

var _ domain.Individual = (*ProductRecommendationIndividual)(nil)

func NewProductRecommendationIndividual(genes []domain.Gene) *ProductRecommendationIndividual {
	var chromosome []ProductRecommendationGene
	for _, gene := range genes {
		if g, ok := gene.(*ProductRecommendationGene); ok {
			chromosome = append(chromosome, *g)
		} else {
			panic("invalid gene type for ProductRecommendationIndividual")
		}
	}
	individual := &ProductRecommendationIndividual{
		chromosome: chromosome,
		fitness:    0,
	}
	individual.CalculateFitness()
	return individual
}

func (p *ProductRecommendationIndividual) Chromosome() []domain.Gene {
	genes := make([]domain.Gene, len(p.chromosome))
	for i, gene := range p.chromosome {
		genes[i] = &gene
	}
	return genes
}

func (p *ProductRecommendationIndividual) CalculateFitness() {
	totalPrice := 0.0
	for _, gene := range p.chromosome {
		totalPrice += float64(gene.Quantity) * float64(gene.Product.Price)
	}

	totalWeight := 0.0
	for i, gene := range p.chromosome {
		totalWeight += float64(gene.Quantity) * float64(products[i].Weight)
	}

	if totalPrice > float64(Budget) || totalWeight > float64(BagCapacity) {
		p.fitness = 0
	}

	totalPreference := 0.0
	for _, gene := range p.chromosome {
		score := preferenceMap[gene.Product.Category]
		totalPreference += float64(gene.Quantity) * score
	}

	p.fitness = totalPreference
}

func (p *ProductRecommendationIndividual) Fitness() float64 {
	return p.fitness
}
