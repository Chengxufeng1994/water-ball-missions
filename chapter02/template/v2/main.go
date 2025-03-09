package main

import (
	"fmt"

	"github.com/Chengxufeng1994/water-ball-missions/chapter02/template/model"
)

func main() {
	strategy := NewLanguageBasedGroupingStrategy()
	students := []*model.Student{
		model.NewStudent("John", 3, "Golang", "Software Engineer", []int{1, 2, 3}),
		model.NewStudent("Jane", 3, "Java", "Software Engineer", []int{1, 2, 3}),
	}
	groups := strategy.Group(students)
	for _, group := range groups {
		fmt.Println(group.Students)
	}
}
