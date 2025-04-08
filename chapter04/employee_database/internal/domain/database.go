package domain

type Database interface {
	GetEmployeeByID(id int) (Employee, error)
}
