package domain

type Employee interface {
	ID() int
	Name() string
	Age() int
	GetSubordinates() []Employee
}
