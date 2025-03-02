package Models

type LessonModel struct {
	Name          string
	PerWeekS1     float32
	PerWeekS2     float32
	Teacher       TeacherModel
	TeacherTwo    TeacherModel
	DoubleTeacher bool
}
