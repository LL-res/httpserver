package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	qpsCounter = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Number of HTTP requests processed.",
	}, []string{"method", "status_code"})
)

func init() {
	prometheus.MustRegister(qpsCounter)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	r := gin.Default()
	r.GET("/random", func(c *gin.Context) {
		qpsCounter.With(prometheus.Labels{
			"method":      "GET",
			"status_code": fmt.Sprintf("%d", c.Writer.Status()),
		}).Inc()

		c.JSON(http.StatusOK, gin.H{
			"random": rand.Intn(100),
		})
	})

	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err := r.Run(fmt.Sprintf(":%s", port)); err != nil {
		panic(err)
	}
}
