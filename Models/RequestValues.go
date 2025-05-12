package Models

import (
	"github.com/xuri/excelize/v2"
)

var Columns = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T",
	"U", "V", "W", "X", "Y", "Z", "AA", "AB", "AC", "AD", "AE", "AF", "AG", "AH", "AI", "AJ", "AK", "AL", "AM", "AN",
	"AO", "AP", "AQ", "AR", "AS"}

type RequestData struct {
	TeachersS1 []TeacherModel
	TeachersS2 []TeacherModel

	Groups []GroupModel

	Schedules []ScheduleModel

	Builds []string

	Director string

	Years       string
	CurrentYear string

	ValidityTerm string

	EndDate string

	DeputyDirectorUR string

	MethodologicalDepartment string

	Extracts []*excelize.File

	Teacher *excelize.File

	Group *excelize.File

	FilesErrors []string

	Failure bool
}
