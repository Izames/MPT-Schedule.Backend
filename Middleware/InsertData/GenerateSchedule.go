package InsertData

import (
	"MPT-Schedule/Middleware/utils"
	"MPT-Schedule/Models"
)

func GenerateSchedule(group Models.GroupModel) {
	var daysForInsert []*Models.ScheduleDay
	var lessonsSemester1 []Models.LessonModel
	var lessonsSemester2 []Models.LessonModel
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
			daysForInsert = append(daysForInsert, &days[i])
		}
	}
	for _, lesson := range group.Lessons {
		if lesson.PerWeekS1 > 0 {
			lessonsSemester1 = append(lessonsSemester1, lesson)
		}
		if lesson.PerWeekS2 > 0 {
			lessonsSemester2 = append(lessonsSemester2, lesson)
		}
	}

	InsertInSemester(lessonsSemester1, daysForInsert, 1)
	schedule := Models.ScheduleModel{
		Group: group.Name,
	}
	schedule.Semester1.Monday = days[0]
	schedule.Semester1.Tuesday = days[1]
	schedule.Semester1.Wednesday = days[2]
	schedule.Semester1.Thursday = days[3]
	schedule.Semester1.Friday = days[4]
	schedule.Semester1.Saturday = days[5]
	//InsertInSemester(lessonsSemester2, daysForInsert)
	schedule.Semester2.Monday = days[0]
	schedule.Semester2.Tuesday = days[1]
	schedule.Semester2.Wednesday = days[2]
	schedule.Semester2.Thursday = days[3]
	schedule.Semester2.Friday = days[4]
	schedule.Semester2.Saturday = days[5]

	Models.Schedules = append(Models.Schedules, schedule)
}
