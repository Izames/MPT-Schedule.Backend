package GetData

import (
	"MPT-Schedule/Middleware/utils"
	"MPT-Schedule/Models"
	"fmt"
	"strconv"
	"strings"
)

func GetTeachers() {
	sheets := Models.Teacher.GetSheetList()

	line, _ := Models.Teacher.GetRows(sheets[0])
	var buildingsName []string
	for i := 1; i < len(line[0]); i++ {
		if i < 14 {
			continue
		}
		buildingsName = append(buildingsName, strings.ToLower(line[0][i]))
	}
	if len(Models.Builds) == 0 {
		Models.Builds = buildingsName
	}
	buildRows := len(line[0]) - 14

	for _, sheet := range sheets {
		rows, _ := Models.Teacher.GetRows(sheet)
		for i, row := range rows {
			var stop = false
			if i == 0 {
				continue
			}
			if len(row) != len(rows[0]) {
				Models.FilesErrors = append(Models.FilesErrors, fmt.Sprintf("Неправильно заполнены ограничения преподавателя %s в листе %s", row[0], sheet))
				continue
			}
			for _, str := range row {
				if str == "" {
					Models.FilesErrors = append(Models.FilesErrors, fmt.Sprintf("Неправильно заполнены ограничения преподавателя %s в листе %s", row[0], sheet))
					break
				}
			}
			if stop {
				continue
			}
			if row[13] != "1" && row[13] != "2" && row[13] != "3" && row[13] != "4" && row[13] != "5" {
				Models.FilesErrors = append(Models.FilesErrors, fmt.Sprintf("Неправильно заполнено ограничение пар преподавателя %s в листе %s. Количество должно быть от 1 до 5", row[0], sheet))
				continue
			}
			lessons, _ := strconv.Atoi(row[13])
			var buildsForTeacher []string
			for j := 0; j < buildRows; j++ {
				if row[14+j] == "+" {
					buildsForTeacher = append(buildsForTeacher, buildingsName[j])
				}
			}

			// Создаем новые срезы для каждого преподавателя
			week1 := make([]Models.LessonDayModel, 6)
			week2 := make([]Models.LessonDayModel, 6)

			// Инициализация Lessons для каждого дня
			for j := range week1 {
				week1[j].Lessons = make([]bool, 5)
				week2[j].Lessons = make([]bool, 5)
			}

			// Заполнение week1 и week2
			freeDay := Models.LessonDayModel{
				Lessons: []bool{false, false, false, false, false},
			}
			week1[0] = utils.CheckPrepodDay(freeDay, Models.LessonDayModel{Lessons: []bool{row[1] == "+", row[2] == "+", row[3] == "+", row[4] == "+", row[5] == "+"}}, row[6] == "+")
			week1[1] = utils.CheckPrepodDay(freeDay, Models.LessonDayModel{Lessons: []bool{row[1] == "+", row[2] == "+", row[3] == "+", row[4] == "+", row[5] == "+"}}, row[7] == "+")
			week1[2] = utils.CheckPrepodDay(freeDay, Models.LessonDayModel{Lessons: []bool{row[1] == "+", row[2] == "+", row[3] == "+", row[4] == "+", row[5] == "+"}}, row[8] == "+")
			week1[3] = utils.CheckPrepodDay(freeDay, Models.LessonDayModel{Lessons: []bool{row[1] == "+", row[2] == "+", row[3] == "+", row[4] == "+", row[5] == "+"}}, row[9] == "+")
			week1[4] = utils.CheckPrepodDay(freeDay, Models.LessonDayModel{Lessons: []bool{row[1] == "+", row[2] == "+", row[3] == "+", row[4] == "+", row[5] == "+"}}, row[10] == "+")
			week1[5] = utils.CheckPrepodDay(freeDay, Models.LessonDayModel{Lessons: []bool{row[1] == "+", row[2] == "+", row[3] == "+", row[4] == "+", row[5] == "+"}}, row[11] == "+")

			week2[0] = utils.CheckPrepodDay(freeDay, Models.LessonDayModel{Lessons: []bool{row[1] == "+", row[2] == "+", row[3] == "+", row[4] == "+", row[5] == "+"}}, row[6] == "+")
			week2[1] = utils.CheckPrepodDay(freeDay, Models.LessonDayModel{Lessons: []bool{row[1] == "+", row[2] == "+", row[3] == "+", row[4] == "+", row[5] == "+"}}, row[7] == "+")
			week2[2] = utils.CheckPrepodDay(freeDay, Models.LessonDayModel{Lessons: []bool{row[1] == "+", row[2] == "+", row[3] == "+", row[4] == "+", row[5] == "+"}}, row[8] == "+")
			week2[3] = utils.CheckPrepodDay(freeDay, Models.LessonDayModel{Lessons: []bool{row[1] == "+", row[2] == "+", row[3] == "+", row[4] == "+", row[5] == "+"}}, row[9] == "+")
			week2[4] = utils.CheckPrepodDay(freeDay, Models.LessonDayModel{Lessons: []bool{row[1] == "+", row[2] == "+", row[3] == "+", row[4] == "+", row[5] == "+"}}, row[10] == "+")
			week2[5] = utils.CheckPrepodDay(freeDay, Models.LessonDayModel{Lessons: []bool{row[1] == "+", row[2] == "+", row[3] == "+", row[4] == "+", row[5] == "+"}}, row[11] == "+")

			// Устанавливаем дни недели
			for j := range week1 {
				week1[j].Day = j + 1
				week1[j].Build = ""
				week2[j].Day = j + 1
				week2[j].Build = ""
			}

			// Добавляем преподавателя в TeachersS1 и TeachersS2
			Models.TeachersS1 = append(Models.TeachersS1, Models.TeacherModel{
				FIO:          row[0],
				Week:         week1,
				Window:       row[12] == "+",
				LessonsInDay: lessons,
				Builds:       buildsForTeacher,
			})
			Models.TeachersS2 = append(Models.TeachersS2, Models.TeacherModel{
				FIO:          row[0],
				Week:         week2,
				Window:       row[12] == "+",
				LessonsInDay: lessons,
				Builds:       buildsForTeacher,
			})
		}
	}
}
