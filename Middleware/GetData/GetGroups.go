package GetData

import (
	"MPT-Schedule/Middleware/utils"
	"MPT-Schedule/Models"
	"fmt"
	"strings"
)

func GetGroups() {
	var unSCDays []string
	var week []Models.Day
	sheet := Models.Group.GetSheetList()[0]
	cols, _ := Models.Group.GetCols(sheet)
	for i, col := range cols[10] {
		if i < 2 {
			continue
		}
		if col != "" {
			unSCDays = append(unSCDays, strings.ToLower(col))
		}
	}
	rows, _ := Models.Group.GetRows(sheet)
	for i, row := range rows {
		var stop = false
		if i < 1 {
			continue
		}
		if len(row) < 8 {
			Models.FilesErrors = append(Models.FilesErrors, fmt.Sprintf("Ошибка заполнения группы %s", row[1]))
			continue
		}
		for j := 0; j < 8; j++ {
			if row[j] == "" {
				Models.FilesErrors = append(Models.FilesErrors, fmt.Sprintf("Ошибка заполнения группы %s", row[1]))
				stop = true
				break
			}
		}
		if stop {
			continue
		}
		for j := 0; j < 6; j++ {
			week = append(week, Models.Day{DayNum: j + 1, Build: row[2+j], UnSCDay: utils.Contains(strings.ToLower(row[2+j]), unSCDays)})
		}
		group := Models.GroupModel{
			Name: row[1],
			Week: week,
		}
		Models.Groups = append(Models.Groups, group)

	}
	println("")
}
