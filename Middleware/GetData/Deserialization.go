package GetData

import (
	"MPT-Schedule/Models"
	"github.com/xuri/excelize/v2"
	"log"
	"mime/multipart"
)

func Deserialization(form *multipart.Form) error {
	Models.Director = form.Value["director"][0]
	Models.Years = form.Value["years"][0]
	Models.ValidityTerm = form.Value["validity_term"][0]
	Models.EndDate = form.Value["end_date"][0]
	Models.DeputyDirectorUR = form.Value["deputy_director_ur"][0]
	Models.MethodologicalDepartment = form.Value["methodological_department"][0]
	Models.CurrentYear = form.Value["current_year"][0]

	files := form.File["extracts"]
	for _, fileHeader := range files {
		file, err := fileHeader.Open()
		if err != nil {
			log.Println("Произошла ошибка при парсинге файлов: ", err)
			return err
		}
		excelFile, err := excelize.OpenReader(file)
		if err != nil {
			log.Println("Произошла ошибка при парсинге файлов: ", err)
			return err
		}
		Models.Extracts = append(Models.Extracts, excelFile)
		file.Close()
		excelFile.Close()
	}

	TeacherHeader := form.File["teacher"][0]
	TeacherFile, err := TeacherHeader.Open()
	if err != nil {
		return err
	}
	Models.Teacher, err = excelize.OpenReader(TeacherFile)
	if err != nil {
		return err
	}
	TeacherFile.Close()
	Models.Teacher.Close()

	GroupHeader := form.File["group"][0]
	GroupFile, err := GroupHeader.Open()
	if err != nil {
		return err
	}
	Models.Group, err = excelize.OpenReader(GroupFile)
	if err != nil {
		return err
	}
	GroupFile.Close()
	Models.Group.Close()

	return nil
}
