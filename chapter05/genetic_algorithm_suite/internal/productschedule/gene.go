package productschedule

import (
	"github.com/Chengxufeng1994/water-ball-missions/chapter05/genetic_alogrithm_suite/internal/domain"
)

type ProductScheduleGene struct {
	ProductName string
	Duration    int
	Demand      int
	MachineID   string
	WorkerID    string
}

var _ domain.Gene = (*ProductScheduleGene)(nil)

func NewProductScheduleGene(ProductName string, Duration int, Demand int, MachineID string, WorkerID string) *ProductScheduleGene {
	return &ProductScheduleGene{
		ProductName: ProductName,
		Duration:    Duration,
		Demand:      Demand,
		MachineID:   MachineID,
		WorkerID:    WorkerID,
	}
}

func (p ProductScheduleGene) Clone() domain.Gene {
	return ProductScheduleGene{
		ProductName: p.ProductName,
		MachineID:   p.MachineID,
		WorkerID:    p.WorkerID,
		Duration:    p.Duration,
		Demand:      p.Demand,
	}
}
