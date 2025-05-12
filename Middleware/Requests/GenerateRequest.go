package Requests

import (
	"MPT-Schedule/Middleware/GetData"
	"MPT-Schedule/Middleware/InsertData"
	"MPT-Schedule/Middleware/WorkWithFiles"
	"MPT-Schedule/Middleware/utils"
	"MPT-Schedule/Models"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func GenerateRequest(context *gin.Context) {
	//парсинг данных в мультипарт режиме
	if err := context.Request.ParseMultipartForm(32 << 22); err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
		log.Println("не удалось распарсить данные для генерации расписания мультипарт формы")
		return
	}
	//парсинг запроса в переменную
	form, err := context.MultipartForm()

	if err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
		log.Println("не удалось распарсить данные для генерации расписания")
		return
	}

	attempts := 0
	for {
		requestData := Models.RequestData{}
		extractFiles := form.File["extracts"]
		if len(extractFiles) == 0 {
			context.JSON(400, gin.H{"error": "no files with the name 'extracts' found"})
			log.Println("не удалось найти файлы с именем 'extracts'")
			return
		}
		//Десериализация данных
		if err = GetData.Deserialization(form, &requestData); err != nil {
			context.JSON(400, gin.H{"error": err.Error()})
			log.Println("не удалось распарсить данные для генерации расписания")
			return
		}
		//собрать ограничения учителей
		GetData.GetTeachers(&requestData)
		GetData.GetGroups(&requestData)
		for i, file := range requestData.Extracts {
			GetData.GetLessons(file, extractFiles[i].Filename, &requestData)
		}
		for _, group := range requestData.Groups {
			InsertData.GenerateSchedule(group, &requestData)
		}

		folder, _ := os.MkdirTemp("", "Расписание групп_*")
		for i := range requestData.Schedules {
			InsertData.GroupSchedules(requestData.Schedules[i], folder)
		}
		files, fileNames := WorkWithFiles.GenerateScheduleFile(&requestData)
		zip := WorkWithFiles.ZippingFiles(files, &requestData, folder)

		for _, file := range fileNames {
			os.Remove(file)
		}

		if utils.TeacherWindowCheck(requestData.TeachersS1) || utils.TeacherWindowCheck(requestData.TeachersS2) {
			requestData.Failure = true
		}
		if requestData.Failure {
			attempts++
			if attempts < 9 {
				continue
			}
		}
		// Устанавливаем заголовки для правильного определения файла
		context.Header("Content-Type", "application/zip")
		context.File(zip)
		os.Remove(folder)
		break
	}

}
