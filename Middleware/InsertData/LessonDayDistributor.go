package InsertData

import (
	"MPT-Schedule/Models"
	"fmt"
)

// LessonDayDistributor класс, который берет массив пар подготовленных для дня недели и распределяет. К примеру с 1 по 4, или со 2 по 5
func LessonDayDistributor(lessons []Models.LessonModelND, dayForInsert Models.ScheduleDay, group string, semester, day int, rData *Models.RequestData, BrokenTry bool) (bool, []Models.LessonModelND) {
	PairCount := len(lessons)
	if len(lessons) == 0 {
		if BrokenTry {
			rData.FilesErrors = append(rData.FilesErrors, fmt.Sprintf("Ошибка! у группы %s %d семестра в %d день недели нету пар", group, semester, day+1))
			rData.Failure = true
		}
		return false, nil
	}
	pairLessons := make([][]*Models.LessonModelND, 5)
	LessonSlotDistributor(lessons, pairLessons, dayForInsert)
	startSlot := 0 //0 - это первый урок 4 - это пятый урок
	//5 - 0
	//4 - 1
	//3 - 2
	//2 - 3
	FailAttempt := true
	for startSlot = 0; startSlot <= 5-PairCount; startSlot++ {
		if Iterate(pairLessons, startSlot, PairCount, group) {
			FailAttempt = false
			break
		}
	}
	if FailAttempt {
		if BrokenTry {
			rData.FilesErrors = append(rData.FilesErrors, fmt.Sprintf("ошибка! не удалось найти верной комбинации пар для группы %s %d семестра %d-го дня недели", group, semester, day+1))
			rData.Failure = true
		}
		return false, nil
	} else {
		TeacherMarking(pairLessons, day)
	}
	println("dfdf")

	return true, lessons
}

func Iterate(pairLessons [][]*Models.LessonModelND, i, PairCount int, group string) bool {
	result := false
	log := fmt.Sprintf("количество пар: %d i: %d группа: %s", PairCount, i, group)
	println(log)
	for j := range pairLessons[i] {
		if !pairLessons[i][j].TryInserted {
			pairLessons[i][j].TryInserted = true
			pairLessons[i][j].SlotInserted = i + 1
			if i+1 < PairCount {
				if Iterate(pairLessons, i+1, PairCount, group) {
					result = true
					break
				} else {
					pairLessons[i][j].TryInserted = false
					pairLessons[i][j].SlotInserted = 0
				}
			} else {
				result = true
				break
			}
		}
	}
	return result
}
func TeacherMarking(pairLessons [][]*Models.LessonModelND, day int) {
	for i := range pairLessons {
		for _, lesson := range pairLessons[i] {
			if lesson.OneND {
				if lesson.NumLesson.DoubleTeacher && lesson.SlotInserted != 0 {
					lesson.NumLesson.Teacher.Week[day].Lessons[lesson.SlotInserted-1] = false
					lesson.NumLesson.TeacherTwo.Week[day].Lessons[lesson.SlotInserted-1] = false
				} else if lesson.SlotInserted != 0 {
					lesson.NumLesson.Teacher.Week[day].Lessons[lesson.SlotInserted-1] = false
				}
			} else {
				if lesson.NumLesson.DoubleTeacher && !lesson.DenLesson.DoubleTeacher && lesson.SlotInserted != 0 {
					lesson.NumLesson.Teacher.Week[day].Lessons[lesson.SlotInserted-1] = false
					lesson.NumLesson.TeacherTwo.Week[day].Lessons[lesson.SlotInserted-1] = false
				} else if !lesson.NumLesson.DoubleTeacher && lesson.DenLesson.DoubleTeacher && lesson.SlotInserted != 0 {
					lesson.DenLesson.Teacher.Week[day].Lessons[lesson.SlotInserted-1] = false
					lesson.DenLesson.TeacherTwo.Week[day].Lessons[lesson.SlotInserted-1] = false
				} else if lesson.NumLesson.DoubleTeacher && lesson.DenLesson.DoubleTeacher && lesson.SlotInserted != 0 {
					lesson.NumLesson.Teacher.Week[day].Lessons[lesson.SlotInserted-1] = false
					lesson.NumLesson.TeacherTwo.Week[day].Lessons[lesson.SlotInserted-1] = false
					lesson.DenLesson.Teacher.Week[day].Lessons[lesson.SlotInserted-1] = false
					lesson.DenLesson.TeacherTwo.Week[day].Lessons[lesson.SlotInserted-1] = false
				} else if !lesson.NumLesson.DoubleTeacher && !lesson.DenLesson.DoubleTeacher && lesson.SlotInserted != 0 {
					lesson.NumLesson.Teacher.Week[day].Lessons[lesson.SlotInserted-1] = false
					lesson.DenLesson.Teacher.Week[day].Lessons[lesson.SlotInserted-1] = false
				}
			}
		}
	}
	println("sd")
}
