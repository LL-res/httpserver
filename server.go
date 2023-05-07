package main

import (
	"flag"
	"fmt"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"math/rand"
	"net/http"
	"os"
	"sync"
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

type Limiter struct {
	sync.Mutex
	rate        int
	cap         int
	cur         int
	lastRequest time.Time
}

func NewLimiter(rate int, cap int) *Limiter {
	return &Limiter{
		rate: rate,
		cap:  cap,
	}
}
func (l *Limiter) Allow() bool {
	l.Lock()
	defer l.Unlock()
	now := time.Now()
	l.cur += int(now.Sub(l.lastRequest).Seconds()) * l.rate
	l.lastRequest = now
	if l.cur > l.cap {
		l.cur = l.cap
	}
	if l.cur > 0 {
		l.cur--
		return true
	}
	return false
}

func main() {
	icap := flag.Int("cap", 10, "每秒最大的请求数")
	rate := flag.Int("rate", 5, "每秒可增长的请求数")
	flag.Parse()
	rand.Seed(time.Now().UnixNano())
	limiter := NewLimiter(*rate, *icap)
	r := gin.Default()
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))
	randRouter := r.Group("/random")
	randRouter.Use(func(c *gin.Context) {
		if !limiter.Allow() {
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
