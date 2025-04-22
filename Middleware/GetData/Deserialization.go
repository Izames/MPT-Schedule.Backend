package GetData

import (
	"MPT-Schedule/Models"
	"github.com/xuri/excelize/v2"
	"log"
	"mime/multipart"
)

func Deserialization(form *multipart.Form, rData *Models.RequestData) error {
	rData.Director = form.Value["director"][0]
	rData.Years = form.Value["years"][0]
	rData.ValidityTerm = form.Value["validity_term"][0]
	rData.EndDate = form.Value["end_date"][0]
	rData.DeputyDirectorUR = form.Value["deputy_director_ur"][0]
	rData.MethodologicalDepartment = form.Value["methodological_department"][0]
	rData.CurrentYear = form.Value["current_year"][0]

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
		rData.Extracts = append(rData.Extracts, excelFile)
		file.Close()
		excelFile.Close()
	}

	TeacherHeader := form.File["teacher"][0]
	TeacherFile, err := TeacherHeader.Open()
	if err != nil {
		return err
	}
	rData.Teacher, err = excelize.OpenReader(TeacherFile)
	if err != nil {
		return err
	}
	TeacherFile.Close()
	rData.Teacher.Close()

	GroupHeader := form.File["group"][0]
	GroupFile, err := GroupHeader.Open()
	if err != nil {
		return err
	}
	rData.Group, err = excelize.OpenReader(GroupFile)
	if err != nil {
		return err
	}
	GroupFile.Close()
	rData.Group.Close()

	return nil
}
