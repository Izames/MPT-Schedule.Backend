package InsertData

import (
	"MPT-Schedule/Middleware/utils"
	"MPT-Schedule/Models"
	"fmt"
)

func InsertInSemester(lessons []Models.LessonModel, daysForInsert []Models.ScheduleDay, group string, semester int) []Models.ScheduleDay {
	lessons = utils.SortLessons(lessons)
	NorDLessons := utils.FindNorDLessons(lessons)
	attemps := 0
	for !(len(lessons) == 0 || attemps > 3) {
		lessonsCount := len(lessons)
		for j := range daysForInsert {
			for c := 0; c < len(lessons); c++ {
				// Проверка на максимальное количество пар в день
				if len(daysForInsert[j].Lessons) == 5 {
					continue
				}

				//проверка на какую пару внедрить
				switch {
				case daysForInsert[j].Lessons == nil:
					goBackOnStep := false
					if utils.InsertDay(&lessons[c], &daysForInsert[j], NorDLessons, 0) {
						for k, l := range lessons {
							if l.PerWeek == 0 {
								lessons = append(lessons[:k], lessons[k+1:]...)
								goBackOnStep = true
							}
						}
						if goBackOnStep {
							c--
						}
						continue
					}

				case len(daysForInsert[j].Lessons) == 1:
					goBackOnStep := false
					if utils.InsertDay(&lessons[c], &daysForInsert[j], NorDLessons, 1) {
						for k, l := range lessons {
							if l.PerWeek == 0 {
								lessons = append(lessons[:k], lessons[k+1:]...)
								goBackOnStep = true
							}
						}
						if goBackOnStep {
							c--
						}
						continue
					}
				case len(daysForInsert[j].Lessons) == 2:
					goBackOnStep := false
					if utils.InsertDay(&lessons[c], &daysForInsert[j], NorDLessons, 2) {
						for k, l := range lessons {
							if l.PerWeek == 0 {
								lessons = append(lessons[:k], lessons[k+1:]...)
								goBackOnStep = true
							}
						}
						if goBackOnStep {
							c--
						}
						continue
					}
				case len(daysForInsert[j].Lessons) == 3:
					goBackOnStep := false
					if utils.InsertDay(&lessons[c], &daysForInsert[j], NorDLessons, 3) {
						for k, l := range lessons {
							if l.PerWeek == 0 {
								lessons = append(lessons[:k], lessons[k+1:]...)
								goBackOnStep = true
							}
						}
						if goBackOnStep {
							c--
						}
						continue
					}
				case len(daysForInsert[j].Lessons) == 4:
					goBackOnStep := false
					if utils.InsertDay(&lessons[c], &daysForInsert[j], NorDLessons, 4) {
						for k, l := range lessons {
							if l.PerWeek == 0 {
								lessons = append(lessons[:k], lessons[k+1:]...)
								goBackOnStep = true
							}
						}
						if goBackOnStep {
							c--
						}
						continue
					}
				case len(daysForInsert[j].Lessons) == 5:
					continue
				}

			}
		}
		if len(lessons) == lessonsCount {
			attemps++
			if attemps == 3 {
				for l := range lessons {
					Models.FilesErrors = append(Models.FilesErrors, fmt.Sprintf("ошибка при генерации расписания, у %s не были вставлены пары: %s в %d семестре. Пар осталось: %f", group, lessons[l].Name, semester, lessons[l].PerWeek))
				}
			}
		}
	}

	return daysForInsert
}
