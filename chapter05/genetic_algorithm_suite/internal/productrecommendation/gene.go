package productrecommendation

import "github.com/Chengxufeng1994/water-ball-missions/chapter05/genetic_alogrithm_suite/internal/domain"

type ProductRecommendationGene struct {
	Product  Product
	Quantity int
}

var _ domain.Gene = (*ProductRecommendationGene)(nil)

func NewProductRecommendationGene(product Product, quantity int) *ProductRecommendationGene {
	return &ProductRecommendationGene{
		Product:  product,
		Quantity: quantity,
	}
}

func (p *ProductRecommendationGene) Clone() domain.Gene {
	return &ProductRecommendationGene{
		Product:  p.Product,
		Quantity: p.Quantity,
	}
}
