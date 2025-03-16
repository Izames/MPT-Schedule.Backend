package utils

import (
	"MPT-Schedule/Models"
	"math"
)

func FindNorDLessons(lessons []Models.LessonModel) []*Models.LessonModel {
	var NorDLessons []*Models.LessonModel
	for i, lesson := range lessons {
		_, frac := math.Modf(float64(lesson.PerWeek))
		if frac != 0 {
			NorDLessons = append(NorDLessons, &lessons[i])
		}
	}
	return NorDLessons
}
