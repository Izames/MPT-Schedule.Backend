package Models

type TeacherModel struct {
	FIO          string
	Week         []LessonDayModel
	Window       bool
	LessonsInDay int
	Builds       []string
}

type LessonDayModel struct {
	Day     int
	Build   string
	Lessons []bool
}
