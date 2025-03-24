package Roaming

import (
	"MPT-Schedule/AuthMiddleware/Auth"
	"MPT-Schedule/AuthMiddleware/JWT"
	"MPT-Schedule/Middleware/Requests"
	con "context"
	"github.com/gin-gonic/gin"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
	"net/http"
	"time"
)

var auth_count int
var generatedSchedules int
var http_requests int

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

	token := "mhXVbCCeKBsbxi5OCoOUBGGVdxazqkJ_9oBHk9b_rjymBspRFd9f61lQ48-_-OUSQBrMGk3MwSRMgTCxTCby0A=="
	url := "http://localhost:8086"
	client := influxdb2.NewClient(url, token)
	org := "MPT"
	bucket := "metrics"
	writeAPI := client.WriteAPIBlocking(org, bucket)
	router := gin.Default()
	router.Use(CORSMiddleware())
	router.Use(func(context *gin.Context) {
		context.Next()
		if context.Request.URL.Path != "/metrics" {
			http_requests++
			tags := map[string]string{
				"method":   context.Request.Method,
				"endpoint": context.FullPath(),
			}
			fields := map[string]interface{}{
				"count": http_requests,
			}
			point := write.NewPoint("http_requests", tags, fields, time.Now())
			writeAPI.WritePoint(con.Background(), point)
		}
		if context.FullPath() == "/generateSchedule" && context.Writer.Status() == 200 {
			generatedSchedules++
			fields := map[string]interface{}{
				"count": generatedSchedules,
			}
			point := write.NewPoint("schedules_generated", nil, fields, time.Now())
			writeAPI.WritePoint(con.Background(), point)
		}

		if context.FullPath() == "/auth" {
			auth_count++
			fields := map[string]interface{}{
				"count": auth_count,
			}
			point := write.NewPoint("user_auth", nil, fields, time.Now())
			writeAPI.WritePoint(con.Background(), point)
		}
	})
	router.POST("/auth", Auth.Authentification)
	router.POST("/register", Auth.Registration)
	router.POST("/sendPin", Auth.SendPinCode)
	router.POST("/updatePassword", Auth.UpdatePassword)
	//приватные ссылки
	privateRoutes := router.Group("/schedule")
	privateRoutes.Use(JWT.JWTAuth())
	privateRoutes.POST("/generateSchedule", Requests.GenerateRequest)

	router.Run(":8091")
}
