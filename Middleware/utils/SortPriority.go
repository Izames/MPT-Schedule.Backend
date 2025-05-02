package utils

import "MPT-Schedule/Models"

func SortPriority(daysForInsert []Models.ScheduleDay, newLessons []Models.LessonModelND) []Models.LessonModelND {
	PrioritySortedLessons := make([][]Models.LessonModelND, len(daysForInsert)*4)
	for _, lesson := range newLessons {
		if lesson.OneND {
			if lesson.NumLesson.DoubleTeacher {
				priority := 0
				for _, day := range lesson.NumLesson.Teacher.Week {
					for _, insertDay := range daysForInsert {
						if day.Day == insertDay.Day && day.Build != "" {
							priority++
						}
					}
				}
				for _, day := range lesson.NumLesson.TeacherTwo.Week {
					for _, insertDay := range daysForInsert {
						if day.Day == insertDay.Day && day.Build != "" {
							priority++
						}
					}
				}
				PrioritySortedLessons[priority] = append(PrioritySortedLessons[priority], lesson)
			} else {
				priority := 0
				for _, day := range lesson.NumLesson.Teacher.Week {
					for _, insertDay := range daysForInsert {
						if day.Day == insertDay.Day && day.Build != "" {
							priority++
						}
					}
				}
				PrioritySortedLessons[priority] = append(PrioritySortedLessons[priority], lesson)
			}
		} else {
			if lesson.NumLesson.DoubleTeacher && !lesson.DenLesson.DoubleTeacher {
				priority := 0
				for _, day := range lesson.NumLesson.Teacher.Week {
					for _, insertDay := range daysForInsert {
						if day.Day == insertDay.Day && day.Build != "" {
							priority++
						}
					}
				}
				for _, day := range lesson.NumLesson.TeacherTwo.Week {
					for _, insertDay := range daysForInsert {
						if day.Day == insertDay.Day && day.Build != "" {
							priority++
						}
					}
				}
				for _, day := range lesson.DenLesson.Teacher.Week {
					for _, insertDay := range daysForInsert {
						if day.Day == insertDay.Day && day.Build != "" {
							priority++
						}
					}
				}
				PrioritySortedLessons[priority] = append(PrioritySortedLessons[priority], lesson)
			} else if !lesson.NumLesson.DoubleTeacher && lesson.DenLesson.DoubleTeacher {
				priority := 0
				for _, day := range lesson.NumLesson.Teacher.Week {
					for _, insertDay := range daysForInsert {
						if day.Day == insertDay.Day && day.Build != "" {
							priority++
						}
					}
				}
				for _, day := range lesson.DenLesson.Teacher.Week {
					for _, insertDay := range daysForInsert {
						if day.Day == insertDay.Day && day.Build != "" {
							priority++
						}
					}
				}
				for _, day := range lesson.DenLesson.TeacherTwo.Week {
					for _, insertDay := range daysForInsert {
						if day.Day == insertDay.Day && day.Build != "" {
							priority++
						}
					}
				}
				PrioritySortedLessons[priority] = append(PrioritySortedLessons[priority], lesson)
			} else if lesson.NumLesson.DoubleTeacher && lesson.DenLesson.DoubleTeacher {
				priority := 0
				for _, day := range lesson.NumLesson.Teacher.Week {
					for _, insertDay := range daysForInsert {
						if day.Day == insertDay.Day && day.Build != "" {
							priority++
						}
					}
				}
				for _, day := range lesson.NumLesson.TeacherTwo.Week {
					for _, insertDay := range daysForInsert {
						if day.Day == insertDay.Day && day.Build != "" {
							priority++
						}
					}
				}
				for _, day := range lesson.DenLesson.Teacher.Week {
					for _, insertDay := range daysForInsert {
						if day.Day == insertDay.Day && day.Build != "" {
							priority++
						}
					}
				}
				for _, day := range lesson.DenLesson.TeacherTwo.Week {
					for _, insertDay := range daysForInsert {
						if day.Day == insertDay.Day && day.Build != "" {
							priority++
						}
					}
				}
				PrioritySortedLessons[priority] = append(PrioritySortedLessons[priority], lesson)
			} else {
				priority := 0
				for _, day := range lesson.NumLesson.Teacher.Week {
					for _, insertDay := range daysForInsert {
						if day.Day == insertDay.Day && day.Build != "" {
							priority++
						}
					}
				}
				for _, day := range lesson.DenLesson.Teacher.Week {
					for _, insertDay := range daysForInsert {
						if day.Day == insertDay.Day && day.Build != "" {
							priority++
						}
					}
				}
				PrioritySortedLessons[priority] = append(PrioritySortedLessons[priority], lesson)
			}
		}
	}
	var EndLesson []Models.LessonModelND

	for i := len(PrioritySortedLessons) - 1; i >= 0; i-- {
		PriorityLessons := PrioritySortedLessons[i]
		for _, Lesson := range PriorityLessons {
			EndLesson = append(EndLesson, Lesson)
		}
	}
	return EndLesson
}
