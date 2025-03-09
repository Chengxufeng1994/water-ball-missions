package model

type Student struct {
	Name               string
	Experience         int
	Language           string
	JobTitle           string
	AvailableTimeSlots []int
}

func NewStudent(name string, experience int, language string, jobTitle string, availableTimeSlots []int) *Student {
	return &Student{
		Name:               name,
		Experience:         experience,
		Language:           language,
		JobTitle:           jobTitle,
		AvailableTimeSlots: availableTimeSlots,
	}
}
