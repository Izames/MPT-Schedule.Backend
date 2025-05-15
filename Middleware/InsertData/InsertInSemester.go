package InsertData

import (
	"MPT-Schedule/Middleware/utils"
	"MPT-Schedule/Models"
	"fmt"
)

func InsertInSemester(lessons []Models.LessonModel, daysForInsert []Models.ScheduleDay, group string, semester int, rData *Models.RequestData) []Models.ScheduleDay {
	attempts := 0
	BrokenTry := false
	var LessonsWeek [][]Models.LessonModelND
	for {
		WeekPairLessons := make([][][]*Models.LessonModelND, len(daysForInsert))
		for i := range WeekPairLessons {
			WeekPairLessons[i] = make([][]*Models.LessonModelND, 5)
		}
		var allPairLessons [][][]*Models.LessonModelND
		LessonsWeek = make([][]Models.LessonModelND, 0)
		if attempts >= 9 {
			BrokenTry = true
		}
		for range daysForInsert {
			LessonsWeek = append(LessonsWeek, []Models.LessonModelND{})
		}
		Lessons := utils.RandomLessonSplitter(lessons, daysForInsert, rData)
		//создаем дубликат недели чтобы просто скидать туда пары в кучу и потом распределить
		for i := range daysForInsert {
			LessonSlotDistributor(Lessons, WeekPairLessons[i], daysForInsert[i])
		}
		j := 0
		maxCycles := len(LessonsWeek)
		currentCycles := 0
		LessonsLen := -1
		LessonsAttempts := 0
		for len(Lessons) > 0 {

			if currentCycles > maxCycles {

				Lessons = append(Lessons, Lessons[0])
				Lessons = Lessons[1:]
				currentCycles = 0
				if len(Lessons) == LessonsLen {
					LessonsAttempts++
				} else {
					LessonsLen = len(Lessons)
					LessonsAttempts = 0
				}
				if LessonsAttempts > 5 {
					break
				}
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
				} else {
					currentCycles++
					j++
					continue
				}
			}
			Able := false
			if len(LessonsWeek[j]) >= 1 {
				TempLessons := make([][]*Models.LessonModelND, 5)
				TempLessonsWeek := append(LessonsWeek[j], Lessons[0])
				LessonSlotDistributor(TempLessonsWeek, TempLessons, daysForInsert[j])
				Count := 0
				for i := range TempLessons {
					if len(TempLessons[i]) >= 1 {
						Count++
					} else {
						Count = 0
					}
					if Count >= len(TempLessonsWeek) {
						Able = true
					}
				}
			} else {
				Able = true
			}
			if Able {
				currentCycles = 0
				LessonsAttempts = 0
				LessonsWeek[j] = append(LessonsWeek[j], Lessons[0])
				Lessons = Lessons[1:]
				j++
			} else {
				currentCycles++
				j++
			}

		}
		if len(Lessons) > 0 && !BrokenTry {
			attempts++
			continue
		}
		for _, lesson := range Lessons {
			if lesson.OneND {
				rData.FilesErrors = append(rData.FilesErrors, fmt.Sprintf("В группу %s %d семестра не была вставлена пара %s", group, semester, lesson.NumLesson.Name))
				rData.Failure = true
			} else {
				rData.FilesErrors = append(rData.FilesErrors, fmt.Sprintf("В группу %s %d семестра не была вставлена пара %s по знаменателю и %s по числителю", group, semester, lesson.NumLesson.Name, lesson.DenLesson.Name))
				rData.Failure = true
			}
		}
		result := true
		for i := range LessonsWeek {
			var pairLessons [][]*Models.LessonModelND
			result, LessonsWeek[i], pairLessons = LessonDayDistributor(LessonsWeek[i], daysForInsert[i], group, semester, daysForInsert[i].Day-1, rData, BrokenTry)
			if !result && !BrokenTry {
				break
			} else {
				allPairLessons = append(allPairLessons, pairLessons)
			}
		}
		if result {
			for i := range allPairLessons {
				utils.TeacherMarking(allPairLessons[i], daysForInsert[i].Day-1)
			}
			break
		}
		attempts++
		if BrokenTry {
			for i := range allPairLessons {
				utils.TeacherMarking(allPairLessons[i], daysForInsert[i].Day-1)
			}
			break
		}
	}
	for i := range LessonsWeek {
		daysForInsert[i].Lessons = make([]Models.ScheduleLesson, 5)
		for c := range LessonsWeek[i] {
			if LessonsWeek[i][c].NumLesson.DoubleTeacher && LessonsWeek[i][c].DenLesson.DoubleTeacher {
				if LessonsWeek[i][c].SlotInserted == 0 {
					continue
				}
				daysForInsert[i].Lessons[LessonsWeek[i][c].SlotInserted-1] = Models.ScheduleLesson{
					NumLessonName: Models.NumDenLesson{
						LessonName: LessonsWeek[i][c].NumLesson.Name,
						Teacher:    LessonsWeek[i][c].NumLesson.Teacher.FIO + " " + LessonsWeek[i][c].NumLesson.TeacherTwo.FIO,
					},
					DenLessonName: Models.NumDenLesson{
						LessonName: LessonsWeek[i][c].DenLesson.Name,
						Teacher:    LessonsWeek[i][c].DenLesson.Teacher.FIO + " " + LessonsWeek[i][c].DenLesson.TeacherTwo.FIO,
					},
				}
			} else if LessonsWeek[i][c].NumLesson.DoubleTeacher && !LessonsWeek[i][c].DenLesson.DoubleTeacher {
				if LessonsWeek[i][c].SlotInserted == 0 {
					continue
				}
				daysForInsert[i].Lessons[LessonsWeek[i][c].SlotInserted-1] = Models.ScheduleLesson{
					NumLessonName: Models.NumDenLesson{
						LessonName: LessonsWeek[i][c].NumLesson.Name,
						Teacher:    LessonsWeek[i][c].NumLesson.Teacher.FIO + " " + LessonsWeek[i][c].NumLesson.TeacherTwo.FIO,
					},
					DenLessonName: Models.NumDenLesson{
						LessonName: LessonsWeek[i][c].DenLesson.Name,
						Teacher:    LessonsWeek[i][c].DenLesson.Teacher.FIO,
					},
				}
			} else if !LessonsWeek[i][c].NumLesson.DoubleTeacher && LessonsWeek[i][c].DenLesson.DoubleTeacher {
				if LessonsWeek[i][c].SlotInserted == 0 {
					continue
				}
				daysForInsert[i].Lessons[LessonsWeek[i][c].SlotInserted-1] = Models.ScheduleLesson{
					NumLessonName: Models.NumDenLesson{
						LessonName: LessonsWeek[i][c].NumLesson.Name,
						Teacher:    LessonsWeek[i][c].NumLesson.Teacher.FIO + " " + LessonsWeek[i][c].NumLesson.TeacherTwo.FIO,
					},
					DenLessonName: Models.NumDenLesson{
						LessonName: LessonsWeek[i][c].DenLesson.Name,
						Teacher:    LessonsWeek[i][c].DenLesson.Teacher.FIO + " " + LessonsWeek[i][c].DenLesson.TeacherTwo.FIO,
					},
				}
			} else {
				if LessonsWeek[i][c].SlotInserted == 0 {
					continue
				}
				daysForInsert[i].Lessons[LessonsWeek[i][c].SlotInserted-1] = Models.ScheduleLesson{
					NumLessonName: Models.NumDenLesson{
						LessonName: LessonsWeek[i][c].NumLesson.Name,
						Teacher:    LessonsWeek[i][c].NumLesson.Teacher.FIO,
					},
					DenLessonName: Models.NumDenLesson{
						LessonName: LessonsWeek[i][c].DenLesson.Name,
						Teacher:    LessonsWeek[i][c].DenLesson.Teacher.FIO,
					},
				}
			}

		}
	}
	return daysForInsert
}
