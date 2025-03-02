package GetData

import (
	"MPT-Schedule/Middleware/utils"
	"MPT-Schedule/Models"
	"fmt"
	"strings"
)

func GetGroups() {
	var unSCDays []string
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
		group := Models.GroupModel{
			Name:      row[1],
			Monday:    Models.Day{Build: row[2], UnSCDay: utils.Contains(strings.ToLower(row[2]), unSCDays)},
			Tuesday:   Models.Day{Build: row[3], UnSCDay: utils.Contains(strings.ToLower(row[3]), unSCDays)},
			Wednesday: Models.Day{Build: row[4], UnSCDay: utils.Contains(strings.ToLower(row[4]), unSCDays)},
			Thursday:  Models.Day{Build: row[5], UnSCDay: utils.Contains(strings.ToLower(row[5]), unSCDays)},
			Friday:    Models.Day{Build: row[6], UnSCDay: utils.Contains(strings.ToLower(row[6]), unSCDays)},
			Saturday:  Models.Day{Build: row[7], UnSCDay: utils.Contains(strings.ToLower(row[7]), unSCDays)},
		}
		Models.Groups = append(Models.Groups, group)

	}
	println("")
}
