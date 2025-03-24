package utils

import "MPT-Schedule/Models"

func InsertDay(lesson *Models.LessonModel, dayInsert *Models.ScheduleDay, NorDLessons []*Models.LessonModel, lessonNum int, group string) bool {
	//проверка на условий вставки
	if !CheckLesson(lesson, lessonNum, dayInsert.Day-1, dayInsert.Building) {
		return false
	}
	var day Models.ScheduleLesson
	//проверка на то как его вставить
	if !HasFractPart(lesson.PerWeek) {
		day = CreateDay(lesson, lesson)
		lesson.Teacher.Week[dayInsert.Day-1].Lessons[lessonNum] = false
		lesson.Teacher.Week[dayInsert.Day-1].Build = dayInsert.Building
		if lesson.DoubleTeacher {
			lesson.TeacherTwo.Week[dayInsert.Day-1].Lessons[lessonNum] = false
			lesson.TeacherTwo.Week[dayInsert.Day-1].Build = dayInsert.Building
		}
		lesson.PerWeek--
	} else {
		DenLesson := FindDorNLesson(NorDLessons, lessonNum, dayInsert.Day-1, dayInsert.Building, lesson)
		if DenLesson != nil {
			day = CreateDay(lesson, DenLesson)
			lesson.Teacher.Week[dayInsert.Day-1].Lessons[lessonNum] = false
			DenLesson.Teacher.Week[dayInsert.Day-1].Lessons[lessonNum] = false
			lesson.Teacher.Week[dayInsert.Day-1].Build = dayInsert.Building
			DenLesson.Teacher.Week[dayInsert.Day-1].Build = dayInsert.Building
			if lesson.DoubleTeacher {
				lesson.TeacherTwo.Week[dayInsert.Day-1].Lessons[lessonNum] = false
				lesson.TeacherTwo.Week[dayInsert.Day-1].Build = dayInsert.Building
			}
			if DenLesson.DoubleTeacher {
				DenLesson.TeacherTwo.Week[dayInsert.Day-1].Lessons[lessonNum] = false
				DenLesson.TeacherTwo.Week[dayInsert.Day-1].Build = dayInsert.Building
			}
			lesson.PerWeek -= 0.5
			DenLesson.PerWeek -= 0.5
		} else {
			if len(NorDLessons) == 1 {
				day = CreateDay(lesson, DenLesson)
				DenLesson.Teacher.Week[dayInsert.Day-1].Lessons[lessonNum] = false
				lesson.Teacher.Week[dayInsert.Day-1].Build = dayInsert.Building
				if lesson.DoubleTeacher {
					lesson.TeacherTwo.Week[dayInsert.Day-1].Lessons[lessonNum] = false
					lesson.TeacherTwo.Week[dayInsert.Day-1].Build = dayInsert.Building
				}
				lesson.PerWeek -= 0.5
			}

		}
	}
	dayInsert.Lessons = append(dayInsert.Lessons, day)
	return true
}
