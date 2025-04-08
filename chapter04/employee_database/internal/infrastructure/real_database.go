package infrastructure

import (
	"errors"
	"strconv"
	"strings"

	"github.com/Chengxufeng1994/water-ball-missions/chapter04/employee_database/internal/domain"
)

type RealDatabase struct {
}

var _ domain.Database = (*RealDatabase)(nil)

const rawData = `id name age subordinateIds
1 waterball 25
2 fixiabis 15 1,3
3 fong 7 1
4 cc 18 1,2,3
5 peterchen 3 1,4
6 handsomeboy 22 1
`

func NewRealDatabase() *RealDatabase {
	return &RealDatabase{}
}

func (r *RealDatabase) GetEmployeeByID(id int) (domain.Employee, error) {
	rows := strings.Split(rawData, "\n")
	rows = rows[1:]
	if id >= len(rows) {
		return nil, errors.New("employee not found")
	}

	row := rows[id-1]
	fields := strings.Fields(row)
	id, _ = strconv.Atoi(fields[0])
	name := fields[1]
	age, _ := strconv.Atoi(fields[2])

	var subordinatesList []int
	if len(fields) > 3 {
		subordinates := fields[3]
		subordinateIds := strings.SplitSeq(subordinates, ",")
		for subordinateId := range subordinateIds {
			id, _ := strconv.Atoi(subordinateId)
			subordinatesList = append(subordinatesList, id)
		}
	}

	return NewRealEmployeeProxy(r, id, name, age, subordinatesList), nil
}
