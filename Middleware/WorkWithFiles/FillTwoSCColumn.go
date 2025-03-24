package WorkWithFiles

import (
	"MPT-Schedule/Models"
	"fmt"
	"github.com/xuri/excelize/v2"
)

func FillTwoSCColumn(SheetName string, file *excelize.File, schedule []Models.ScheduleModel, semester int) *excelize.File {
	var Week1 []Models.ScheduleDay
	var Week2 []Models.ScheduleDay
	var Weeks [][]Models.ScheduleDay
	if semester == 1 {
		Week1 = schedule[0].Semester1.Week
		Week2 = schedule[1].Semester1.Week
	} else {
		Week1 = schedule[0].Semester2.Week
		Week2 = schedule[1].Semester2.Week
	}
	Weeks = append(Weeks, Week1)
	Weeks = append(Weeks, Week2)

	whiteStyle, _ := file.NewStyle(&excelize.Style{
		Fill: excelize.Fill{
			Type:    "pattern",
			Color:   []string{"#FFFFFF"}, // Белый цвет
			Pattern: 1,
		},
	})
	RightLine, _ := file.NewStyle(&excelize.Style{
		Border: []excelize.Border{
			{
				Type:  "right",
				Color: "#000000",
				Style: 1,
			},
		},
		Fill: excelize.Fill{
			Type:    "pattern",
			Color:   []string{"#FFFFFF"}, // Белый цвет
			Pattern: 1,
		},
	})
	RightAndDownLine, _ := file.NewStyle(&excelize.Style{
		Border: []excelize.Border{
			{
				Type:  "right",
				Color: "#000000",
				Style: 1,
			},
			{
				Type:  "bottom",
				Color: "#000000",
				Style: 1,
			},
		},
		Fill: excelize.Fill{
			Type:    "pattern",
			Color:   []string{"#FFFFFF"}, // Белый цвет
			Pattern: 1,
		},
		Alignment: &excelize.Alignment{
			Horizontal: "center",
		},
	})
	BottomLine, _ := file.NewStyle(&excelize.Style{
		Border: []excelize.Border{
			{
				Type:  "bottom",
				Color: "#000000",
				Style: 1,
			},
		},
		Fill: excelize.Fill{
			Type:    "pattern",
			Color:   []string{"#FFFFFF"}, // Белый цвет
			Pattern: 1,
		},
	})
	FullLineStyle, _ := file.NewStyle(&excelize.Style{
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
		Fill: excelize.Fill{
			Type:    "pattern",
			Color:   []string{"#FFFFFF"}, // Белый цвет
			Pattern: 1,
		},
	})
	boldStyle, _ := file.NewStyle(&excelize.Style{
		Fill: excelize.Fill{
			Type:    "pattern",
			Color:   []string{"#FFFFFF"}, // Белый цвет
			Pattern: 1,
		},
		Font: &excelize.Font{
			Bold: true, // Жирный шрифт
		},
		Alignment: &excelize.Alignment{
			Horizontal: "center",
		},
	})
	WeekDayStyle, _ := file.NewStyle(&excelize.Style{
		Fill: excelize.Fill{
			Type:    "pattern",
			Color:   []string{"#FFFFFF"}, // Белый цвет
			Pattern: 1,
		},
		Font: &excelize.Font{
			Bold: true, // Жирный шрифт
		},
		Alignment: &excelize.Alignment{
			Horizontal: "center", // Горизонтальное выравнивание по центру
			Vertical:   "center",
		},
		Border: []excelize.Border{
			{Type: "left", Color: "#000000", Style: 1},   // Левая граница
			{Type: "right", Color: "#000000", Style: 1},  // Правая граница
			{Type: "top", Color: "#000000", Style: 1},    // Верхняя граница
			{Type: "bottom", Color: "#000000", Style: 1}, // Нижняя граница
		},
	})
	centerTextStyle, _ := file.NewStyle(&excelize.Style{
		Fill: excelize.Fill{
			Type:    "pattern",
			Color:   []string{"#FFFFFF"}, // Белый цвет
			Pattern: 1,
		},
		Alignment: &excelize.Alignment{
			Horizontal: "center",
		},
	})
	weekDays := []string{"П\nО\nН\nЕ\nД\nЕ\nЛ\nЬ\nН\nИ\nК", "В\nТ\nО\nР\nН\nИ\nК", "С\nР\nЕ\nД\nА", "Ч\nЕ\nТ\nВ\nЕ\nР\nГ", "П\nЯ\nТ\nН\nИ\nЦ\nА", "С\nУ\nБ\nБ\nО\nТ\nА"}
	groups := ""
	for _, sc := range schedule {
		if groups == "" {
			groups += sc.Group
		} else {
			groups += ", " + sc.Group
		}
	}
	//создание шапки
	file.SetCellStyle(SheetName, fmt.Sprintf("%s%d", Models.Columns[0], 1), fmt.Sprintf("%s%d", Models.Columns[8], 96), whiteStyle)
	file.SetCellValue(SheetName, fmt.Sprintf("%s%d", Models.Columns[6], 1), "УТВЕРЖДАЮ")
	file.SetCellValue(SheetName, fmt.Sprintf("%s%d", Models.Columns[6], 2), "Директор Московского приборостроительного техникума")
	file.SetCellValue(SheetName, fmt.Sprintf("%s%d", Models.Columns[6], 3), fmt.Sprintf("____________________________ %s", Models.Director))
	file.SetCellValue(SheetName, fmt.Sprintf("%s%d", Models.Columns[6], 4), fmt.Sprintf("\"________\" _______________________%s.г", Models.CurrentYear))
	file.MergeCell(SheetName, fmt.Sprintf("%s%d", Models.Columns[0], 6), fmt.Sprintf("%s%d", Models.Columns[8], 6))
	file.SetCellStyle(SheetName, fmt.Sprintf("%s%d", Models.Columns[0], 6), fmt.Sprintf("%s%d", Models.Columns[8], 6), boldStyle)
	file.SetCellValue(SheetName, fmt.Sprintf("%s%d", Models.Columns[6], 6), fmt.Sprintf("Расписание учебных занятий на %d семестр %s учебного года", semester, Models.Years))
	file.MergeCell(SheetName, fmt.Sprintf("%s%d", Models.Columns[0], 7), fmt.Sprintf("%s%d", Models.Columns[8], 7))
	file.SetCellStyle(SheetName, fmt.Sprintf("%s%d", Models.Columns[0], 7), fmt.Sprintf("%s%d", Models.Columns[8], 7), boldStyle)
	file.SetCellValue(SheetName, fmt.Sprintf("%s%d", Models.Columns[6], 7), fmt.Sprintf("групп %s", groups))
	file.SetCellStyle(SheetName, fmt.Sprintf("%s%d", Models.Columns[0], 8), fmt.Sprintf("%s%d", Models.Columns[8], 8), centerTextStyle)
	file.MergeCell(SheetName, fmt.Sprintf("%s%d", Models.Columns[0], 8), fmt.Sprintf("%s%d", Models.Columns[8], 8))
	file.SetCellValue(SheetName, fmt.Sprintf("%s%d", Models.Columns[6], 8), fmt.Sprintf("действует с %s по %s", Models.ValidityTerm, Models.EndDate))
	//создание footer
	file.SetCellValue(SheetName, fmt.Sprintf("%s%d", Models.Columns[1], 95), fmt.Sprintf("Заместитель директора по УР                                           %s", Models.DeputyDirectorUR))
	file.SetCellStyle(SheetName, fmt.Sprintf("%s%d", Models.Columns[5], 95), fmt.Sprintf("%s%d", Models.Columns[5], 95), centerTextStyle)
	file.SetCellValue(SheetName, fmt.Sprintf("%s%d", Models.Columns[5], 95), fmt.Sprintf("Начальник учебно-методического отдела                     %s", Models.MethodologicalDepartment))
	file.MergeCell(SheetName, fmt.Sprintf("%s%d", Models.Columns[5], 95), fmt.Sprintf("%s%d", Models.Columns[8], 95))
	//создание шаблона

	file.SetCellStyle(SheetName, fmt.Sprintf("%s%d", Models.Columns[0], 9), fmt.Sprintf("%s%d", Models.Columns[5], 9), FullLineStyle)
	file.SetCellValue(SheetName, fmt.Sprintf("%s%d", Models.Columns[2], 9), schedule[0].Group)
	file.SetCellValue(SheetName, fmt.Sprintf("%s%d", Models.Columns[5], 9), schedule[1].Group)
	for j := 0; j < 2; j++ {
		file.SetCellValue(SheetName, fmt.Sprintf("%s%d", Models.Columns[0+j*3], 9), "День")
		file.SetCellValue(SheetName, fmt.Sprintf("%s%d", Models.Columns[1+j*3], 9), "пара")

		for i := 0; i < 6; i++ {
			file.MergeCell(SheetName, fmt.Sprintf("%s%d", Models.Columns[0+j*3], 10+i*14), fmt.Sprintf("%s%d", Models.Columns[0+j*3], 23+i*14))
		}
		for i := 0; i < 6; i++ {
			file.SetCellValue(SheetName, fmt.Sprintf("%s%d", Models.Columns[0+j*3], 10+i*14), weekDays[i])
		}
		for i := 0; i < 6; i++ {
			file.SetCellStyle(SheetName, fmt.Sprintf("%s%d", Models.Columns[0+j*3], 10+i*14), fmt.Sprintf("%s%d", Models.Columns[0+j*3], 23+i*14), WeekDayStyle)
		}
		for i := 0; i < 6; i++ {
			for c := 0; c < 7; c++ {
				file.SetCellStyle(SheetName, fmt.Sprintf("%s%d", Models.Columns[1+j*3], 10+i*14+c*2), fmt.Sprintf("%s%d", Models.Columns[1+j*3], 10+i*14+c*2), RightLine)
				file.SetCellStyle(SheetName, fmt.Sprintf("%s%d", Models.Columns[1+j*3], 11+i*14+c*2), fmt.Sprintf("%s%d", Models.Columns[1+j*3], 11+i*14+c*2), RightAndDownLine)
				file.SetCellValue(SheetName, fmt.Sprintf("%s%d", Models.Columns[1+j*3], 11+i*14+c*2), c)

				file.SetCellStyle(SheetName, fmt.Sprintf("%s%d", Models.Columns[2+j*3], 11+i*14+c*2), fmt.Sprintf("%s%d", Models.Columns[2+j*3], 11+i*14+c*2), BottomLine)
			}
		}
	}
	for i := 0; i < 3; i++ {
		file.SetColWidth(SheetName, Models.Columns[0+i*3], Models.Columns[1+i*3], 5)
		file.SetColWidth(SheetName, Models.Columns[2+i*3], Models.Columns[2+i*3], 100)
	}
	//заполнение шаблона

	for j := 0; j < 2; j++ {
		for i := 0; i < 6; i++ {
			if !Weeks[j][i].StudyingDay {
				for c := 0; c < 7; c++ {
					file.MergeCell(SheetName, fmt.Sprintf("%s%d", Models.Columns[2+j*3], 10+c*2+i*14), fmt.Sprintf("%s%d", Models.Columns[2+j*3], 11+c*2+i*14))
					file.SetCellValue(SheetName, fmt.Sprintf("%s%d", Models.Columns[2+j*3], 10+c*2+i*14), Weeks[j][i].Building)
				}
			} else {
				file.MergeCell(SheetName, fmt.Sprintf("%s%d", Models.Columns[2+j*3], 10+i*14), fmt.Sprintf("%s%d", Models.Columns[2+j*3], 11+i*14))
				file.SetCellValue(SheetName, fmt.Sprintf("%s%d", Models.Columns[2+j*3], 10+i*14), Weeks[j][i].Building)
				for c, lesson := range Weeks[j][i].Lessons {
					if lesson.NumLessonName == lesson.DenLessonName {
						file.SetCellValue(SheetName, fmt.Sprintf("%s%d", Models.Columns[2+j*3], 12+i*14+c*2), lesson.NumLessonName.LessonName)
						file.SetCellValue(SheetName, fmt.Sprintf("%s%d", Models.Columns[2+j*3], 13+i*14+c*2), lesson.NumLessonName.Teacher)
					} else {
						file.SetCellValue(SheetName, fmt.Sprintf("%s%d", Models.Columns[2+j*3], 12+i*14+c*2), lesson.NumLessonName.LessonName+" "+lesson.NumLessonName.Teacher)
						file.SetCellValue(SheetName, fmt.Sprintf("%s%d", Models.Columns[2+j*3], 13+i*14+c*2), lesson.DenLessonName.LessonName+" "+lesson.NumLessonName.Teacher)
					}

				}
			}
		}
	}
	return file
}
