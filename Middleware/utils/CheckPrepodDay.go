package utils

import "MPT-Schedule/Models"

func CheckPrepodDay(freeDay Models.LessonDayModel, lessonsDay Models.LessonDayModel, result bool) Models.LessonDayModel {
	if result {
		return lessonsDay
	} else {
		return freeDay
	}
}
