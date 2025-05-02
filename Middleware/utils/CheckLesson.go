package utils

import (
	"MPT-Schedule/Models"
)

func CheckLesson(lesson *Models.LessonModel, weekDay int, build string, rData *Models.RequestData, LessonsDay []Models.LessonModelND) (bool, bool) {
	reservedLessons, FarRange, OnlyEndFree := TeacherLessonsReserved(lesson.Teacher.Week[weekDay])
	if len(LessonsDay) >= 5 {
		return false, FarRange
	}
	if lesson.DoubleTeacher {
		reservedLessons2, FarRange2, OnlyEndFree2 := TeacherLessonsReserved(lesson.TeacherTwo.Week[weekDay])
		if reservedLessons < reservedLessons2 {
			reservedLessons = reservedLessons2
		}
		if FarRange || FarRange2 {
			FarRange = true
		}
		if OnlyEndFree || OnlyEndFree2 {
			OnlyEndFree = true
		}
	}
	if len(LessonsDay) > 0 {
		if LessonsDay[0].OneFarAlready && FarRange {
			return false, FarRange
		}
	}
	if lesson.DoubleTeacher {
		if !Contains(build, lesson.Teacher.Builds) &&
			!Contains(build, lesson.TeacherTwo.Builds) {
			return false, FarRange
		}
		if lesson.Teacher.Week[weekDay].Build != "" {
			if build != "Дистанционно" && lesson.Teacher.Week[weekDay].Build != build {
				return false, FarRange
			}
		} else {
			lesson.Teacher.Week[weekDay].Build = build
		}
		if lesson.TeacherTwo.Week != nil && lesson.TeacherTwo.Week[weekDay].Build != "" {
			if lesson.TeacherTwo.Week[weekDay].Build != build && build != "Дистанционно" {
				return false, FarRange
			}
		} else {
			lesson.Teacher.Week[weekDay].Build = build
		}
		reservedLessons2, _, _ := TeacherLessonsReserved(lesson.TeacherTwo.Week[weekDay])
		if reservedLessons == lesson.Teacher.LessonsInDay || reservedLessons2 == lesson.TeacherTwo.LessonsInDay {
			return false, FarRange
		}
	} else {
		if !Contains(build, lesson.Teacher.Builds) {
			return false, FarRange
		}
		if lesson.Teacher.Week[weekDay].Build != "" && build != "Дистанционно" {
			if lesson.Teacher.Week[weekDay].Build != build {
				return false, FarRange
			}
		} else {
			lesson.Teacher.Week[weekDay].Build = build
		}
		if reservedLessons == lesson.Teacher.LessonsInDay {
			return false, FarRange
		}
	}
	//проверка чтоб один препод не больше двух раз в неделю и если у него свободна только одна пара, то не лез
	repeats := 0
	for _, lessonDay := range LessonsDay {
		if lessonDay.OneND {
			if lessonDay.NumLesson.DoubleTeacher {
				if lesson.DoubleTeacher {
					if lesson.Teacher.FIO == lessonDay.NumLesson.Teacher.FIO || lesson.Teacher.FIO == lessonDay.NumLesson.TeacherTwo.FIO {
						repeats++
					}
					if lesson.TeacherTwo.FIO == lessonDay.NumLesson.Teacher.FIO || lesson.TeacherTwo.FIO == lessonDay.NumLesson.TeacherTwo.FIO {
						repeats++
					}
				} else {
					if lesson.Teacher.FIO == lessonDay.NumLesson.Teacher.FIO || lesson.Teacher.FIO == lessonDay.NumLesson.TeacherTwo.FIO {
						repeats++
					}
				}
			} else {
				if lesson.DoubleTeacher {
					if lesson.Teacher.FIO == lessonDay.NumLesson.Teacher.FIO || lesson.TeacherTwo.FIO == lessonDay.NumLesson.TeacherTwo.FIO {
						repeats++
					}
				} else {
					if lesson.Teacher.FIO == lessonDay.NumLesson.Teacher.FIO {
						repeats++
					}
				}
			}
		} else {
			if lessonDay.NumLesson.DoubleTeacher {
				if lesson.DoubleTeacher {
					if lesson.Teacher.FIO == lessonDay.NumLesson.Teacher.FIO || lesson.Teacher.FIO == lessonDay.NumLesson.TeacherTwo.FIO {
						repeats++
					}
					if lesson.TeacherTwo.FIO == lessonDay.NumLesson.Teacher.FIO || lesson.TeacherTwo.FIO == lessonDay.NumLesson.TeacherTwo.FIO {
						repeats++
					}
				} else {
					if lesson.Teacher.FIO == lessonDay.NumLesson.Teacher.FIO || lesson.Teacher.FIO == lessonDay.NumLesson.TeacherTwo.FIO {
						repeats++
					}
				}
			} else {
				if lesson.DoubleTeacher {
					if lesson.Teacher.FIO == lessonDay.NumLesson.Teacher.FIO || lesson.TeacherTwo.FIO == lessonDay.NumLesson.TeacherTwo.FIO {
						repeats++
					}
				} else {
					if lesson.Teacher.FIO == lessonDay.NumLesson.Teacher.FIO {
						repeats++
					}
				}
			}
			if lessonDay.DenLesson.DoubleTeacher {
				if lesson.DoubleTeacher {
					if lesson.Teacher.FIO == lessonDay.DenLesson.Teacher.FIO || lesson.Teacher.FIO == lessonDay.DenLesson.TeacherTwo.FIO {
						repeats++
					}
					if lesson.TeacherTwo.FIO == lessonDay.DenLesson.Teacher.FIO || lesson.TeacherTwo.FIO == lessonDay.DenLesson.TeacherTwo.FIO {
						repeats++
					}
				} else {
					if lesson.Teacher.FIO == lessonDay.DenLesson.Teacher.FIO || lesson.Teacher.FIO == lessonDay.DenLesson.TeacherTwo.FIO {
						repeats++
					}
				}
			} else {
				if lesson.DoubleTeacher {
					if lesson.Teacher.FIO == lessonDay.DenLesson.Teacher.FIO || lesson.TeacherTwo.FIO == lessonDay.DenLesson.TeacherTwo.FIO {
						repeats++
					}
				} else {
					if lesson.Teacher.FIO == lessonDay.DenLesson.Teacher.FIO {
						repeats++
					}
				}
			}
		}
	}
	if repeats > 1 {
		return false, FarRange
	}
	if (repeats == 1 && reservedLessons == 4) || (repeats == 1 && FarRange) {
		return false, FarRange
	}
	if len(LessonsDay) > 0 {
		if LessonsDay[0].OneFarAlready && OnlyEndFree {
			return false, FarRange
		}
	}
	return true, FarRange
}
