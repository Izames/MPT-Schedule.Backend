package utils

import (
	"MPT-Schedule/Models"
)

func CheckLesson(lesson *Models.LessonModel, lessonNum, weekDay int, build string) bool {
	if lesson.DoubleTeacher {
		if !Contains(build, lesson.Teacher.Builds) &&
			!Contains(build, lesson.TeacherTwo.Builds) {
			return false
		}
		if lesson.Teacher.Week[weekDay].Build != "" {
			if lesson.Teacher.Week[weekDay].Build != build {
				return false
			}
		}
		if lesson.TeacherTwo.Week != nil && lesson.TeacherTwo.Week[weekDay].Build != "" {
			if lesson.TeacherTwo.Week[weekDay].Build != build {
				return false
			}
		}
		if TeacherLessonsReserved(lesson.Teacher.Week[weekDay]) == lesson.Teacher.LessonsInDay ||
			TeacherLessonsReserved(lesson.TeacherTwo.Week[weekDay]) == lesson.TeacherTwo.LessonsInDay {
			return false
		}
		if !lesson.Teacher.Week[weekDay].Lessons[lessonNum] && !lesson.TeacherTwo.Week[weekDay].Lessons[lessonNum] {
			return false
		}
	} else {
		if !Contains(build, lesson.Teacher.Builds) {
			return false
		}
		if lesson.Teacher.Week[weekDay].Build != "" {
			if lesson.Teacher.Week[weekDay].Build != build {
				return false
			}
		}
		if TeacherLessonsReserved(lesson.Teacher.Week[weekDay]) == lesson.Teacher.LessonsInDay {
			return false
		}
		if !lesson.Teacher.Week[weekDay].Lessons[lessonNum] {
			return false
		}
	}
	return true
}
