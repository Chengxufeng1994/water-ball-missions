package infrastructure

import "github.com/Chengxufeng1994/water-ball-missions/chapter04/employee_database/internal/domain"

type RealEmployee struct {
	id           int
	name         string
	age          int
	subordinates []domain.Employee
}

var _ domain.Employee = (*RealEmployee)(nil)

func NewRealEmployee(id int, name string, age int, subordinates []domain.Employee) *RealEmployee {
	return &RealEmployee{
		id:           id,
		name:         name,
		age:          age,
		subordinates: subordinates,
	}
}

// ID implements domain.Employee.
func (r *RealEmployee) ID() int {
	return r.id
}

// Name implements domain.Employee.
func (r *RealEmployee) Name() string {
	return r.name
}

// Age implements domain.Employee.
func (r *RealEmployee) Age() int {
	return r.age
}

// Subordinates implements domain.Employee.
func (r *RealEmployee) GetSubordinates() []domain.Employee {
	return r.subordinates
}
