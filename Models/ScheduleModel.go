package Models

type ScheduleModel struct {
	File      string
	List      string
	Group     string
	Semester1 Semester
	Semester2 Semester
}
type Semester struct {
	Week []ScheduleDay
}
type ScheduleDay struct {
	Day         int
	StudyingDay bool
	Building    string
	Lessons     []ScheduleLesson
}
type ScheduleLesson struct {
	NumLessonName NumDenLesson
	DenLessonName NumDenLesson
}
type NumDenLesson struct {
	LessonName string
	Teacher    string
}
