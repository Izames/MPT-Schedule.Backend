package Models

type LessonModel struct {
	Name          string
	PerWeek       float32
	Teacher       TeacherModel
	TeacherTwo    TeacherModel
	DoubleTeacher bool
}

type LessonModelND struct {
	OneND         bool
	NumLesson     LessonModel
	DenLesson     LessonModel
	TryInserted   bool
	SlotInserted  int
	OneFarAlready bool
}
