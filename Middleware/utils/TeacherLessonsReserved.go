package utils

import "MPT-Schedule/Models"

func TeacherLessonsReserved(day Models.LessonDayModel) (int, bool, bool) {
	count := 0
	//уроки 1 и 5, противоположности
	FarRange := false
	//свободен только один край (1 или 5 пара)
	OnlyEndFree := false
	for _, lesson := range day.Lessons {
		if lesson == false {
			count++
		}
	}
	if day.Lessons[0] && day.Lessons[4] && !day.Lessons[1] && !day.Lessons[2] && !day.Lessons[3] {
		FarRange = true
	}
	if count == 4 && (day.Lessons[0] || day.Lessons[4]) {
		OnlyEndFree = true
	}
	return count, FarRange, OnlyEndFree
}
