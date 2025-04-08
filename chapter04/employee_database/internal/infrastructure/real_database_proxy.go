package infrastructure

import (
	"errors"
	"os"

	"github.com/Chengxufeng1994/water-ball-missions/chapter04/employee_database/internal/domain"
)

type RealDatabaseProxy struct {
	*RealDatabase
}

var _ domain.Database = (*RealDatabaseProxy)(nil)

func NewRealDatabaseProxy(database *RealDatabase) *RealDatabaseProxy {
	return &RealDatabaseProxy{
		RealDatabase: database,
	}
}

func (r *RealDatabaseProxy) GetEmployeeByID(id int) (domain.Employee, error) {
	if err := r.ProtectedPassword(); err != nil {
		return nil, err
	}
	return r.RealDatabase.GetEmployeeByID(id)
}

func (r *RealDatabaseProxy) ProtectedPassword() error {
	password := os.Getenv("PASSWORD")

	if password != "1qaz2wsx" {
		return errors.New("wrong password")
	}

	return nil
}
