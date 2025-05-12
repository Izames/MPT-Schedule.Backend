package InsertData

import (
	"MPT-Schedule/Middleware/WorkWithFiles"
	"MPT-Schedule/Models"
	"github.com/xuri/excelize/v2"
	"path/filepath"
	"strings"
)

func GroupSchedules(group Models.ScheduleModel, folder string) {
	file1 := excelize.NewFile()
	file2 := excelize.NewFile()
	file1.SetSheetName("Sheet1", "Числитель")
	file2.SetSheetName("Sheet1", "Числитель")
	file1.NewSheet("Знаменатель")
	file2.NewSheet("Знаменатель")
	WorkWithFiles.CreateGroupSC("Числитель", file1, group, 1)
	WorkWithFiles.CreateGroupSC("Знаменатель", file1, group, 1)
	WorkWithFiles.CreateGroupSC("Числитель", file2, group, 2)
	WorkWithFiles.CreateGroupSC("Знаменатель", file2, group, 2)
	replacer := strings.NewReplacer(
		"/", "&",
		"\\", "&",
	)
	filePath1 := filepath.Join(folder, "1 семестр "+replacer.Replace(group.Group)+".xlsx")
	filePath2 := filepath.Join(folder, "2 семестр "+replacer.Replace(group.Group)+".xlsx")
	file1.SaveAs(filePath1)
	file2.SaveAs(filePath2)
}
