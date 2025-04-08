package infrastructure

import (
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/employee_database/internal/domain"
)

type RealEmployeeProxy struct {
	database domain.Database
	*RealEmployee
	SubordinateIDs []int
}

var _ domain.Employee = (*RealEmployeeProxy)(nil)

func NewRealEmployeeProxy(
	database domain.Database, id int, name string, age int, subordinates []int,
) *RealEmployeeProxy {
	return &RealEmployeeProxy{
		database:       database,
		RealEmployee:   NewRealEmployee(id, name, age, nil),
		SubordinateIDs: subordinates,
	}
}

func (r *RealEmployeeProxy) GetSubordinates() []domain.Employee {
	for _, id := range r.SubordinateIDs {
		employee, _ := r.database.GetEmployeeByID(id)
		r.RealEmployee.subordinates = append(r.RealEmployee.subordinates, employee)
	}
	return r.RealEmployee.subordinates
}
