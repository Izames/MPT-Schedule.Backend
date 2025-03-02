package Models

type TeacherModel struct {
	FIO          string
	Monday       LessonDayModel
	Tuesday      LessonDayModel
	Wednesday    LessonDayModel
	Thursday     LessonDayModel
	Friday       LessonDayModel
	Saturday     LessonDayModel
	Window       bool
	LessonsInDay int
	Builds       []string
}

type LessonDayModel struct {
	FirstLesson  bool
	SecondLesson bool
	ThirdLesson  bool
	FourthLesson bool
	FifthLesson  bool
}
