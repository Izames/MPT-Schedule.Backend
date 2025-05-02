package InsertData

import (
	"MPT-Schedule/Middleware/utils"
	"MPT-Schedule/Models"
	"fmt"
)

func InsertInSemester(lessons []Models.LessonModel, daysForInsert []Models.ScheduleDay, group string, semester int, rData *Models.RequestData) []Models.ScheduleDay {
	attempts := 0
	BrokenTry := false
	for {
		if attempts >= 9 {
			BrokenTry = true
		}
		var LessonsWeek [][]Models.LessonModelND
		var BrokenLessons []Models.LessonModelND
		for range daysForInsert {
			LessonsWeek = append(LessonsWeek, []Models.LessonModelND{})
		}
		Lessons := utils.RandomLessonSplitter(lessons, daysForInsert)
		//создаем дубликат недели чтобы просто скидать туда пары в кучу и потом распределить
		j := 0
		maxCycles := len(LessonsWeek) + 1
		currentCycles := 0
		for len(Lessons) > 0 {

			if currentCycles > maxCycles {
				BrokenLessons = append(BrokenLessons, Lessons[0])
				Lessons = Lessons[1:]
				currentCycles = 0
				continue
			}

			if j == len(LessonsWeek) {
				j = 0
			}
			if Lessons[0].OneND {
				result, FarRange := utils.CheckLesson(&Lessons[0].DenLesson, daysForInsert[j].Day-1, daysForInsert[j].Building, rData, LessonsWeek[j])
				if result {
					if FarRange {
						Lessons[0].OneFarAlready = true
					}
					currentCycles = 0
				} else {
					currentCycles++
					j++
					continue
				}
			} else {
				result1, FarRange1 := utils.CheckLesson(&Lessons[0].DenLesson, daysForInsert[j].Day-1, daysForInsert[j].Building, rData, LessonsWeek[j])
				result2, FarRange2 := utils.CheckLesson(&Lessons[0].NumLesson, daysForInsert[j].Day-1, daysForInsert[j].Building, rData, LessonsWeek[j])
				if result1 && result2 {
					if FarRange1 || FarRange2 {
						Lessons[0].OneFarAlready = true
					}
					currentCycles = 0
				} else {
					currentCycles++
					j++
					continue
				}

			}
			LessonsWeek[j] = append(LessonsWeek[j], Lessons[0])
			Lessons = Lessons[1:]
			j++

		}
		if len(BrokenLessons) > 0 && !BrokenTry {
			attempts++
			continue
		}
		for _, lesson := range BrokenLessons {
			if lesson.OneND {
				rData.FilesErrors = append(rData.FilesErrors, fmt.Sprintf("В группу %s %d семестра не была вставлена пара %s", group, semester, lesson.NumLesson.Name))
			} else {
				rData.FilesErrors = append(rData.FilesErrors, fmt.Sprintf("В группу %s %d семестра не была вставлена пара %s по знаменателю и %s по числителю", group, semester, lesson.NumLesson.Name, lesson.DenLesson.Name))
			}
		}
		result := true
		for i := range LessonsWeek {
			result = LessonDayDistributor(LessonsWeek[i], daysForInsert[i], group, semester, daysForInsert[i].Day-1, rData, BrokenTry)
			if !result {
				break
			}
		}
		if result {
			break
		}
		attempts++
		if BrokenTry {
			break
		}
	}
	println("df")
	//начало алгоритма. Алгоритм перечисляет пары
	//for !(len(lessons) == 0 || attemps > 3) {
	//	lessonsCount := len(lessons)
	//	//берутся дни для вставления
	//	for j := range daysForInsert {
	//	LessonLoop:
	//		for c := 0; c < len(lessons); c++ {
	//			// Проверка на максимальное количество пар в день
	//			if len(daysForInsert[j].Lessons) == 5 {
	//				continue
	//			}
	//
	//			//проверка на какую пару внедрить
	//			switch {
	//			case daysForInsert[j].Lessons == nil:
	//				goBackOnStep := false
	//				randomLesson := rand.Intn(len(lessons))
	//				if utils.InsertDay(&lessons[randomLesson], &daysForInsert[j], NorDLessons, 0, group) {
	//					for k, l := range lessons {
	//						if l.PerWeek == 0 {
	//							lessons = append(lessons[:k], lessons[k+1:]...)
	//							goBackOnStep = true
	//						}
	//					}
	//					if goBackOnStep {
	//						c--
	//					}
	//					break LessonLoop
	//				}
	//
	//			case len(daysForInsert[j].Lessons) == 1:
	//				goBackOnStep := false
	//				randomLesson := rand.Intn(len(lessons))
	//				if utils.InsertDay(&lessons[randomLesson], &daysForInsert[j], NorDLessons, 1, group) {
	//					for k, l := range lessons {
	//						if l.PerWeek == 0 {
	//							lessons = append(lessons[:k], lessons[k+1:]...)
	//							goBackOnStep = true
	//						}
	//					}
	//					if goBackOnStep {
	//						c--
	//					}
	//					break LessonLoop
	//				}
	//			case len(daysForInsert[j].Lessons) == 2:
	//				goBackOnStep := false
	//				randomLesson := rand.Intn(len(lessons))
	//				if utils.InsertDay(&lessons[randomLesson], &daysForInsert[j], NorDLessons, 2, group) {
	//					for k, l := range lessons {
	//						if l.PerWeek == 0 {
	//							lessons = append(lessons[:k], lessons[k+1:]...)
	//							goBackOnStep = true
	//						}
	//					}
	//					if goBackOnStep {
	//						c--
	//					}
	//					break LessonLoop
	//				}
	//			case len(daysForInsert[j].Lessons) == 3:
	//				goBackOnStep := false
	//				randomLesson := rand.Intn(len(lessons))
	//				if utils.InsertDay(&lessons[randomLesson], &daysForInsert[j], NorDLessons, 3, group) {
	//					for k, l := range lessons {
	//						if l.PerWeek == 0 {
	//							lessons = append(lessons[:k], lessons[k+1:]...)
	//							goBackOnStep = true
	//						}
	//					}
	//					if goBackOnStep {
	//						c--
	//					}
	//					break LessonLoop
	//				}
	//			case len(daysForInsert[j].Lessons) == 4:
	//				goBackOnStep := false
	//				randomLesson := rand.Intn(len(lessons))
	//				if utils.InsertDay(&lessons[randomLesson], &daysForInsert[j], NorDLessons, 4, group) {
	//					for k, l := range lessons {
	//						if l.PerWeek == 0 {
	//							lessons = append(lessons[:k], lessons[k+1:]...)
	//							goBackOnStep = true
	//						}
	//					}
	//					if goBackOnStep {
	//						c--
	//					}
	//					break LessonLoop
	//				}
	//			case len(daysForInsert[j].Lessons) == 5:
	//				continue
	//			}
	//
	//		}
	//	}
	//	if len(lessons) == lessonsCount {
	//		attemps++
	//		if attemps == 3 {
	//			for l := range lessons {
	//				rData.FilesErrors = append(rData.FilesErrors, fmt.Sprintf("ошибка при генерации расписания, у %s не были вставлены пары: %s в %d семестре. Пар осталось: %f", group, lessons[l].Name, semester, lessons[l].PerWeek))
	//			}
	//		}
	//	}
	//}

	return daysForInsert
}
