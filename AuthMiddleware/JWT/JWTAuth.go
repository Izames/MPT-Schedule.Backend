package JWT

import (
	"MPT-Schedule/DataBase"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func JWTAuth() gin.HandlerFunc {
	//перехватывание запроса
	return func(context *gin.Context) {
		context.Set("db", DataBase.DB)

		//получаем заголовок
		authHeader := context.GetHeader("Authorization")
		if authHeader == "" {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "Token is required"})
			context.Abort()
			return
		}
		//получаем сам токен из заголовка
		tokenString := strings.Split(authHeader, " ")[1]

		//парсинг токена и проверка на подлинность
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return privateKey, nil
		})

		if err != nil {
			context.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			context.Abort()
			return
		}
		if !token.Valid {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "Token is invalid"})
			context.Abort()
			return
		}

		context.Next() // Передача управления следующему обработчику
	}
}
