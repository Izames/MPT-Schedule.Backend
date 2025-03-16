package utils

import "MPT-Schedule/Models"

func CreateDay(numLesson, denLesson *Models.LessonModel) Models.ScheduleLesson {
	var day Models.ScheduleLesson
	if denLesson == nil {
		denLesson = &Models.LessonModel{
			Name: "",
			Teacher: Models.TeacherModel{
				FIO: "",
			},
			DoubleTeacher: false,
		}
	}
	if numLesson.DoubleTeacher && denLesson.DoubleTeacher {
		day = Models.ScheduleLesson{
			NumLessonName: Models.NumDenLesson{
				LessonName: numLesson.Name,
				Teacher:    numLesson.Teacher.FIO + " " + numLesson.TeacherTwo.FIO,
			},
			DenLessonName: Models.NumDenLesson{
				LessonName: denLesson.Name,
				Teacher:    denLesson.Teacher.FIO + " " + denLesson.TeacherTwo.FIO,
			},
		}
	} else if numLesson.DoubleTeacher {
		day = Models.ScheduleLesson{
			NumLessonName: Models.NumDenLesson{
				LessonName: numLesson.Name,
				Teacher:    numLesson.Teacher.FIO + " " + numLesson.TeacherTwo.FIO,
			},
			DenLessonName: Models.NumDenLesson{
				LessonName: denLesson.Name,
				Teacher:    denLesson.Teacher.FIO,
			},
		}
	} else if denLesson.DoubleTeacher {
		day = Models.ScheduleLesson{
			NumLessonName: Models.NumDenLesson{
				LessonName: numLesson.Name,
				Teacher:    numLesson.Teacher.FIO,
			},
			DenLessonName: Models.NumDenLesson{
				LessonName: denLesson.Name,
				Teacher:    denLesson.Teacher.FIO + " " + denLesson.TeacherTwo.FIO,
			},
		}
	} else {
		day = Models.ScheduleLesson{
			NumLessonName: Models.NumDenLesson{
				LessonName: numLesson.Name,
				Teacher:    numLesson.Teacher.FIO,
			},
			DenLessonName: Models.NumDenLesson{
				LessonName: denLesson.Name,
				Teacher:    denLesson.Teacher.FIO,
			},
		}
	}

	return day
}
