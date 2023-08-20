package limiter

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"httpserver/consts"
	"httpserver/promc"
	"net/http"
	"strconv"
	"time"
)

var (
	Rate Limiter
	Max  Limiter
)

type Limiter interface {
	Allow() bool
}

func Init(cap, rate int) {
	Rate = NewR(rate, cap)
	Max = NewM(cap)
}
func New(cap, rate int, kind string) Limiter {
	switch kind {
	case "rate":
		return NewR(rate, cap)
	case "max":
		return NewM(cap)
	}
	return nil
}
func Handler(c *gin.Context) {
	if !Max.Allow() {
		promc.QPSCounter.With(prometheus.Labels{
			consts.METHOD:      c.Request.Method,
			consts.STATUS_CODE: strconv.Itoa(http.StatusInternalServerError),
		}).Inc()
		promc.T3Counter.With(prometheus.Labels{
			consts.METHOD:      c.Request.Method,
			consts.STATUS_CODE: strconv.Itoa(http.StatusInternalServerError),
		})
		promc.ResetByInterval(3*time.Second, promc.T3Counter, &promc.T3Flag)
		c.JSON(http.StatusInternalServerError, "can not handle")
		c.Abort()
	}
	promc.QPSCounter.With(prometheus.Labels{
		consts.METHOD:      c.Request.Method,
		consts.STATUS_CODE: strconv.Itoa(http.StatusOK),
	}).Inc()
	promc.T3Counter.With(prometheus.Labels{
		consts.METHOD:      c.Request.Method,
		consts.STATUS_CODE: strconv.Itoa(http.StatusOK),
	})
	promc.ResetByInterval(3*time.Second, promc.T3Counter, &promc.T3Flag)
	c.Next()
}
