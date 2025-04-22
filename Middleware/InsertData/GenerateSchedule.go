package InsertData

import (
	"MPT-Schedule/Middleware/utils"
	"MPT-Schedule/Models"
)

func GenerateSchedule(group Models.GroupModel, rData *Models.RequestData) {
	//сортировка дней для заполнения
	var daysForInsert []Models.ScheduleDay
	days := []Models.ScheduleDay{
		{Day: 1}, // Понедельник
		{Day: 2}, // Вторник
		{Day: 3}, // Среда
		{Day: 4}, // Четверг
		{Day: 5}, // Пятница
		{Day: 6}, // Суббота
	}
	var daysForInsert2 []Models.ScheduleDay
	days2 := []Models.ScheduleDay{
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
		days2[i], result = utils.FilterUnSKDays(group, days[i])
		if result {
			daysForInsert = append(daysForInsert, days[i])
			daysForInsert2 = append(daysForInsert2, days2[i])
		}
	}
	//формирование семестра 1 и 2
	daysForInsert = InsertInSemester(group.LessonsS1, daysForInsert, group.Name, 1, rData)
	//название расписания
	schedule := Models.ScheduleModel{
		Group: group.Name,
	}
	schedule.List = group.ListName
	schedule.File = group.FileName
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

	daysForInsert2 = InsertInSemester(group.LessonsS2, daysForInsert2, group.Name, 2, rData)
	for i, day := range days2 {
		for _, dayI := range daysForInsert2 {
			if dayI.Day == day.Day {
				days2[i] = dayI
			}
		}
	}
	for _, day := range days2 {
		schedule.Semester2.Week = append(schedule.Semester2.Week, day)
	}

	rData.Schedules = append(rData.Schedules, schedule)
}
