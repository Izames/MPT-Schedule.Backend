package Roaming

import (
	"MPT-Schedule/AuthMiddleware/Auth"
	"MPT-Schedule/Middleware/Requests"
	con "context"
	"github.com/gin-gonic/gin"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
	"time"
)

var auth_count int
var generatedSchedules int
var http_requests int

func Routes() {

	token := "ZtioJd3fB0gaZRNuF1RFkeHrhuvelVwYrroVBgTXdPG70XjQACcJUQZW-AFfs1LG9B4RRIK3MXxVnw1BFscLHg=="
	url := "http://localhost:8086"
	client := influxdb2.NewClient(url, token)
	org := "MPT"
	bucket := "metrics"
	writeAPI := client.WriteAPIBlocking(org, bucket)

	router := gin.Default()
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
	router.POST("/generateSchedule", Requests.GenerateRequest)

	router.Run(":8091")
}
