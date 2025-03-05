package InsertData

import (
	"MPT-Schedule/Middleware/utils"
	"MPT-Schedule/Models"
)

func InsertInSemester(lessons []Models.LessonModel, daysForInsert []Models.ScheduleDay) []Models.ScheduleDay {
	lessons = utils.SortLessons(lessons)
	for j, _ := range daysForInsert {
		for _, lesson := range lessons {

			//проверка на препода
			if len(daysForInsert[j].Lessons) == 5 {
				continue
			}
			if lesson.DoubleTeacher {
				if !utils.Contains(daysForInsert[j].Building, lesson.Teacher.Builds) && !utils.Contains(daysForInsert[j].Building, lesson.TeacherTwo.Builds) {
					continue
				}
				if utils.TeacherLessonsReserved(lesson.Teacher.Week[daysForInsert[j].Day-1]) == lesson.Teacher.LessonsInDay || utils.TeacherLessonsReserved(lesson.TeacherTwo.Week[daysForInsert[j].Day-1]) == lesson.TeacherTwo.LessonsInDay {
					continue
				}
			} else {
				if !utils.Contains(daysForInsert[j].Building, lesson.Teacher.Builds) {
					continue
				}
				if utils.TeacherLessonsReserved(lesson.Teacher.Week[daysForInsert[j].Day-1]) == lesson.Teacher.LessonsInDay {
					continue
				}
			}

			//проверка на какую пару внедрить
			if daysForInsert[j].Lessons == nil {
				if lesson.Teacher.Week[daysForInsert[j].Day-1].Lessons[0] {
					//проверить на окно
					//проверить пару на числитель/знаменатель
					//в случае, если пара полная, то просто вставить и вычесть 1
					//в случае, если пара числитель/знаменатель, найти ей такую же пару и отнять 0.5
					//если пара = 0, то убрать ее из пар
				}
			}
			if len(daysForInsert[j].Lessons) == 1 {

			}
			if len(daysForInsert[j].Lessons) == 2 {

			}
			if len(daysForInsert[j].Lessons) == 3 {

			}
			if len(daysForInsert[j].Lessons) == 4 {

			}

		}
	}
	return daysForInsert
}
