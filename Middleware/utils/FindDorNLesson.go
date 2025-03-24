package utils

import (
	"MPT-Schedule/Models"
)

// поиск пары для числителя/знаменателя
func FindDorNLesson(DnNLessons []*Models.LessonModel, lessonNum, weekDay int, build string, startLesson *Models.LessonModel) *Models.LessonModel {
	for _, lesson := range DnNLessons {
		if lesson.DoubleTeacher {
			if !Contains(build, lesson.Teacher.Builds) &&
				!Contains(build, lesson.TeacherTwo.Builds) {
				continue
			}
			if TeacherLessonsReserved(lesson.Teacher.Week[weekDay]) == lesson.Teacher.LessonsInDay ||
				TeacherLessonsReserved(lesson.TeacherTwo.Week[weekDay]) == lesson.TeacherTwo.LessonsInDay {
				continue
			}
			if !lesson.Teacher.Week[weekDay].Lessons[lessonNum] && !lesson.TeacherTwo.Week[weekDay].Lessons[lessonNum] {
				continue
			}
		} else {
			if !Contains(build, lesson.Teacher.Builds) {
				continue
			}
			if TeacherLessonsReserved(lesson.Teacher.Week[weekDay]) == lesson.Teacher.LessonsInDay {
				continue
			}
			if !lesson.Teacher.Week[weekDay].Lessons[lessonNum] {
				continue
			}
		}
		if startLesson == lesson {
			continue
		}
		return lesson
	}
	return nil
}
