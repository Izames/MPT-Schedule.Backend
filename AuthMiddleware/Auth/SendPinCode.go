package Auth

import (
	"MPT-Schedule/DataBase"
	"MPT-Schedule/Mail"
	"MPT-Schedule/Models"
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"regexp"
)

func SendPinCode(context *gin.Context) {
	var pin Models.PinCode
	var oldPin Models.PinCode
	var user Models.User
	//парсинг данных
	if err := context.ShouldBindJSON(&pin); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "ошибка парсинга пин кода"})
	}
	//проверка формата почты
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !re.MatchString(pin.Email) {
		context.JSON(http.StatusUnprocessableEntity, gin.H{"error": "неверный формат почты"})
		return
	}
	//проверка существует ли пользователь
	result := DataBase.DB.Where("email=?", pin.Email).First(&user)
	if result.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "такого пользователя нет в системе"})
		return
	}
	//проверка есть ли в бд уже нужный пин код пользователя
	result = DataBase.DB.Where("email=?", pin.Email).First(&oldPin)

	//обработка результата и отправка нового, или старого пин кода
	if result.Error != nil {
		pin.Pin = fmt.Sprintf("%06d", rand.Intn(1000000))
		if err := DataBase.CreatePin(pin.Email, pin.Pin); err != "" {
			context.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		if err := Mail.SendEmail("Ваш ПинКод : "+pin.Pin, pin.Email, "ПинКод для восстановления пароля"); err != "" {
			context.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
	} else {
		if err := Mail.SendEmail("Ваш ПинКод : "+oldPin.Pin, oldPin.Email, "ПинКод для восстановления пароля"); err != "" {
			context.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
	}

	context.JSON(http.StatusAccepted, gin.H{"status": "success"})
}
