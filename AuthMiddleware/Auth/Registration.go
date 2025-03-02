package Auth

import (
	"MPT-Schedule/DataBase"
	"MPT-Schedule/Models"
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
)

func Registration(context *gin.Context) {
	var input Models.User
	var user Models.User

	//парсинг данных
	if err := context.ShouldBind(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	//проверка формата почты
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !re.MatchString(input.Email) {
		context.JSON(http.StatusUnprocessableEntity, gin.H{"error": "неверный формат почты"})
		return
	}

	//проверка не существует ли пользователь уже в системе
	result := DataBase.DB.Where("email=?", input.Email).First(&user)
	if result.Error == nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "пользователь уже есть в системе"})
		return
	}
	//проверка длины пароя
	if len(input.Password) < 6 {
		context.JSON(http.StatusBadRequest, gin.H{"error": "пароль должен быть длиннее 6 символов"})
		return
	}
	//регистрация пользователя
	DataBase.CreateUser(input.Email, input.Password)
	context.JSON(http.StatusCreated, gin.H{"user": input.Email})
}
