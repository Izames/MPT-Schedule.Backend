package Requests

import (
	"MPT-Schedule/Middleware/GetData"
	"MPT-Schedule/Middleware/InsertData"
	"MPT-Schedule/Models"
	"github.com/gin-gonic/gin"
	"log"
)

func GenerateRequest(context *gin.Context) {
	//парсинг данных в мультипарт режиме
	if err := context.Request.ParseMultipartForm(32 << 22); err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
		log.Println("не удалось распарсить данные для генерации расписания")
		return
	}
	//парсинг запроса в переменную
	form, err := context.MultipartForm()

	if err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
		log.Println("не удалось распарсить данные для генерации расписания")
		return
	}
	extractFiles := form.File["extracts"]
	if len(extractFiles) == 0 {
		context.JSON(400, gin.H{"error": "no files with the name 'extracts' found"})
		log.Println("не удалось найти файлы с именем 'extracts'")
		return
	}
	//Десериализация данных
	if err = GetData.Deserialization(form); err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
		log.Println("не удалось распарсить данные для генерации расписания")
		return
	}
	//собрать ограничения учителей
	GetData.GetTeachers()
	GetData.GetGroups()
	for _, file := range Models.Extracts {
		GetData.GetLessons(file)
	}
	for _, group := range Models.Groups {
		InsertData.GenerateSchedule(group)
	}

	println("")
	//дублируем будущее расписание
	//WorkWithFiles.DuplicateFile()
	//file := InsertData.CreateSchedule()
	//path := WorkWithFiles.Zipping(file)
	//file.Close()
	//context.File(path)
	//os.Remove(path)
	//os.Remove("Schedule1.xlsx")
	//os.Remove("Schedule2.xlsx")
	defer Models.Clean()
}
