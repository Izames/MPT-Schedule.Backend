package WorkWithFiles

import (
	"MPT-Schedule/Models"
	"fmt"
	"github.com/xuri/excelize/v2"
)

func CreateSCFile(schedules [][]Models.ScheduleModel, rData *Models.RequestData) (*excelize.File, *excelize.File, string, string) {
	newFile1S := excelize.NewFile()
	newFile2S := excelize.NewFile()
	filePath1 := ""
	filePath2 := ""
	for _, schedule := range schedules {
		if len(schedule) <= 3 {
			newFile1S.NewSheet(schedule[0].List)
			newFile2S.NewSheet(schedule[0].List)
			newFile1S.DeleteSheet("Sheet1")
			newFile2S.DeleteSheet("Sheet1")
			if len(schedule) == 3 {
				FillThreeSCColumn(schedule[0].List, newFile1S, schedule, 1, rData)
				FillThreeSCColumn(schedule[0].List, newFile2S, schedule, 2, rData)
			}
			if len(schedule) == 2 {
				FillTwoSCColumn(schedule[0].List, newFile1S, schedule, 1, rData)
				FillTwoSCColumn(schedule[0].List, newFile2S, schedule, 2, rData)
			}
			if len(schedule) == 1 {
				FillOneSCColumn(schedule[0].List, newFile1S, schedule, 1, rData)
				FillOneSCColumn(schedule[0].List, newFile2S, schedule, 2, rData)
			}
		} else {
			// Определяем, как делить элементы
			switch len(schedule) % 3 {
			case 0:
				for i := 0; i < len(schedule); i += 3 {
					var sc []Models.ScheduleModel
					for j := 0; j < 3; j++ {
						sc = append(sc, schedule[j+i])
					}
					newFile1S.NewSheet(fmt.Sprintf("%s.%d", schedule[i].List, i))
					newFile2S.NewSheet(fmt.Sprintf("%s.%d", schedule[i].List, i))
					FillThreeSCColumn(fmt.Sprintf("%s.%d", schedule[i].List, i), newFile1S, sc, 1, rData)
					FillThreeSCColumn(fmt.Sprintf("%s.%d", schedule[i].List, i), newFile2S, sc, 2, rData)
				}
			case 1, 2:
				if len(schedule)%2 == 0 {
					for i := 0; i < len(schedule); i += 2 {
						var sc []Models.ScheduleModel
						for j := 0; j < 2; j++ {
							sc = append(sc, schedule[j+i])
						}
						newFile1S.NewSheet(fmt.Sprintf("%s.%d", schedule[i].List, i))
						newFile2S.NewSheet(fmt.Sprintf("%s.%d", schedule[i].List, i))
						FillTwoSCColumn(fmt.Sprintf("%s.%d", schedule[i].List, i), newFile1S, sc, 1, rData)
						FillTwoSCColumn(fmt.Sprintf("%s.%d", schedule[i].List, i), newFile2S, sc, 2, rData)
					}
				} else {
					var sc []Models.ScheduleModel
					for j := 0; j < 3; j++ {
						sc = append(sc, schedule[j])
					}
					newFile1S.NewSheet(fmt.Sprintf("%s.%d", schedule[0].List, 1))
					newFile2S.NewSheet(fmt.Sprintf("%s.%d", schedule[0].List, 1))
					FillThreeSCColumn(fmt.Sprintf("%s.%d", schedule[0].List, 1), newFile1S, sc, 1, rData)
					FillThreeSCColumn(fmt.Sprintf("%s.%d", schedule[0].List, 1), newFile2S, sc, 2, rData)
					for i := 3; i < len(schedule); i += 2 {
						var scs []Models.ScheduleModel
						for j := 0; j < 2; j++ {
							scs = append(scs, schedule[j+i])
						}
						newFile1S.NewSheet(fmt.Sprintf("%s.%d", schedule[i].List, i))
						newFile2S.NewSheet(fmt.Sprintf("%s.%d", schedule[i].List, i))
						FillTwoSCColumn(fmt.Sprintf("%s.%d", schedule[i].List, i), newFile1S, scs, 1, rData)
						FillTwoSCColumn(fmt.Sprintf("%s.%d", schedule[i].List, i), newFile2S, scs, 2, rData)
					}
				}
			}
		}
		filePath1 = schedules[0][0].File + " 1 семестр.xlsx"
		filePath2 = schedules[0][0].File + " 2 семестр.xlsx"
		newFile1S.SaveAs(filePath1)
		newFile2S.SaveAs(filePath2)
	}
	return newFile1S, newFile2S, filePath1, filePath2
}
