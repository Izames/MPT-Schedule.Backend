package WorkWithFiles

import (
	"MPT-Schedule/Models"
	"fmt"
	"github.com/xuri/excelize/v2"
)

func CreateGroupSC(SheetName string, file *excelize.File, group Models.ScheduleModel, semester int) *excelize.File {
	AllBorder, _ := file.NewStyle(&excelize.Style{
		Border: []excelize.Border{
			{
				Type:  "top",
				Color: "#000000",
				Style: 1,
			},
			{
				Type:  "bottom",
				Color: "#000000",
				Style: 1,
			},
			{
				Type:  "left",
				Color: "#000000",
				Style: 1,
			},
			{
				Type:  "right",
				Color: "#000000",
				Style: 1,
			},
		},
		Alignment: &excelize.Alignment{
			Horizontal: "center",
			Vertical:   "center",
		},
	})
	var Week = []string{"П\nО\nН\nЕ\nД\nЕ\nЛ\nЬ\nН\nИ\nК", "В\nТ\nО\nР\nН\nИ\nК", "С\nР\nЕ\nД\nА", "Ч\nЕ\nТ\nВ\nЕ\nР\nГ", "П\nЯ\nТ\nН\nИ\nЦ\nА", "С\nУ\nБ\nБ\nО\nТ\nА"}
	for i := 0; i < 6; i++ {
		file.MergeCell(SheetName, fmt.Sprintf("%s%d", Models.Columns[1], 2+i*15), fmt.Sprintf("%s%d", Models.Columns[2], 2+i*15))
		file.SetCellValue(SheetName, fmt.Sprintf("%s%d", Models.Columns[1], 2+i*15), group.Semester1.Week[i].Building)
		file.MergeCell(SheetName, fmt.Sprintf("%s%d", Models.Columns[0], 2+i*15), fmt.Sprintf("%s%d", Models.Columns[0], 16+i*15))
		file.SetCellValue(SheetName, fmt.Sprintf("%s%d", Models.Columns[0], 2+i*15), Week[i])
		for j := 0; j < 7; j++ {
			file.MergeCell(SheetName, fmt.Sprintf("%s%d", Models.Columns[1], 3+i*15+j*2), fmt.Sprintf("%s%d", Models.Columns[1], 4+i*15+j*2))
			file.SetCellValue(SheetName, fmt.Sprintf("%s%d", Models.Columns[1], 2+i*15+j*2+1), j)
		}
	}
	for i := 1; i < 92; i++ {
		file.SetRowHeight(SheetName, i, 25)
	}
	if semester == 1 {
		for i, lesson := range group.Semester1.Week {
			if !lesson.StudyingDay {
				for j := 0; j < 7; j++ {
					file.MergeCell(SheetName, fmt.Sprintf("%s%d", Models.Columns[2], 3+i*15+j*2), fmt.Sprintf("%s%d", Models.Columns[2], 4+i*15+j*2))
					file.SetCellValue(SheetName, fmt.Sprintf("%s%d", Models.Columns[2], 3+i*15+j*2), lesson.Building)
				}
			} else {
				for j, pair := range lesson.Lessons {
					if SheetName == "Числитель" {
						file.SetCellValue(SheetName, fmt.Sprintf("%s%d", Models.Columns[2], 3+i*15+j*2), pair.NumLessonName.LessonName)
						file.SetCellValue(SheetName, fmt.Sprintf("%s%d", Models.Columns[2], 4+i*15+j*2), pair.NumLessonName.Teacher)
					} else if SheetName == "Знаменатель" {
						file.SetCellValue(SheetName, fmt.Sprintf("%s%d", Models.Columns[2], 3+i*15+j*2), pair.DenLessonName.LessonName)
						file.SetCellValue(SheetName, fmt.Sprintf("%s%d", Models.Columns[2], 4+i*15+j*2), pair.DenLessonName.Teacher)
					}
				}
			}
		}
	} else {
		for i, lesson := range group.Semester2.Week {
			if !lesson.StudyingDay {
				for j := 0; j < 7; j++ {
					file.MergeCell(SheetName, fmt.Sprintf("%s%d", Models.Columns[2], 3+i*15+j*2), fmt.Sprintf("%s%d", Models.Columns[2], 4+i*15+j*2))
					file.SetCellValue(SheetName, fmt.Sprintf("%s%d", Models.Columns[2], 3+i*15+j*2), lesson.Building)
				}
			} else {
				for j, pair := range lesson.Lessons {
					if SheetName == "Числитель" {
						file.SetCellValue(SheetName, fmt.Sprintf("%s%d", Models.Columns[2], 3+i*15+j*2), pair.NumLessonName.LessonName)
						file.SetCellValue(SheetName, fmt.Sprintf("%s%d", Models.Columns[2], 4+i*15+j*2), pair.NumLessonName.Teacher)
					} else if SheetName == "Знаменатель" {
						file.SetCellValue(SheetName, fmt.Sprintf("%s%d", Models.Columns[2], 3+i*15+j*2), pair.DenLessonName.LessonName)
						file.SetCellValue(SheetName, fmt.Sprintf("%s%d", Models.Columns[2], 4+i*15+j*2), pair.DenLessonName.Teacher)
					}
				}
			}
		}
	}
	file.SetCellValue(SheetName, fmt.Sprintf("%s%d", Models.Columns[0], 1), "день")
	file.SetCellValue(SheetName, fmt.Sprintf("%s%d", Models.Columns[1], 1), "пара")
	file.SetColWidth(SheetName, "C", "C", 80)
	file.SetCellStyle(SheetName, fmt.Sprintf("%s%d", Models.Columns[0], 1), fmt.Sprintf("%s%d", Models.Columns[2], 91), AllBorder)
	return file
}
