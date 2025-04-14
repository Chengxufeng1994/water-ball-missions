package productschedule

import (
	"fmt"
	"math/rand"

	"github.com/Chengxufeng1994/water-ball-missions/chapter05/genetic_alogrithm_suite/internal/domain"
)

type ProductScheduleGeneFactory struct{}

var _ domain.GeneFactory = (*ProductScheduleGeneFactory)(nil)

func NewProductScheduleGeneFactory() *ProductScheduleGeneFactory {
	return &ProductScheduleGeneFactory{}
}

func (p *ProductScheduleGeneFactory) CreateGene(num int) domain.Gene {
	// 隨機挑選一個產品
	product := products[rand.Intn(len(products))]

	// 隨機選一台機器（M1、M2）
	machineID := fmt.Sprintf("M%d", machines[rand.Intn(len(machines))].ID)

	// 隨機選一位工人（W1~W4）
	workerID := fmt.Sprintf("W%d", workers[rand.Intn(len(workers))].ID)

	// 隨機產生一小部分需求（例如 1~10）
	randomDemand := rand.Intn(10) + 1

	return NewProductScheduleGene(product.Name, product.Duration, randomDemand, machineID, workerID)
}
