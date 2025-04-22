package utils

import "MPT-Schedule/Models"

func SortLessons(lessons []Models.LessonModel, rData *Models.RequestData) []Models.LessonModel {
	var sortedLessons []Models.LessonModel
	var highPriority []Models.LessonModel
	var lowPriority []Models.LessonModel
	for _, lesson := range lessons {
		if lesson.DoubleTeacher {
			if (len(lesson.Teacher.Builds) < len(rData.Builds)) || (len(lesson.TeacherTwo.Builds) < len(rData.Builds)) {
				highPriority = append(highPriority, lesson)
			} else {
				lowPriority = append(lowPriority, lesson)
			}
		} else {
			if len(lesson.Teacher.Builds) < len(rData.Builds) {
				highPriority = append(highPriority, lesson)
			} else {
				lowPriority = append(lowPriority, lesson)
			}
		}
	}
	for _, lesson := range highPriority {
		sortedLessons = append(sortedLessons, lesson)
	}
	for _, lesson := range lowPriority {
		sortedLessons = append(sortedLessons, lesson)
	}
	return sortedLessons
}
