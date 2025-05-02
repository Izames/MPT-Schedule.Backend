package InsertData

import "MPT-Schedule/Models"

func LessonSlotDistributor(lessons []Models.LessonModelND, pairLessons [][]*Models.LessonModelND, daysForInsert Models.ScheduleDay) {
	for _, lesson := range lessons {
		if lesson.OneND {
			if lesson.NumLesson.DoubleTeacher {
				for j := 0; j < 5; j++ {
					if lesson.NumLesson.Teacher.Week[daysForInsert.Day-1].Lessons[j] && lesson.NumLesson.TeacherTwo.Week[daysForInsert.Day-1].Lessons[j] {
						pairLessons[j] = append(pairLessons[j], &lesson)
					}
				}
			} else {
				for j := 0; j < 5; j++ {
					if lesson.NumLesson.Teacher.Week[daysForInsert.Day-1].Lessons[j] {
						pairLessons[j] = append(pairLessons[j], &lesson)
					}
				}
			}
		} else {
			if lesson.NumLesson.DoubleTeacher && !lesson.DenLesson.DoubleTeacher {
				for j := 0; j < 5; j++ {
					if lesson.NumLesson.Teacher.Week[daysForInsert.Day-1].Lessons[j] && lesson.NumLesson.TeacherTwo.Week[daysForInsert.Day-1].Lessons[j] && lesson.DenLesson.Teacher.Week[daysForInsert.Day-1].Lessons[j] {
						pairLessons[j] = append(pairLessons[j], &lesson)
					}
				}
			} else if !lesson.NumLesson.DoubleTeacher && lesson.DenLesson.DoubleTeacher {
				for j := 0; j < 5; j++ {
					if lesson.NumLesson.Teacher.Week[daysForInsert.Day-1].Lessons[j] && lesson.DenLesson.TeacherTwo.Week[daysForInsert.Day-1].Lessons[j] && lesson.DenLesson.Teacher.Week[daysForInsert.Day-1].Lessons[j] {
						pairLessons[j] = append(pairLessons[j], &lesson)
					}
				}
			} else if lesson.NumLesson.DoubleTeacher && lesson.DenLesson.DoubleTeacher {
				for j := 0; j < 5; j++ {
					if lesson.NumLesson.Teacher.Week[daysForInsert.Day-1].Lessons[j] && lesson.NumLesson.TeacherTwo.Week[daysForInsert.Day-1].Lessons[j] && lesson.DenLesson.Teacher.Week[daysForInsert.Day-1].Lessons[j] && lesson.DenLesson.TeacherTwo.Week[daysForInsert.Day-1].Lessons[j] {
						pairLessons[j] = append(pairLessons[j], &lesson)
					}
				}
			} else {
				for j := 0; j < 5; j++ {
					if lesson.NumLesson.Teacher.Week[daysForInsert.Day-1].Lessons[j] && lesson.DenLesson.Teacher.Week[daysForInsert.Day-1].Lessons[j] {
						pairLessons[j] = append(pairLessons[j], &lesson)
					}
				}
			}
		}
	}
}
