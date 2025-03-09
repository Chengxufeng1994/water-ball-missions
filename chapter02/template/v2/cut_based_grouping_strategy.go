package main

import (
	"github.com/Chengxufeng1994/water-ball-missions/chapter02/template/model"
)

const GROUP_MIN_SIZE = 6

type CutBasedGroupingStrategy struct {
}

func NewCutBasedGroupingStrategy() *CutBasedGroupingStrategy {
	return &CutBasedGroupingStrategy{}
}

var _ model.GroupingStrategy = (*CutBasedGroupingStrategy)(nil)

func (c *CutBasedGroupingStrategy) Group(students []*model.Student) []*model.Group {
	firstCutMap := make(map[string]*model.Group)
	for _, student := range students {
		key := c.CutBy(student)
		if _, ok := firstCutMap[key]; !ok {
			firstCutMap[key] = model.NewGroup()
		}
		group := firstCutMap[key]
		group.AddStudents(student)
		firstCutMap[key] = group
	}

	secondCutGroups := make([]*model.Group, 0)
	for _, group := range firstCutMap {
		groups := group.SplitGroupBySize(GROUP_MIN_SIZE)
		secondCutGroups = append(secondCutGroups, groups...)
	}

	nonFullGroups := make([]*model.Group, 0)
	fullGroups := make([]*model.Group, 0)
	for _, group := range secondCutGroups {
		if group.Size() < GROUP_MIN_SIZE {
			nonFullGroups = append(nonFullGroups, group)
		} else {
			fullGroups = append(fullGroups, group)
		}
	}

	if len(fullGroups) == 0 {
		return nonFullGroups
	}

	for _, nonFullGroup := range nonFullGroups {
		for _, fullGroup := range fullGroups {
			if c.MeetMergeCriteria(nonFullGroup, fullGroup) {
				fullGroup.AddStudents(nonFullGroup.Students...)
			}
		}
	}

	return fullGroups
}

func (c CutBasedGroupingStrategy) CutBy(student *model.Student) string {
	return student.Language
}

func (c CutBasedGroupingStrategy) MeetMergeCriteria(nonFullGroup *model.Group, fullGroup *model.Group) bool {
	return true
}
