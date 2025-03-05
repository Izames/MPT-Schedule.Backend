package utils

import "MPT-Schedule/Models"

func FilterUnSKDays(group Models.GroupModel, day Models.ScheduleDay) (Models.ScheduleDay, bool) {
	if group.Week[day.Day-1].UnSCDay {
		day.StudyingDay = false
		day.Building = group.Week[day.Day-1].Build
		return day, false
	} else {
		day.Building = group.Week[day.Day-1].Build
		day.StudyingDay = true
		return day, true
	}
}
