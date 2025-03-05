package Models

import (
	"github.com/xuri/excelize/v2"
	"os"
)

var TeachersS1 []TeacherModel
var TeachersS2 []TeacherModel

var Groups []GroupModel

var Schedules []ScheduleModel

var Builds []string

var Columns = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T",
	"U", "V", "W", "X", "Y", "Z", "AA", "AB", "AC", "AD", "AE", "AF", "AG", "AH", "AI", "AJ", "AK", "AL", "AM", "AN",
	"AO", "AP", "AQ", "AR", "AS"}

var Director string

var Years string

var ValidityTerm string

var EndDate string

var DeputyDirectorUR string

var DeputyDirectorUMR string

var MethodologicalDepartment string

var Extracts []*excelize.File

var Teacher *excelize.File

var Group *excelize.File

var FilesErrors []string

func Clean() {
	Director = ""
	Years = ""
	ValidityTerm = ""
	EndDate = ""
	DeputyDirectorUR = ""
	DeputyDirectorUMR = ""
	MethodologicalDepartment = ""
	for _, file := range Extracts {
		file.Close()
		os.Remove(file.Path)
	}
	Extracts = nil
	TeachersS1 = []TeacherModel{}
	TeachersS2 = []TeacherModel{}
	Groups = []GroupModel{}
	Schedules = []ScheduleModel{}
	Builds = []string{}
	Teacher = nil
	Group = nil
	FilesErrors = []string{}
}
