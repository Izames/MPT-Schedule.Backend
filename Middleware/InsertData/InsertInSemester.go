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
		LessonsWeek = make([][]Models.LessonModelND, 0)
		if attempts >= 9 {
			BrokenTry = true
		}
		var BrokenLessons []Models.LessonModelND
		for range daysForInsert {
			LessonsWeek = append(LessonsWeek, []Models.LessonModelND{})
		}
		Lessons := utils.RandomLessonSplitter(lessons, daysForInsert, rData)
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
				rData.Failure = true
			} else {
				rData.FilesErrors = append(rData.FilesErrors, fmt.Sprintf("В группу %s %d семестра не была вставлена пара %s по знаменателю и %s по числителю", group, semester, lesson.NumLesson.Name, lesson.DenLesson.Name))
				rData.Failure = true
			}
		}
		result := true
		for i := range LessonsWeek {
			result, LessonsWeek[i] = LessonDayDistributor(LessonsWeek[i], daysForInsert[i], group, semester, daysForInsert[i].Day-1, rData, BrokenTry)
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
