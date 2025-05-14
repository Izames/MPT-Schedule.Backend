package utils

import (
	"MPT-Schedule/Models"
	"math"
	"math/rand"
	"time"
)

// RandomLessonSplitter отделяет пары друг от друга: составляет симбиоз и половинчатых пар,
// а также делает из (пара X 2 раза в неделю -> Пара X 1 раз в неделю, Пара X 1 раз в неделю)
func RandomLessonSplitter(lessons []Models.LessonModel, daysForInsert []Models.ScheduleDay, rData *Models.RequestData) []Models.LessonModelND {
	var newLessons []Models.LessonModelND
	var NorDLessons []Models.LessonModel
	var NDLessons []Models.LessonModel

	//отделяем пары от половинчатых
	for _, lesson := range lessons {
		_, frac := math.Modf(float64(lesson.PerWeek))
		if frac != 0 {
			NorDLesson := lesson
			NorDLesson.PerWeek = 0.5
			NorDLessons = append(NorDLessons, NorDLesson)

			NDLesson := lesson
			NDLesson.PerWeek -= 0.5
			NDLessons = append(NDLessons, NDLesson)
		} else {
			NDLessons = append(NDLessons, lesson)
		}
	}

	//обрабатываем целые пары
	for i := range NDLessons {
		lesson := NDLessons[i]
		if lesson.PerWeek == 1 {
			newLessons = append(newLessons, Models.LessonModelND{DenLesson: lesson, NumLesson: lesson, OneND: true})
		} else {
			lessonsCount := int(lesson.PerWeek)
			for j := 0; j < lessonsCount; j++ {
				newLesson := lesson
				newLesson.PerWeek = 1
				newLessons = append(newLessons, Models.LessonModelND{DenLesson: newLesson, NumLesson: newLesson, OneND: true})
			}
		}
	}

	//добавляем половинчатые пары, где одна соединена с другой
	for i := 0; i < len(NorDLessons); i += 2 {
		if i+1 == len(NorDLessons) {
			var per float32
			per = 0
			newLessons = append(newLessons, Models.LessonModelND{DenLesson: NorDLessons[i], NumLesson: Models.LessonModel{Name: "", PerWeek: per}, OneND: false})
		} else {
			newLessons = append(newLessons, Models.LessonModelND{DenLesson: NorDLessons[i], NumLesson: NorDLessons[i+1], OneND: false})
		}
	}

	//перемешать их всех
	for range 10 {
		rand.Seed(time.Now().UnixNano()) // Инициализация генератора случайных чисел
		rand.Shuffle(len(newLessons), func(i, j int) {
			newLessons[i], newLessons[j] = newLessons[j], newLessons[i]
		})
	}
	return SortPriority(daysForInsert, newLessons)
}
