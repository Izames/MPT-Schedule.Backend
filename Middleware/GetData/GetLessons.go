package GetData

import (
	"MPT-Schedule/Middleware/utils"
	"MPT-Schedule/Models"
	"fmt"
	"github.com/xuri/excelize/v2"
	"regexp"
	"strconv"
	"strings"
)

func GetLessons(file *excelize.File, fileName string, rData *Models.RequestData) {
	//листы для считывания информации
	sheets := file.GetSheetList()
	re := regexp.MustCompile(`^ПП`)

	//перебор листов
	for _, sheet := range sheets {

		//список групп на этом листе
		var groupsId []int

		//сокращенный массив групп состоящий из их названий
		var groupsForCheck []string
		for _, group := range rData.Groups {
			groupsForCheck = append(groupsForCheck, group.Name)
		}

		//пропуск листа если это лист практики
		if re.MatchString(sheet) {
			continue
		}

		//получение строк с листа
		rows, _ := file.GetRows(sheet)

		//переключение с режима обработки групп на режим обработки предметов
		switchToLessons := false

		//столбец с которого начинаются учителя и группы
		prepodsStart := 0

		//столбец количества пар в неделю
		PerWeekValue1S := 0
		PerWeekValue2S := 0

		//пропуск листа
		stop := false

		//перебор строк для обработки
		for _, row := range rows {
			if row == nil {
				continue
			}
			if stop {
				continue
			}
			if row[0] == "ИТОГО" {
				break
			}

			//режим обработки предметов
			if switchToLessons {
				re = regexp.MustCompile(`^[А-ЯЁ]\.[А-ЯЁ]\.\s*[А-ЯЁ][а-яё]+\s+[А-ЯЁ]\.[А-ЯЁ]\.\s*[А-ЯЁ][а-яё]+$`)
				next := false
				if len(row) < prepodsStart+len(groupsId) {
					continue
				}
				if row[PerWeekValue1S] == row[PerWeekValue2S] && row[PerWeekValue2S] == "" {
					continue
				}
				for i := prepodsStart; i < len(groupsId)-1+prepodsStart; i++ {
					if row[i] == "" {
						next = true
						break
					}
				}
				if next {
					continue
				}
				for i := prepodsStart; i < len(groupsId)+prepodsStart; i++ {
					teacher1S1 := &Models.TeacherModel{}
					teacher2S1 := &Models.TeacherModel{}
					teacher1S2 := &Models.TeacherModel{}
					teacher2S2 := &Models.TeacherModel{}
					var doubleTeacher bool
					//проверка на английский
					if re.MatchString(row[i]) {
						doubleTeacher = true
						match1, match2, _ := strings.Cut(row[i], "\n")
						teacher1S1 = utils.FindTeacherByName(match1, 1, rData)
						teacher2S1 = utils.FindTeacherByName(match2, 1, rData)
						teacher1S2 = utils.FindTeacherByName(match1, 2, rData)
						teacher2S2 = utils.FindTeacherByName(match2, 2, rData)
						if teacher1S1 == nil || teacher2S2 == nil {
							stop = true
							rData.FilesErrors = append(rData.FilesErrors, fmt.Sprintf("Один из преподавателей английского не найден в системе %s на листе %s", row[i], sheet))
							break
						}

					} else {
						doubleTeacher = false
						match1 := row[i]
						teacher1S1 = utils.FindTeacherByName(match1, 1, rData)
						teacher1S2 = utils.FindTeacherByName(match1, 2, rData)
						if teacher1S1 == nil {
							stop = true
							rData.FilesErrors = append(rData.FilesErrors, fmt.Sprintf("Неопознанный преподаватель %s на листе %s", row[i], sheet))
							break
						}
					}

					Value1S, _ := strconv.Atoi(row[PerWeekValue1S])
					Value2S, _ := strconv.Atoi(row[PerWeekValue2S])
					if Value1S > 0 {
						lessonS1 := Models.LessonModel{
							Name:          row[1],
							PerWeek:       float32(Value1S) / 2,
							Teacher:       *teacher1S1,
							TeacherTwo:    *teacher2S1,
							DoubleTeacher: doubleTeacher,
						}
						rData.Groups[groupsId[i-prepodsStart]].LessonsS1 = append(rData.Groups[groupsId[i-prepodsStart]].LessonsS1, lessonS1)
						rData.Groups[groupsId[i-prepodsStart]].ListName = sheet
						rData.Groups[groupsId[i-prepodsStart]].FileName = fileName
					}

					if Value2S > 0 {
						lessonS2 := Models.LessonModel{
							Name:          row[1],
							PerWeek:       float32(Value2S) / 2,
							Teacher:       *teacher1S2,
							TeacherTwo:    *teacher2S2,
							DoubleTeacher: doubleTeacher,
						}
						rData.Groups[groupsId[i-prepodsStart]].LessonsS2 = append(rData.Groups[groupsId[i-prepodsStart]].LessonsS2, lessonS2)
						rData.Groups[groupsId[i-prepodsStart]].ListName = sheet
						rData.Groups[groupsId[i-prepodsStart]].FileName = fileName

					}
				}
			}

			//режим обработки групп и столбцов
			if !switchToLessons {
				secondSemester := false
				check := false
				for i := 0; i < len(row); i++ {
					if check {
						groupId := utils.FindGroupByName(row[i], rData)
						if groupId != -1 {
							groupsId = append(groupsId, groupId)
							switchToLessons = true
						} else {
							rData.FilesErrors = append(rData.FilesErrors, fmt.Sprintf("Неопознанная группа %s на листе %s", row[i], sheet))
							stop = true
						}
					}
					if strings.ToLower(row[i]) == "экз" {
						if secondSemester {
							check = true
							i = i + 1
							prepodsStart = i + 1
						}
						secondSemester = true
					}
					if strings.ToLower(row[i]) == "час в нед" {
						if PerWeekValue1S == 0 {
							PerWeekValue1S = i
						} else {
							PerWeekValue2S = i
						}
					}
				}
			}

		}

	}
	println("")
}
