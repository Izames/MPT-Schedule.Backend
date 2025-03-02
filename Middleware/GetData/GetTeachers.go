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
			lessons, _ := strconv.Atoi(row[12])
			var buildsForTeacher []string
			for j := 0; j < buildRows; j++ {
				if row[14+j] == "+" {
					buildsForTeacher = append(buildsForTeacher, buildingsName[j])
				}
			}
			lessonsForDay := Models.LessonDayModel{
				FirstLesson:  row[1] == "+",
				SecondLesson: row[2] == "+",
				ThirdLesson:  row[3] == "+",
				FourthLesson: row[4] == "+",
				FifthLesson:  row[5] == "+",
			}
			freeDay := Models.LessonDayModel{
				FirstLesson:  false,
				SecondLesson: false,
				ThirdLesson:  false,
				FourthLesson: false,
				FifthLesson:  false,
			}
			Models.Teachers = append(Models.Teachers, Models.TeacherModel{
				FIO:          row[0],
				Monday:       utils.CheckPrepodDay(freeDay, lessonsForDay, row[6] == "+"),
				Tuesday:      utils.CheckPrepodDay(freeDay, lessonsForDay, row[7] == "+"),
				Wednesday:    utils.CheckPrepodDay(freeDay, lessonsForDay, row[8] == "+"),
				Thursday:     utils.CheckPrepodDay(freeDay, lessonsForDay, row[9] == "+"),
				Friday:       utils.CheckPrepodDay(freeDay, lessonsForDay, row[10] == "+"),
				Saturday:     utils.CheckPrepodDay(freeDay, lessonsForDay, row[11] == "+"),
				Window:       row[12] == "+",
				LessonsInDay: lessons,
				Builds:       buildsForTeacher,
			})
		}
	}
}
