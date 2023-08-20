package main

/*import (
	"flag"
	"fmt"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"httpserver/limiter"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
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
	icap := flag.Int("cap", 10, "每秒最大的请求数")
	rate := flag.Int("rate", 5, "每秒可增长的请求数")
	kind := flag.String("kind", "max", "kind")
	flag.Parse()
	rand.Seed(time.Now().UnixNano())
	lmtr := limiter.New(*icap, *rate, *kind)
	r := gin.Default()
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))
	randRouter := r.Group("/random")
	randRouter.Use(func(c *gin.Context) {
		if !lmtr.Allow() {
			qpsCounter.With(prometheus.Labels{
				"method":      "GET",
				"status_code": "500",
			}).Inc()
			c.JSON(http.StatusInternalServerError, "can not handle")
			c.Abort()
		}
		c.Next()
	})
	randRouter.GET("", func(c *gin.Context) {
		qpsCounter.With(prometheus.Labels{
			"method":      "GET",
			"status_code": "200",
		}).Inc()

		c.JSON(http.StatusOK, gin.H{
			"random": rand.Intn(100),
		})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err := r.Run(fmt.Sprintf(":%s", port)); err != nil {
		panic(err)
	}
}
*/
