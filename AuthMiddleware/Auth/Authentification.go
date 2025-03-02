package Auth

import (
	"MPT-Schedule/AuthMiddleware/JWT"
	"MPT-Schedule/DataBase"
	"MPT-Schedule/Models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

func Authentification(context *gin.Context) {
	var input Models.User
	//парсинг данных
	if err := context.ShouldBind(&input); err != nil {
		log.Println("не удалось распарсить данные аутентификации")
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//проверка существует ли пользователь при аутентификации
	var user Models.User
	result := DataBase.DB.Where("email=?", input.Email).First(&user)
	if result.Error != nil {
		log.Println("Пользователь [", input.Email, "] пытался войти в систему")
		context.JSON(http.StatusBadRequest, gin.H{"error": "пользователь не существует"})
		return
	}
	//проверка хэшей паролей
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		log.Println("в аккаунт пользователя [", user.Email, "] пытались войти с ложным паролем")
		context.JSON(http.StatusBadRequest, gin.H{"error": "пароли не совпадают"})
		return
	}
	//генерация токена
	token, err := JWT.GenerateJWT(user)
	if err != nil {
		log.Println("у пользователя [", input.Email, "] произошла ошибка генерации JWT токена", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	//отправка токена
	log.Println("пользователь [", input.Email, "] успешно вошел в систему")
	context.JSON(http.StatusOK, gin.H{"jwt": token})
}
