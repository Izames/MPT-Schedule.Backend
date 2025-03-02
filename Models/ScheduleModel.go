package Models

type ScheduleModel struct {
	Group     string
	Semester1 Semester
	Semester2 Semester
}
type Semester struct {
	Monday    ScheduleDay
	Tuesday   ScheduleDay
	Wednesday ScheduleDay
	Thursday  ScheduleDay
	Friday    ScheduleDay
	Saturday  ScheduleDay
}
type ScheduleDay struct {
	Day          int
	StudyingDay  bool
	Building     string
	FirstLesson  ScheduleLesson
	SecondLesson ScheduleLesson
	ThirdLesson  ScheduleLesson
	FourthLesson ScheduleLesson
	FifthLesson  ScheduleLesson
}
type ScheduleLesson struct {
	NumLessonName string
	DenLessonName string
	Teacher       string
}
