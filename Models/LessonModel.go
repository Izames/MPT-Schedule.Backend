package Models

type LessonModel struct {
	Name          string
	PerWeek       float32
	Teacher       TeacherModel
	TeacherTwo    TeacherModel
	DoubleTeacher bool
}
