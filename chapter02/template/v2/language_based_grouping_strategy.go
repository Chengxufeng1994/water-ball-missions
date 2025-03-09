package main

import "github.com/Chengxufeng1994/water-ball-missions/chapter02/template/model"

type LanguageBasedGroupingStrategy struct {
	CutBasedGroupingStrategy
}

func NewLanguageBasedGroupingStrategy() *LanguageBasedGroupingStrategy {
	return &LanguageBasedGroupingStrategy{}
}

var _ model.GroupingStrategy = (*LanguageBasedGroupingStrategy)(nil)

func (l *LanguageBasedGroupingStrategy) CutBy(student *model.Student) string {
	return student.Language
}

func (l *LanguageBasedGroupingStrategy) MeetMergeCriteria(nonFullGroup *model.Group, fullGroup *model.Group) bool {
	return nonFullGroup.Students[0].Language == fullGroup.Students[0].Language
}
