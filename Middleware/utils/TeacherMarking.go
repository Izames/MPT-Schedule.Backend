package utils

import "MPT-Schedule/Models"

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
				if lesson.SlotInserted != 0 {
					if lesson.NumLesson.Name != "" && lesson.DenLesson.Name != "" {
						TeacherMark(&lesson.NumLesson, lesson.NumLesson.DoubleTeacher, day, lesson.SlotInserted-1)
						TeacherMark(&lesson.DenLesson, lesson.DenLesson.DoubleTeacher, day, lesson.SlotInserted-1)
					} else if lesson.NumLesson.Name != "" {
						TeacherMark(&lesson.NumLesson, lesson.NumLesson.DoubleTeacher, day, lesson.SlotInserted-1)
					} else {
						TeacherMark(&lesson.DenLesson, lesson.DenLesson.DoubleTeacher, day, lesson.SlotInserted-1)
					}
				}
			}
		}
	}
}
func TeacherMark(lesson *Models.LessonModel, doubleTeacher bool, day, lessonNum int) {
	if doubleTeacher {
		lesson.Teacher.Week[day].Lessons[lessonNum] = false
		lesson.TeacherTwo.Week[day].Lessons[lessonNum] = false
	} else {
		lesson.Teacher.Week[day].Lessons[lessonNum] = false
	}
}
