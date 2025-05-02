package GetData

import (
	"MPT-Schedule/Middleware/utils"
	"MPT-Schedule/Models"
	"fmt"
	"strings"
)

func GetGroups(rData *Models.RequestData) {
	var unSCDays []string
	sheet := rData.Group.GetSheetList()[0]
	cols, _ := rData.Group.GetCols(sheet)
	for i, col := range cols[10] {
		if i < 2 {
			continue
		}
		if col != "" {
			unSCDays = append(unSCDays, strings.ToLower(col))
		}
	}
	rows, _ := rData.Group.GetRows(sheet)
	for i, row := range rows {
		var week []Models.Day
		var stop = false
		if i < 1 {
			continue
		}
		if len(row) < 8 {
			rData.FilesErrors = append(rData.FilesErrors, fmt.Sprintf("Ошибка заполнения группы %s", row[1]))
			continue
		}
		for j := 0; j < 8; j++ {
			if row[j] == "" {
				rData.FilesErrors = append(rData.FilesErrors, fmt.Sprintf("Ошибка заполнения группы %s", row[1]))
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
		rData.Groups = append(rData.Groups, group)

	}
	println("")
}
