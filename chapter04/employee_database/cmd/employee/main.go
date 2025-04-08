package main

import (
	"fmt"

	"github.com/Chengxufeng1994/water-ball-missions/chapter04/employee_database/internal/domain"
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/employee_database/internal/infrastructure"
)

func main() {
	database := infrastructure.NewRealDatabase()
	databaseProxy := infrastructure.NewRealDatabaseProxy(database)
	employee, err := databaseProxy.GetEmployeeByID(6)
	if err != nil {
		panic(err)
	}
	fmt.Println(employee.GetSubordinates())
}

type Client struct {
	database domain.Database
}
