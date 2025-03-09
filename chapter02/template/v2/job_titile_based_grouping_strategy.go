package main

import "github.com/Chengxufeng1994/water-ball-missions/chapter02/template/model"

type JobTitleBasedGroupingStrategy struct {
	*CutBasedGroupingStrategy
}

var _ model.GroupingStrategy = (*JobTitleBasedGroupingStrategy)(nil)

func (j JobTitleBasedGroupingStrategy) CutBy(student *model.Student) string {
	return student.JobTitle
}

func (j JobTitleBasedGroupingStrategy) MeetMergeCriteria(nonFullGroup *model.Group, fullGroup *model.Group) bool {
	return nonFullGroup.Students[0].JobTitle == fullGroup.Students[0].JobTitle
}
