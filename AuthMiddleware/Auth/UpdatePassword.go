package Auth

import (
	"MPT-Schedule/DataBase"
	"MPT-Schedule/Models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func UpdatePassword(context *gin.Context) {
	var passwordData Models.PasswordUpdate
	//парсинг данных
	err := context.ShouldBindJSON(&passwordData)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "ошибка парсинга данных смены пароля"})
		return
	}
	//нахождение действующего пин кода и есть ли он вообще
	pin, err := DataBase.FindPinByMail(passwordData.Email)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error ": "действующего пин кода не найдено!"})
		log.Println("пользователь [ ", passwordData.Email, " ] пытался восстановить пин код без действующего пин кода")
		return
	}
	//проверка пин кода с введенным
	if pin.Pin != passwordData.Pin {
		context.JSON(http.StatusBadRequest, gin.H{"error ": "неверный пин код!"})
		log.Println("пользователь [ ", passwordData.Email, " ] пытался войти по неверному пин коду")
		return
	}
	//проверка длины нового пароля
	if len(passwordData.Password) < 6 {
		context.JSON(http.StatusBadRequest, gin.H{"error": "пароль должен быть длиннее 6 символов"})
		return
	}
	//проверка нужного пользователя
	user := DataBase.FindUserByEmail(passwordData.Email)

	//обновление пароля
	user.Password = passwordData.Password
	result := DataBase.Update_user(user)
	if result != "" {
		log.Println(result)
		context.JSON(http.StatusInternalServerError, gin.H{"error": result})
	}
	context.JSON(http.StatusOK, gin.H{"success": "success"})
}
