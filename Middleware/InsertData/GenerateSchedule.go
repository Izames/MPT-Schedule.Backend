package InsertData

import (
	"MPT-Schedule/Middleware/utils"
	"MPT-Schedule/Models"
)

func GenerateSchedule(group Models.GroupModel) {
	var daysForInsert []Models.ScheduleDay
	days := []Models.ScheduleDay{
		{Day: 1}, // Понедельник
		{Day: 2}, // Вторник
		{Day: 3}, // Среда
		{Day: 4}, // Четверг
		{Day: 5}, // Пятница
		{Day: 6}, // Суббота
	}

	for i := range days {
		var result bool
		days[i], result = utils.FilterUnSKDays(group, days[i])
		if result {
			daysForInsert = append(daysForInsert, days[i])
		}
	}

	daysForInsert = InsertInSemester(group.LessonsS1, daysForInsert)
	schedule := Models.ScheduleModel{
		Group: group.Name,
	}
	for i, day := range days {
		for _, dayI := range daysForInsert {
			if dayI.Day == day.Day {
				days[i] = dayI
			}
		}
	}
	for _, day := range days {
		schedule.Semester1.Week = append(schedule.Semester1.Week, day)
	}
	daysForInsert = InsertInSemester(group.LessonsS2, daysForInsert)
	for i, day := range days {
		for _, dayI := range daysForInsert {
			if dayI.Day == day.Day {
				days[i] = dayI
			}
		}
	}
	for _, day := range days {
		schedule.Semester1.Week = append(schedule.Semester1.Week, day)
	}

	Models.Schedules = append(Models.Schedules, schedule)
}
