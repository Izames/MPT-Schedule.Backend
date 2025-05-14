package InsertData

import "MPT-Schedule/Models"

func LessonSlotDistributor(lessons []Models.LessonModelND, pairLessons [][]*Models.LessonModelND, daysForInsert Models.ScheduleDay) {
	for l := range lessons {
		if lessons[l].OneND {
			if lessons[l].NumLesson.DoubleTeacher {
				for j := 0; j < 5; j++ {
					if lessons[l].NumLesson.Teacher.Week[daysForInsert.Day-1].Lessons[j] && lessons[l].NumLesson.TeacherTwo.Week[daysForInsert.Day-1].Lessons[j] {
						pairLessons[j] = append(pairLessons[j], &lessons[l])
					}
				}
			} else {
				for j := 0; j < 5; j++ {
					if lessons[l].NumLesson.Teacher.Week[daysForInsert.Day-1].Lessons[j] {
						pairLessons[j] = append(pairLessons[j], &lessons[l])
					}
				}
			}
		} else {
			if lessons[l].NumLesson.DoubleTeacher && !lessons[l].DenLesson.DoubleTeacher {
				for j := 0; j < 5; j++ {
					if lessons[l].NumLesson.Teacher.Week[daysForInsert.Day-1].Lessons[j] && lessons[l].NumLesson.TeacherTwo.Week[daysForInsert.Day-1].Lessons[j] && lessons[l].DenLesson.Teacher.Week[daysForInsert.Day-1].Lessons[j] {
						pairLessons[j] = append(pairLessons[j], &lessons[l])
					}
				}
			} else if !lessons[l].NumLesson.DoubleTeacher && lessons[l].DenLesson.DoubleTeacher {
				for j := 0; j < 5; j++ {
					if lessons[l].NumLesson.Teacher.Week[daysForInsert.Day-1].Lessons[j] && lessons[l].DenLesson.TeacherTwo.Week[daysForInsert.Day-1].Lessons[j] && lessons[l].DenLesson.Teacher.Week[daysForInsert.Day-1].Lessons[j] {
						pairLessons[j] = append(pairLessons[j], &lessons[l])
					}
				}
			} else if lessons[l].NumLesson.DoubleTeacher && lessons[l].DenLesson.DoubleTeacher {
				for j := 0; j < 5; j++ {
					if lessons[l].NumLesson.Teacher.Week[daysForInsert.Day-1].Lessons[j] && lessons[l].NumLesson.TeacherTwo.Week[daysForInsert.Day-1].Lessons[j] && lessons[l].DenLesson.Teacher.Week[daysForInsert.Day-1].Lessons[j] && lessons[l].DenLesson.TeacherTwo.Week[daysForInsert.Day-1].Lessons[j] {
						pairLessons[j] = append(pairLessons[j], &lessons[l])
					}
				}
			} else {
				for j := 0; j < 5; j++ {
					if !(lessons[l].DenLesson.Name == "") && !(lessons[l].NumLesson.Name == "") {
						if lessons[l].NumLesson.Teacher.Week[daysForInsert.Day-1].Lessons[j] && lessons[l].DenLesson.Teacher.Week[daysForInsert.Day-1].Lessons[j] {
							pairLessons[j] = append(pairLessons[j], &lessons[l])
						}
					} else if lessons[l].NumLesson.Name == "" {
						if lessons[l].DenLesson.Teacher.Week[daysForInsert.Day-1].Lessons[j] {
							pairLessons[j] = append(pairLessons[j], &lessons[l])
						}
					} else {
						if lessons[l].NumLesson.Teacher.Week[daysForInsert.Day-1].Lessons[j] {
							pairLessons[j] = append(pairLessons[j], &lessons[l])
						}
					}

				}
			}
		}
	}
}
