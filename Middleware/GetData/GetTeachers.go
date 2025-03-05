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
	var week []Models.LessonDayModel
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
			lessonsForDay := Models.LessonDayModel{
				Lessons: []bool{row[1] == "+", row[2] == "+", row[3] == "+", row[4] == "+", row[5] == "+"},
			}
			freeDay := Models.LessonDayModel{
				Lessons: []bool{false, false, false, false, false},
			}
			for j := 0; j < 6; j++ {
				week = append(week, utils.CheckPrepodDay(freeDay, lessonsForDay, row[6] == "+"))
			}
			for j, _ := range week {
				week[j].Day = j + 1
			}
			Models.TeachersS1 = append(Models.TeachersS1, Models.TeacherModel{
				FIO:          row[0],
				Week:         week,
				Window:       row[12] == "+",
				LessonsInDay: lessons,
				Builds:       buildsForTeacher,
			})
			Models.TeachersS2 = append(Models.TeachersS2, Models.TeacherModel{
				FIO:          row[0],
				Week:         week,
				Window:       row[12] == "+",
				LessonsInDay: lessons,
				Builds:       buildsForTeacher,
			})
			week = []Models.LessonDayModel{}
		}
	}
}
