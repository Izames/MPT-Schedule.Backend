package InsertData

import (
	"MPT-Schedule/Middleware/utils"
	"MPT-Schedule/Models"
	"fmt"
	"math/rand"
)

func InsertInSemester(lessons []Models.LessonModel, daysForInsert []Models.ScheduleDay, group string, semester int, rData *Models.RequestData) []Models.ScheduleDay {
	//сортировка по доступным строениям
	lessons = utils.SortLessons(lessons, rData)
	NorDLessons := utils.FindNorDLessons(lessons)
	attemps := 0
	//начало алгоритма. Алгоритм перечисляет пары
	for !(len(lessons) == 0 || attemps > 3) {
		lessonsCount := len(lessons)
		//берутся дни для вставления
		for j := range daysForInsert {
		LessonLoop:
			for c := 0; c < len(lessons); c++ {
				// Проверка на максимальное количество пар в день
				if len(daysForInsert[j].Lessons) == 5 {
					continue
				}

				//проверка на какую пару внедрить
				switch {
				case daysForInsert[j].Lessons == nil:
					goBackOnStep := false
					randomLesson := rand.Intn(len(lessons))
					if utils.InsertDay(&lessons[randomLesson], &daysForInsert[j], NorDLessons, 0, group) {
						for k, l := range lessons {
							if l.PerWeek == 0 {
								lessons = append(lessons[:k], lessons[k+1:]...)
								goBackOnStep = true
							}
						}
						if goBackOnStep {
							c--
						}
						break LessonLoop
					}

				case len(daysForInsert[j].Lessons) == 1:
					goBackOnStep := false
					randomLesson := rand.Intn(len(lessons))
					if utils.InsertDay(&lessons[randomLesson], &daysForInsert[j], NorDLessons, 1, group) {
						for k, l := range lessons {
							if l.PerWeek == 0 {
								lessons = append(lessons[:k], lessons[k+1:]...)
								goBackOnStep = true
							}
						}
						if goBackOnStep {
							c--
						}
						break LessonLoop
					}
				case len(daysForInsert[j].Lessons) == 2:
					goBackOnStep := false
					randomLesson := rand.Intn(len(lessons))
					if utils.InsertDay(&lessons[randomLesson], &daysForInsert[j], NorDLessons, 2, group) {
						for k, l := range lessons {
							if l.PerWeek == 0 {
								lessons = append(lessons[:k], lessons[k+1:]...)
								goBackOnStep = true
							}
						}
						if goBackOnStep {
							c--
						}
						break LessonLoop
					}
				case len(daysForInsert[j].Lessons) == 3:
					goBackOnStep := false
					randomLesson := rand.Intn(len(lessons))
					if utils.InsertDay(&lessons[randomLesson], &daysForInsert[j], NorDLessons, 3, group) {
						for k, l := range lessons {
							if l.PerWeek == 0 {
								lessons = append(lessons[:k], lessons[k+1:]...)
								goBackOnStep = true
							}
						}
						if goBackOnStep {
							c--
						}
						break LessonLoop
					}
				case len(daysForInsert[j].Lessons) == 4:
					goBackOnStep := false
					randomLesson := rand.Intn(len(lessons))
					if utils.InsertDay(&lessons[randomLesson], &daysForInsert[j], NorDLessons, 4, group) {
						for k, l := range lessons {
							if l.PerWeek == 0 {
								lessons = append(lessons[:k], lessons[k+1:]...)
								goBackOnStep = true
							}
						}
						if goBackOnStep {
							c--
						}
						break LessonLoop
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
					rData.FilesErrors = append(rData.FilesErrors, fmt.Sprintf("ошибка при генерации расписания, у %s не были вставлены пары: %s в %d семестре. Пар осталось: %f", group, lessons[l].Name, semester, lessons[l].PerWeek))
				}
			}
		}
	}

	return daysForInsert
}
