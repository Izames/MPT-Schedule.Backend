package Models

type GroupModel struct {
	FileName  string
	ListName  string
	Name      string
	Week      []Day
	LessonsS1 []LessonModel
	LessonsS2 []LessonModel
}

type Day struct {
	DayNum  int
	Build   string
	UnSCDay bool
}
