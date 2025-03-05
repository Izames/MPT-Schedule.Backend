package utils

import "MPT-Schedule/Models"

func TeacherLessonsReserved(day Models.LessonDayModel) int {
	count := 0
	for _, lesson := range day.Lessons {
		if lesson == false {
			count++
		}
	}
	return count
}
