package utils

import "MPT-Schedule/Models"

func FilterUnSKDays(group Models.GroupModel, day Models.ScheduleDay) (Models.ScheduleDay, bool) {
	day.Building = group.Monday.Build
	if group.Monday.UnSCDay {
		day.StudyingDay = false
		return day, false
	} else {
		day.StudyingDay = true
		return day, true
	}
}
