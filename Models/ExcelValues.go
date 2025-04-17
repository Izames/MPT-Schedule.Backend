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
var CurrentYear string

var ValidityTerm string

var EndDate string

var DeputyDirectorUR string

var MethodologicalDepartment string

var Extracts []*excelize.File

var Teacher *excelize.File

var Group *excelize.File

var FilesErrors []string

func Clean(files []string) {
	Director = ""
	Years = ""
	CurrentYear = ""
	ValidityTerm = ""
	EndDate = ""
	DeputyDirectorUR = ""
	MethodologicalDepartment = ""
	Extracts = nil
	TeachersS1 = []TeacherModel{}
	TeachersS2 = []TeacherModel{}
	Groups = []GroupModel{}
	Schedules = []ScheduleModel{}
	Builds = []string{}
	Teacher = nil
	Group = nil
	FilesErrors = []string{}
	for _, file := range files {
		os.Remove(file)
	}
}
