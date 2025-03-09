package main

import "github.com/Chengxufeng1994/water-ball-missions/chapter02/template/model"

type TimeslotsBasedGroupingStrategy struct{}

var _ model.GroupingStrategy = (*TimeslotsBasedGroupingStrategy)(nil)

func (t *TimeslotsBasedGroupingStrategy) Group(students []*model.Student) []*model.Group {
	panic("unimplemented")
}
