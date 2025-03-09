package main

import "github.com/Chengxufeng1994/water-ball-missions/chapter02/template/model"

const GROUP_MIN_SIZE = 6

type LanguageBasedGroupingStrategy struct{}

var _ model.GroupingStrategy = (*LanguageBasedGroupingStrategy)(nil)

func (l *LanguageBasedGroupingStrategy) Group(students []*model.Student) []*model.Group {
	languageGroupMap := make(map[string]*model.Group)
	for _, student := range students {
		if _, ok := languageGroupMap[student.Language]; !ok {
			languageGroupMap[student.Language] = model.NewGroup()
		}
		group := languageGroupMap[student.Language]
		group.AddStudents(student)
		languageGroupMap[student.Language] = group
	}

	secondCutGroups := make([]*model.Group, 0)
	for _, group := range languageGroupMap {
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

	for _, nonFullGroup := range nonFullGroups {
		for _, group := range fullGroups {
			if group.Students[0].Language == nonFullGroup.Students[0].Language {
				group.AddStudents(nonFullGroup.Students...)
			}
		}
	}

	return fullGroups
}
