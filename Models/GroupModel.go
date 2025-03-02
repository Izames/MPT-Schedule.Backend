package Models

type GroupModel struct {
	Name      string
	Monday    Day
	Tuesday   Day
	Wednesday Day
	Thursday  Day
	Friday    Day
	Saturday  Day
	Lessons   []LessonModel
}

type Day struct {
	Build   string
	UnSCDay bool
}
