package router

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"httpserver/handler"
	"httpserver/limiter"
)

func New() *gin.Engine {
	r := gin.Default()
	r.GET("/", handler.Index)
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))
	r.POST("/purge", handler.PurgeReqTotal)
	limitGroup := r.Group("/product")
	{
		limitGroup.Use(limiter.Handler)
		limitGroup.POST("/create", handler.CreateProduct)
		limitGroup.GET("/list", handler.GetAllProducts)
	}

	return r
}
