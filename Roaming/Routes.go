package Roaming

import (
	"MPT-Schedule/AuthMiddleware/Auth"
	"MPT-Schedule/Middleware/Requests"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	httpRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of all HTTP requests",
		},
		[]string{"Method", "endpoint"},
	)
	schedulesGenerated = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "schedules_generated_total",
			Help: "Total number of all generated schedules",
		},
	)
	userAuth = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "user_enter_total",
			Help: "user authed",
		},
	)
)

func init() {
	prometheus.MustRegister(httpRequestsTotal)
	prometheus.MustRegister(schedulesGenerated)
	prometheus.MustRegister(userAuth)
}

func Routes() {
	router := gin.Default()
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))
	router.Use(func(context *gin.Context) {
		context.Next()

		httpRequestsTotal.WithLabelValues(context.Request.Method, context.FullPath()).Inc()

		if context.FullPath() == "/generateSchedule" && context.Writer.Status() == 200 {
			schedulesGenerated.Inc()
		}
		if context.FullPath() == "/auth" {
			userAuth.Inc()
		}
	})
	router.POST("/auth", Auth.Authentification)
	router.POST("/register", Auth.Registration)
	router.POST("/sendPin", Auth.SendPinCode)
	router.POST("/updatePassword", Auth.UpdatePassword)
	router.POST("/generateSchedule", Requests.GenerateRequest)

	router.Run(":8091")
}
