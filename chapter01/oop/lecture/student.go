package lecture

type Student struct {
	Name              string
	LectureAttendance []*LectureAttendance
}

func NewStudent(name string) *Student {
	return &Student{
		Name:              name,
		LectureAttendance: make([]*LectureAttendance, 0),
	}
}

func (s *Student) AddLectureAttendance(la *LectureAttendance) {
	s.LectureAttendance = append(s.LectureAttendance, la)
}

func (s *Student) RemoveLectureAttendance(la *LectureAttendance) {
	for i, v := range s.LectureAttendance {
		if v == la {
			s.LectureAttendance = append(s.LectureAttendance[:i], s.LectureAttendance[i+1:]...)
		}
	}
}
