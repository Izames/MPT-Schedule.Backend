package Roaming

import (
	"MPT-Schedule/AuthMiddleware/Auth"
	"MPT-Schedule/AuthMiddleware/JWT"
	"MPT-Schedule/Middleware/Requests"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}
func Routes() {
	router := gin.Default()
	router.Use(CORSMiddleware())
	router.POST("/auth", Auth.Authentification)
	router.POST("/register", Auth.Registration)
	router.POST("/sendPin", Auth.SendPinCode)
	router.POST("/updatePassword", Auth.UpdatePassword)
	//приватные ссылки
	privateRoutes := router.Group("/schedule")
	privateRoutes.Use(JWT.JWTAuth())
	privateRoutes.POST("/generateSchedule", Requests.GenerateRequest)

	router.Run(":8090")
}
