package WorkWithFiles

import (
	"MPT-Schedule/Models"
	"github.com/xuri/excelize/v2"
)

func GenerateScheduleFile() ([]*excelize.File, []string) {
	var schedules [][][]Models.ScheduleModel
	schedulesGroupName := ""
	schedulesGroupFile := ""
	id := -1
	fileId := -1
	for _, sc := range Models.Schedules {
		if schedulesGroupFile != sc.File {
			fileId++
			schedulesGroupFile = sc.File
			if schedulesGroupName != sc.List {
				id++
				schedulesGroupName = sc.List
				schedules = append(schedules, [][]Models.ScheduleModel{{sc}})
			}
		} else {
			if schedulesGroupName != sc.List {
				id++
				schedulesGroupName = sc.List
				schedules[fileId] = append(schedules[fileId], []Models.ScheduleModel{sc})
			} else {
				schedules[fileId][id] = append(schedules[fileId][id], sc)
			}
		}

	}
	var files []*excelize.File
	var filesPath []string
	for _, sc := range schedules {
		file1, file2, path1, path2 := CreateSCFile(sc)
		files = append(files, file1)
		files = append(files, file2)
		filesPath = append(filesPath, path1)
		filesPath = append(filesPath, path2)
	}
	return files, filesPath
}
