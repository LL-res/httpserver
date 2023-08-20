package main

import (
	"httpserver/limiter"
	"httpserver/promc"
	"httpserver/router"
)

func main() {
	promc.Init()
	limiter.Init(20, 10)
	r := router.New()
	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
